/*
Test Service for Jive-SDK written in Go Lang
*/

package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	
	"github.com/jivesoftware/jive-sdk-go-gpm"
)

var clientID string
var clientSecret string

/* 
// Use this initializing function if clientID and clientSecret will be stored 
// in a DB or file and passed in procedurally or any activities before runtime
func init {
   // initialization code here    
}
*/

func main (){
    http.HandleFunc("/register", func(rw http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		data := jive_sdk.Payload{}
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
		}
		statusCode := jive_sdk.IsValidRegistraton(data, clientSecret)
		if statusCode{
			rw.WriteHeader(http.StatusNoContent)
		}else{
			rw.WriteHeader(http.StatusForbidden)	
		}
	})
 
 	http.HandleFunc("/unregister", func(rw http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		data := jive_sdk.Payload{}
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
		}
		statusCode := jive_sdk.IsValidRegistraton(data, clientSecret)
		if statusCode{
			rw.WriteHeader(http.StatusNoContent)
		}else{
			rw.WriteHeader(http.StatusForbidden)	
		}
	})
 
	http.ListenAndServe(":8090", nil)
}