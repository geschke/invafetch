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
	"github.com/spf13/viper"
)

var repository *invdb.Repository
var collectProcessData []golrackpi.ProcessData

type CollectDaemon struct {
	AuthData golrackpi.AuthClient
	lib      *golrackpi.AuthClient
}

/*type JsonProcessdataValues []golrackpi.ProcessDataValues

func (pdv JsonProcessdataValues) MarshalJSON() ([]byte, error) {
	type localPdv JsonProcessdataValues
	valueJson, err := json.Marshal(localPdv(pdv))
	if err != nil {
		return nil, err
	}

	fmt.Println("in MarshalJSON")
	//cnt := "foobar" + strconv.Itoa(int(c.Found))
	//fmt.Println(cnt)
	//fmt.Println(strconv.Itoa(int(c.Found)))
	return json.Marshal(map[string]interface{}{
		"moduleid": map[string]interface{}{
			"foo": json.RawMessage(valueJson),
			//		cnt:      json.RawMessage(valueJson),
		},
	})
}
*/

type PdvMap map[string]golrackpi.ProcessDataValue

//func convertPdvMap(pdv []ProcessDataValueJSON) PdvMap {
func convertPdvMap(pdv []golrackpi.ProcessDataValue) PdvMap {
	pdvmap := make(PdvMap)

	for i := range pdv {
		fmt.Println(pdv[i].Id)
		pdvmap[pdv[i].Id] = pdv[i]
	}
	return pdvmap
}

func convertPdvsMap(pdvs []golrackpi.ProcessDataValues) map[string]PdvMap {
	pdvsmap := make(map[string]PdvMap)
	for i := range pdvs {
		fmt.Println(pdvs[i].ModuleId)
		pdvsmap[pdvs[i].ModuleId] = convertPdvMap(pdvs[i].ProcessData)
	}
	return pdvsmap
}

func convertStuff(pd []golrackpi.ProcessDataValues) {

	//jpd := JsonProcessdataValues(pd)
	pdvsmap := convertPdvsMap(pd)
	foo, _ := json.Marshal(pdvsmap)
	// todo: error handling
	fmt.Println(string(foo))
	repository.AddData(string(foo))
	//panic("die hard")
}

func (cd *CollectDaemon) genNewId(id int) int {
	id++
	return id
}

func (cd *CollectDaemon) innerLoop(ctx context.Context, i int) int {
	log.Println("in innerLoop mit i ", i)

	timer2 := time.NewTimer(30 * time.Second)
	ticker := time.NewTicker(3 * time.Second)

	for active := true; active; {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
			//pd, err := cd.lib.ProcessDataModule("devices:local")
			pd, err := cd.lib.ProcessDataValues(collectProcessData)
			if err != nil {
				fmt.Println(err)
				panic("hard error")
			}
			fmt.Println(pd)
			convertStuff(pd)

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

	return i
}

func (cd *CollectDaemon) outerLoop(ctx context.Context) {
	timer1 := time.NewTimer(10 * time.Hour)

	log.Println("in outerLoop start")
	id := 0
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
				log.Println("in for mit id ", id, " und cnt:", cnt, " and before innerLoop", time.Now())
				err := cd.logoutLogin()
				if err != nil {
					fmt.Println(err)
					panic("hard error 2") // todo error handling
				}
				id = cd.innerLoop(ctx, id)
				log.Println("after innerLoop id", id, time.Now())
				id = cd.genNewId(id)
				log.Println("after genNewId:", id, " time:", time.Now())
				cnt++
			}
		}
	}()

	for active := true; active; {
		select {

		case <-timer1.C:
			log.Println("timer1 Out fired, simulates end of program, give time to end innerLoop", time.Now())
			done <- true
			time.Sleep(100 * time.Millisecond)
			active = false

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
	return
}

func (cd *CollectDaemon) logoutLogin() error {
	ok, err := cd.lib.Logout()
	if err != nil {
		//fmt.Println("logout error", err)
		return fmt.Errorf("logout error: %s", err)
	}
	fmt.Println("logout ok?", ok)

	fmt.Println("Try another login...")
	cd.lib = golrackpi.NewWithParameter(cd.AuthData)

	fmt.Println(cd.lib.SessionId)

	sessionId, err := cd.lib.Login()
	if err != nil {
		fmt.Println("An error occurred:", err)
		return fmt.Errorf("login error: %s", err)
	}
	fmt.Println("SessionId", sessionId)

	return nil
}

func GetDbConfig() dbconn.DatabaseConfiguration {
	// could something go wrong here?
	fmt.Println(viper.Get("dbName"))
	fmt.Println(viper.Get("dbHost"))
	fmt.Println(viper.Get("dbUser"))
	fmt.Println(viper.Get("dbPassword"))
	fmt.Println(viper.Get("dbPort"))
	var config dbconn.DatabaseConfiguration
	config.DBHost = viper.GetString("dbHost")
	config.DBName = viper.GetString("dbName")
	config.DBPassword = viper.GetString("dbPassword")
	config.DBUser = viper.GetString("dbUser")
	config.DBPort = viper.GetString("dbPort")
	return config
}

type ProcessDataValueJSON golrackpi.ProcessDataValue

//type PdvMap map[string]ProcessDataValueJSON

func (pdv ProcessDataValueJSON) MarshalJSON() ([]byte, error) {
	type localPdv ProcessDataValueJSON

	valueJson, err := json.Marshal(localPdv(pdv))
	if err != nil {
		return nil, err
	}

	fmt.Println("in MarshalJSON")
	messageid := pdv.Id
	//cnt := "foobar" + strconv.Itoa(int(c.Found))
	//fmt.Println(cnt)
	//fmt.Println(strconv.Itoa(int(c.Found)))
	return json.Marshal(map[string]interface{}{
		messageid: json.RawMessage(valueJson),
		//		cnt:      json.RawMessage(valueJson),
	},
	)
}

func (cd *CollectDaemon) Start(configProcessData []golrackpi.ProcessData) {
	/*
		pdv := golrackpi.ProcessDataValue{
			Unit:  "",
			Id:    "Statistic:CO2Saving:Day",
			Value: 16672.1581222652,
		}

		pdv2 := golrackpi.ProcessDataValue{
			Unit:  "",
			Id:    "Statistic:Foo",
			Value: 10,
		}

		pdv3 := golrackpi.ProcessDataValue{
			Unit:  "",
			Id:    "Statistic:CO2Saving:Day",
			Value: 16672.1581222652,
		}

		pdv4 := golrackpi.ProcessDataValue{
			Unit:  "",
			Id:    "Statistic:Bar",
			Value: 20000,
		}

		//pdvj := ProcessDataValueJSON(pdv)
		//pdvj2 := ProcessDataValueJSON(pdv2)

		var pdvarr []golrackpi.ProcessDataValue
		var pdvarr2 []golrackpi.ProcessDataValue
		//pdvarr = append(pdvarr, pdvj, pdvj2)
		pdvarr = append(pdvarr, pdv, pdv2)
		pdvarr2 = append(pdvarr2, pdv2, pdv3, pdv4)

		var pdvsarr []golrackpi.ProcessDataValues

		pdvs := golrackpi.ProcessDataValues{
			ModuleId:    "scb:statistic:EnergyFlow",
			ProcessData: pdvarr,
		}

		pdvs2 := golrackpi.ProcessDataValues{
			ModuleId:    "scb:statistic:Nothing",
			ProcessData: pdvarr2,
		}

		pdvsarr = append(pdvsarr, pdvs, pdvs2)

		result := convertPdvsMap(pdvsarr)

		//result := convertPdvMap(pdvarr)

		//fmt.Println(pdv)
		//fmt.Println(pdv2)
		//fmt.Println(pdvarr)

		fmt.Println("after convert", result)
		jsonPdv, err := json.Marshal(result)
		if err != nil {
			fmt.Println("json convert error")
		}
		jsonOut := string(jsonPdv)
		fmt.Println(jsonOut)
		return
	*/

	cd.lib = golrackpi.NewWithParameter(cd.AuthData)

	config := dbconn.ConnectDB(GetDbConfig())

	repository = invdb.NewRepository(config)

	collectProcessData = configProcessData

	fmt.Println(cd.lib.SessionId)

	sessionId, err := cd.lib.Login()
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	fmt.Println("SessionId", sessionId)

	/*	err = cd.logoutLogin()
		if err != nil {
			fmt.Println("Error when connecting again", err)
			return
		}

		ok, err := cd.lib.Logout()
		if err != nil {
			fmt.Println("logout error", err)
			return
		}
		fmt.Println("logout ok?", ok)

		return*/

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
		log.Println("Abbruch mit Ctrl+C")

		cancel()
	}()

	cd.outerLoop(ctx)

	//innerLoop(id)
	log.Println("after outerLoop")

	cd.PrintMemUsage()

}

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
