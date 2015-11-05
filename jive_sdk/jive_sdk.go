package jive_sdk

import (
	"fmt"
	"reflect"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"net/http"
	"bytes"
	"encoding/hex"
	"sort"
	"unicode"
)

// This struct outlines the parameters of the payload from Jive
// includes field translation from JSON to more Go friendly syntax
type Payload struct {
	ClientId string `json:"clientID"`
	Code string `json:"code"`
	Scope string `json:"scope"`
	TenantId string `json:"tenantID"`
	JiveSignatureURL string `json:"JiveSignatureURL"`
	ClientSecret string `json:"clientSecret,omitempty"`
	JiveSignature string `json:"jiveSignature"`
	JiveUrl string `json:"jiveURL"`
	Timestamp string `json:"timestamp"`
}

func IsValidRegistraton(payload Payload, existingSecret string) bool{
	//Using this later
	jiveSignatureURL := ""
	jiveSignature := ""
	
	// Make a map from the payload and filter out the JiveSignature 
	validationBlock := make(map[string]string)
	payloadValues := reflect.ValueOf(payload)
	payloadTypes := reflect.TypeOf(payload)
	size := payloadValues.NumField()
	for i := 0; i < size; i++ {
		key := payloadTypes.Field(i).Name
		value := payloadValues.Field(i).Interface().(string)
		if key != "JiveSignature"{
			validationBlock[key] = value
		}else {
			jiveSignature = value
		}
	}
	
	// Logic for clientSecret
	// Check if there is a clientSecret that exists outside of the payload
	if existingSecret == ""{
		if validationBlock["ClientSecret"] == ""{
			panic("Registration event with no clientSecret. Invalid payload")
			return false
		}else{
			secret := []byte(validationBlock["ClientSecret"])                                                           
			h := sha256.New()
			h.Write(secret)                                                    
			validationBlock["ClientSecret"] = hex.EncodeToString(h.Sum(nil))
		}
	}else{
		if len(validationBlock["ClientSecret"]) != 0{
			fmt.Println("ClientSecret already in payload, ignoring argument. Make sure you are not passing in clientID on register events")
		}else{
			validationBlock["ClientSecret"] = existingSecret
		}
	}
	
	// Byte string of all keys to be sorted in alphabetical order
	// After sorting, keys and values are strung together to be sent to JiveSignatureURL  
	keys := []string{}
	for k := range validationBlock {
		keys = append(keys, k)
	}
	
	body := ""
	sort.Strings(keys)
	for _, k := range keys {
		lowerK :=[]rune(k)
		lowerK[0] = unicode.ToLower(lowerK[0])
		key := string(lowerK)
		body += key + ":" + validationBlock[k] + "\n"
	}
	
	// Post a request to Jive Market to Validate the Signature
    jiveSignatureURL = validationBlock["JiveSignatureURL"]
	req, err := http.NewRequest("POST", jiveSignatureURL, bytes.NewBufferString(body))
	req.Header.Set("X-Jive-MAC", jiveSignature)
	fmt.Printf("\n\nJive URL: %s\n\nBody: %v\n\nSignature: %s\n", jiveSignatureURL, bytes.NewBufferString(body), jiveSignature)
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	if resp.StatusCode == 204{
		fmt.Printf("Validation Successful [%v]\n",resp.Status)
		return true
	}
	fmt.Printf("Validation Failed [%v]", resp.Status)
	return false
}


func IsValidJiveRequest(authorization string, clientId string, clientSecret string) bool {
	if (authorization == "") {
		fmt.Println("Invalid Authorization (null/empty)")
		return false
	}
	tokens := strings.Split(authorization, " ")
	
	if (tokens[0] != "JiveEXTN") {
		fmt.Printf("Invalid Authorization Type [%T]\n",tokens[0])
		return false
	}
	
	payload := make(map[string]string)
	
	for _, param := range strings.Split(tokens[1],"&") {
		idx := strings.Index(param, "=")
		key := param[:idx]
		value := param[idx+1:]
		payload[key] = value
	}
	
	if (payload["client_id"] != clientId) {
		return false
	}
	
	signature := payload["signature"]
	delete(payload,"signature")
	
	message := ""
	for key, value := range payload {
		message += "&" + key + "=" + value
	}
	message = message[1:]
	
	if (strings.HasSuffix(clientSecret,".s")) {
		clientSecret = clientSecret[:len(clientSecret)-len(".s")]
	}
	
	secret, _ := base64.StdEncoding.DecodeString(clientSecret)
	
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(message))
	
	expectedSignature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	
	if (expectedSignature != signature) {
		fmt.Printf("Signatures did NOT match! [Expected: %s] [Actual: %s]", expectedSignature, signature)
	}
	
	return true
}