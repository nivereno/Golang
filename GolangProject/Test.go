package main

import (
	"log"
	"net/http"
	"os"

	//"handlers"
	"github.com/nivereno/Golang/tree/main/GolangProject/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", nil)
}
