package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/kaakaa/matterpoll-emoji/poll"
)

var port = flag.Int("p", 8505, "port number")

func main() {
	flag.Parse()

	c, err := poll.LoadConf("config.json")
	if err != nil {
		log.Fatal(err)
	}
	poll.Conf = c
	http.HandleFunc("/poll", poll.PollCmd)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatal(err)
	}
}
