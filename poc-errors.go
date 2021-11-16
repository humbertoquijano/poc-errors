package main

import (
	"log"
	"os"
	"strconv"

	looptry "mti.com/poc-errors/loop-try"
)

func main() {
	miArg := os.Args[1]

	miVal, err := strconv.Atoi(miArg)

	if err != nil {
		log.Fatal("no se puede convertir argumento a entero")
	}

	if miVal < 0 {
		log.Fatal("valor de argumento no válido")
	}

	for i := 0; i < miVal; i++ {
		looptry.TestLoop(i)
	}

	log.Println("test OK")
}
