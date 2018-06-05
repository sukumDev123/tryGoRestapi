package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

func main(){
	// Init Router
	r := mux.NetRouter()
	// Router Handlers /enpoints
	r.HandlerFunc("/api/books",getBook)
	r.HandlerFunc("/api/books",getBook)
	r.HandlerFunc("/api/books",getBook)
	r.HandlerFunc("/api/books",getBook)
	
}