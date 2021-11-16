package looptry

import (
	"log"
	"net/http"
)

func TestLoop(ciclos int) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		log.Panic(err)
	}

	log.Println(resp)

	if ciclos > 7 {
		log.Panic("limite excedido")
	}
}
