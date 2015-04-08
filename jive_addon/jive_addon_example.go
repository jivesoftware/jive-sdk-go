package main

import (
	"net/http"
	"github.com/codegangsta/negroni"
	"fmt"
	/*
	"os"
	"log"
	"encoding/json"
	*/
)

func main() {
	
	mux := http.NewServeMux()
	
	/*** ADD-ON ***/
	mux.Handle("/configure", http.StripPrefix("/configure", http.FileServer(http.Dir("addon/configure")))).Methods("GET","HEAD")
	//TODO:
	mux.HandleFunc("/register", handleAddOnRegister).Methods("POST")

	//TODO: REPLACE WITH func(w http.ResponseWriter, r *http.Request)!!

	/*** TILES ***/
	mux.Handle("/tiles/example-stream-tile/configure", http.StripPrefix("/tiles/example-stream-tile/configure", http.FileServer(http.Dir("tiles/example-stream-tile/configure")))).Methods("GET","HEAD","POST")
	//TODO:
	mux.HandleFunc("/tiles/example-stream-tile/register", handleTileRegister).Methods("POST")
	
	/*** APPS ***/
	mux.Handle("/apps", http.StripPrefix("/apps", http.FileServer(http.Dir("apps")))).Methods("GET","HEAD")
	
	/*** WEBHOOKS ***/
	//TODO:
	mux.HandleFunc("/webhooks/content", handleWebhooks).Methods("POST")
	//TODO:
	mux.HandleFunc("/webhooks/system", handleWebhooks).Methods("POST")
	
	n := negroni.Classic()
	n.UseHandler(mux)
    n.Run(":8090")
	
} // end main

func handleAddOnConfigure(res http.ResponseWriter, req *http.Request) {
	fmt.Println("handleAddOnConfigure called...")
}

func handleAddOnRegister(res http.ResponseWriter, req *http.Request) {
	fmt.Println("handleAddOnRegister called...")
}

func handleTileRegister(res http.ResponseWriter, req *http.Request) {
	fmt.Println("handleTileRegister called...")
}

func handleTileConfigure(res http.ResponseWriter, req *http.Request) {
	fmt.Println("handleTileConfigure called...")
}

func handleAppXml(res http.ResponseWriter, req *http.Request) {
	 fmt.Println("handleAppXml called...")
}

func handleApp(res http.ResponseWriter, req *http.Request) {
	fmt.Println("handleApp called...")
}

func handleWebhooks(res http.ResponseWriter, req *http.Request) {
	fmt.Println("handleWebhooks called...")
}

