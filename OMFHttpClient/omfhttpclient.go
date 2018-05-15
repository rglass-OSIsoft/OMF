package omfhttpclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateOMFType( jsonTypeString string ){
	// STEP 1 - Create a Type
	fmt.Println("Sending type creation request");
	makeWebRequest(jsonTypeString, "type", "create");
}

func CreateOMFStream( jsonStreamString string ){
	// STEP 2 - Create a Stream/Container
	fmt.Println("Sending stream creation request");
	makeWebRequest(jsonStreamString, "container", "create");
}

func SendOMFData(jsonDataString string){
	// STEP 3 - Send data associated with a stream
	fmt.Println("Sending data values");
	makeWebRequest(jsonDataString, "data", "create");
}

// from an example at https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
func makeWebRequest(jsonToSend, omfMessageType, omfAction string) {
	url := "https://localhost:5000/edge/omf/tenants/default/namespaces/data"
	//fmt.Println("URL:>", url)
	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	var jsonStr = []byte(jsonToSend)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("omfversion", "1.0")
	req.Header.Set("messageformat", "json")
	req.Header.Set("compression", "none")
	req.Header.Set("producertoken", "LzBzgY1zPnxQ55TX6O4WEcj2i1lfL47fDQM57ectmmDjShQvIsQCbMKbKh4i7be")
	//req.Header.Set("X-Custom-Header", "myvalue")

	//add the headers that vary per request - "messagetype" can be "type", "container", or "data"
	//req.Header.Set("messagetype", "type")
	req.Header.Set("messagetype", omfMessageType)
	// "action" can be "create", "update", or "delete"
	req.Header.Set("action", omfAction)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}


