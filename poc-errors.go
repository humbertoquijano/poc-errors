package main

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"

	looptry "mti.com/poc-errors/loop-try"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup

	miArg := os.Args[1]

	miVal, err := strconv.Atoi(miArg)

	if err != nil {
		log.Fatal("no se puede convertir argumento a entero")
	}

	if miVal < 0 {
		log.Fatal("valor de argumento no válido")
	}

	miVal++
	for i := 1; i < miVal; i++ {
		wg.Add(1)
		go looptry.TestLoop(i, &wg)
	}

	wg.Wait()

	log.Println("test OK")
}
