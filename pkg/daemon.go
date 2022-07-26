// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/geschke/golrackpi"
	"github.com/geschke/invafetch/pkg/dbconn"
	"github.com/geschke/invafetch/pkg/invdb"
)

var repository *invdb.Repository
var collectProcessData []golrackpi.ProcessData

type CollectDaemon struct {
	AuthData golrackpi.AuthClient
	DbConfig dbconn.DatabaseConfiguration
	lib      *golrackpi.AuthClient
}

// PdvMap defines a map of ProcessDataValue type with a string as key
type PdvMap map[string]golrackpi.ProcessDataValue

// convertPdvMap converts a slice of ProcessDataValue into a map of ProcessDataValue type with the Id as key
func convertPdvMap(pdv []golrackpi.ProcessDataValue) PdvMap {
	pdvmap := make(PdvMap)

	for i := range pdv {
		pdvmap[pdv[i].Id] = pdv[i]
	}
	return pdvmap
}

// convertPdvsMap converts a slice of ProcessDataValues into a map of PdvMap type with the ModuleId as key
func convertPdvsMap(pdvs []golrackpi.ProcessDataValues) map[string]PdvMap {
	pdvsmap := make(map[string]PdvMap)
	for i := range pdvs {
		//fmt.Println(pdvs[i].ModuleId)
		pdvsmap[pdvs[i].ModuleId] = convertPdvMap(pdvs[i].ProcessData)
	}
	return pdvsmap
}

// convertProcessDataValues converts a slice of processdata values into JSON and returns the result as slice of bytes, which can be stored in the MariaDB database
func convertProcessDataValues(pd []golrackpi.ProcessDataValues) ([]byte, error) {
	var pdvsJSON []byte
	pdvsmap := convertPdvsMap(pd)
	pdvsJSON, err := json.Marshal(pdvsmap)
	if err != nil {
		return pdvsJSON, err
	}
	return pdvsJSON, nil
}

// innerLoop gets, converts and stores inverter values every tickerTimeSeconds seconds. It runs a duration of newLoginTimeMinutes,
// then the timer and ticker are stopped.
func (cd *CollectDaemon) innerLoop(ctx context.Context, newLoginTimeMinutes int64, tickerTimeSeconds int64) {
	log.Println("start innerLoop")

	timer2 := time.NewTimer(time.Duration(newLoginTimeMinutes) * time.Minute)
	ticker := time.NewTicker(time.Duration(tickerTimeSeconds) * time.Second)

	// todo: if something failed, repeat 5 or 10 times, then set fail status
	for active := true; active; {
		select {
		//case t := <-ticker.C:
		case <-ticker.C:

			//fmt.Println("Tick at", t)

			pd, err := cd.lib.ProcessDataValues(collectProcessData)

			if err != nil {
				fmt.Println(err)
				panic("hard error")
			}
			//fmt.Println(pd)
			pdvsJSON, err := convertProcessDataValues(pd)
			//err = errors.New("foo error")
			if err != nil {
				// fail silently
				continue
			}
			//lastId, err := repository.AddData(string(pdvsJSON))
			_, err = repository.AddData(string(pdvsJSON))
			if err != nil {
				// fail silently
				continue
			}
			//log.Println("entry", lastId, "added")

		case <-timer2.C:
			log.Println("timer2 fired")
			ticker.Stop()
			timer2.Stop()
			active = false

		case <-ctx.Done():
			log.Println("ctx done in inner fired!", time.Now())
			timer2.Stop()
			ticker.Stop()
			active = false
		}
	}
	timer2.Stop()
	ticker.Stop()
	log.Print("end innerLoop\n\n")

}

// outerLoop logs into inverter and opens a database connection. Then it gives control to innerLoop, which gets ProcessDataValues.
// After newLoginTimeMinutes a logout and connection close is performed, so the loop starts again. The whole process can be stopped by a context signal.
func (cd *CollectDaemon) outerLoop(ctx context.Context, newLoginTimeMinutes int64, tickerTime int64) {
	//timer1 := time.NewTimer(10 * time.Hour)

	log.Println("in outerLoop start")

	cnt := 0
	done := make(chan bool)
	go func() {
		for active1 := true; active1; {
			select {
			case <-done:
				log.Println("in outer done received, set active to false!")
				active1 = false
			default:
				cd.PrintMemUsage()
				err := cd.login()
				if err != nil {
					log.Println(err)
					panic("hard error 1")
				}
				log.Println("in for cnt:", cnt, " and before innerLoop", time.Now())
				err = cd.openDbRepository()
				if err != nil {
					log.Println(err)
					panic("hard error 4")
				}

				cd.innerLoop(ctx, newLoginTimeMinutes, tickerTime)
				log.Println("after innerLoop", time.Now())
				err = cd.logout()
				if err != nil {
					log.Println(err)
					panic("hard error 2") // todo error handling
				}
				err = cd.closeDbRepository()
				if err != nil {
					log.Println(err)
					panic("hard error 3")
				}

				cnt++
			}
		}
	}()

	for active := true; active; {
		select {

		/*case <-timer1.C:
		log.Println("timer1 Out fired, simulates end of program, give time to end innerLoop", time.Now())
		done <- true
		time.Sleep(100 * time.Millisecond)
		active = false
		*/
		case <-ctx.Done():
			log.Println("ctx done in main fired!")
			log.Println("Wait time of innerLoop to let it end gracefully...")
			done <- true
			time.Sleep(500 * time.Millisecond)
			log.Println("ticker ctx stopped in outer select", time.Now())
			active = false
		}

	}
	log.Println("outerLoop ends normal")
}

// logout performs a logout from inverter
func (cd *CollectDaemon) logout() error {
	ok, err := cd.lib.Logout()
	if err != nil {
		return fmt.Errorf("logout error: %s", err)
	}
	log.Println("logout ok?", ok)
	return nil
}

// login performs a login into the inverter
func (cd *CollectDaemon) login() error {
	log.Println("Try another login...")
	cd.lib = golrackpi.NewWithParameter(cd.AuthData)

	sessionId, err := cd.lib.Login()
	if err != nil {
		log.Println("An error occurred:", err)
		return fmt.Errorf("login error: %s", err)
	}
	log.Println("Login ok with SessionId", sessionId)
	return nil
}

// openDbRepository opens the database connection
func (cd *CollectDaemon) openDbRepository() error {
	db, err := dbconn.ConnectDB(cd.DbConfig)
	if err != nil {
		return err
	}
	repository = invdb.NewRepository(db)
	return nil
}

// closeDbRepository closes the database connection
func (cd *CollectDaemon) closeDbRepository() error {
	err := repository.Close()
	if err != nil {
		return err
	}
	return nil
}

// Start takes configuration parameters and starts the data collecting process
func (cd *CollectDaemon) Start(configProcessData []golrackpi.ProcessData, newLoginTimeMinutes int64, tickerTimeSeconds int64) {

	cd.lib = golrackpi.NewWithParameter(cd.AuthData)
	collectProcessData = configProcessData

	cd.PrintMemUsage()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//ctx, cancel = context.WithTimeout(ctx, time.Duration(10*time.Second))

	c := make(chan os.Signal)

	//c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, os.Interrupt)
		//signal.Notify(c, os.Kill)
		<-c
		log.Println("Stopped by Ctrl+C")

		cancel()
	}()

	cd.outerLoop(ctx, newLoginTimeMinutes, tickerTimeSeconds)

	log.Println("Daemon stopped")

	cd.PrintMemUsage()

}

// PrintMemUsage writes some memory statistics of the running Go process
// https://golangcode.com/print-the-current-memory-usage/
func (cd *CollectDaemon) PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", cd.bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", cd.bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", cd.bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Println()
}

func (cd *CollectDaemon) bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
