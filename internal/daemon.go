package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"
)

type CollectDaemon struct {
}

func (cd *CollectDaemon) genNewId(id int) int {
	id++
	return id
}

func (cd *CollectDaemon) innerLoop(ctx context.Context, i int) int {
	log.Println("in innerLoop mit i ", i)

	timer2 := time.NewTimer(10000 * time.Millisecond)
	ticker := time.NewTicker(950 * time.Millisecond)

	for active := true; active; {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)

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

func (cd *CollectDaemon) Start() {

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
