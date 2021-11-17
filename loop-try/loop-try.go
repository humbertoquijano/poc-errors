package looptry

import (
	"bufio"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func TestLoop(ciclos int, wg *sync.WaitGroup) {
	idx := ciclos % 10
	if idx == 0 {
		idx = 10
	}

	url := "https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(idx)

	resp, err := http.Get(url)

	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	var sb strings.Builder

	for i := 0; scanner.Scan(); i++ {
		sb.WriteString(scanner.Text())
	}

	log.Println(sb.String())

	if ciclos > 7 {
		log.Panic("limite excedido")
	}

	defer wg.Done()
}
