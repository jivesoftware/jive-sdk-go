package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"jive_sdk"
)

var clientID string
var clientSecret string

/* 
// Use this initializing function if clientID and clientSecret will be stored 
// in a DB or file and passed in procedurally
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
		}
			jive_sdk.IsValidRegistraton(data, clientSecret)
        })
 
 http.HandleFunc("/unregister", func(rw http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		data := jive_sdk.Payload{}
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(err)
		}
		secret := clientSecret
			jive_sdk.IsValidRegistraton(data, secret)
        })
 
    http.ListenAndServe(":8090", nil)
}