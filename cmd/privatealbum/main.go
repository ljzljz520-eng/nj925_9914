package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"privatealbum/internal/flow062"
	"privatealbum/internal/store"
)

func main() {
	path := flag.String("db", "privatealbum.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	s, e := store.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	fmt.Printf("private album service listening on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, flow062.New(s).Handler()))
}
