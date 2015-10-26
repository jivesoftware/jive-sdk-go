package main

import ("fmt"; "net/http";"encoding/json"; "reflect"; "crypto/sha256")
// import "crypto/sha256" 

type Payload struct {
	ClientID string `json:"clientID"`
	Code string `json:"code"`
	Scope string `json:"scope"`
	TenantID string `json:"tenantID"`
	JiveSignatureURL string `json:"JiveSignatureURL"`
	ClientSecret string `json:"clientSecret,omitempty"`
	JiveSignature string `json:"jiveSignature"`
	JiveURL string `json:"jiveURL"`
	Timestamp string `json:"timestamp"`
}

func MarshalSelectFields(structa interface{},
    includeFields map[string]bool) (jsona []byte, status error) {
    value := reflect.ValueOf(structa)
    typa := reflect.TypeOf(structa)
    size := value.NumField()
    jsona = append(jsona, '{')
    for i := 0; i < size; i++ {
        structValue := value.Field(i)
        var fieldName string = typa.Field(i).Name
        if marshalledField, marshalStatus := json.Marshal((structValue).Interface()); marshalStatus != nil {
            return []byte{}, marshalStatus
        } else {
            if includeFields[fieldName] {
                jsona = append(jsona, '"')
                jsona = append(jsona, []byte(fieldName)...)
                jsona = append(jsona, '"')
                jsona = append(jsona, ':')
                jsona = append(jsona, (marshalledField)...)
                if i+1 != len(includeFields) {
                    jsona = append(jsona, ',')
                }
            }
        }
    }
    jsona = append(jsona, '}')
    return
}

func is_valid_registration_notification (payload Payload){
	// Needed in unregister when a clientSecret isn't sent
	var clientSecret string

	// Saving these for later
	jiveSignatureURL := payload.JiveSignatureURL
	jiveSignature := payload.JiveSignature
	
	fmt.Println(payload.ClientSecret)
	if clientSecret != ""{
		if payload.ClientSecret != ""{
			fmt.Println("Registration event with no clientSecret, Invalid payload")
			return
		}else{
			tempSecret := []byte(reflect.ValueOf(payload.ClientSecret))
			payload.ClientSecret = sha256.New(tempSecret)
			fmt.Println(payload.ClientSecret)
		}
	}
	
	
		// Takes off JiveSignature and creates a byte string to send
	if data, status := MarshalSelectFields(payload, map[string]bool{
		"ClientID": true,
		"Code" : true,
		"Scope" : true,
		"TenantID" : true,
		"JiveSignatureURL" : true,
		"ClientSecret" : true,
		"JiveSignature" : false,
		"JiveURL" : true,
		"Timestamp" : true}); status != nil {
        	println("error")
    } else {
		// fmt.Println(payload)
    }
	
}

func main (){
    http.HandleFunc("/register", func(rw http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		data := Payload{}
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println(err)
		}
			is_valid_registration_notification(data)
        })
 
    http.ListenAndServe(":8090", nil)
}