package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Creating a GET Request
	res, err := http.Get("Type in the URL of the api")
	//handle the error
	if err != nil {
		log.Fatal(err)
	}

	//read the body of the response from the API
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", body)

	//Creating a POST request
	pbody, _ := json.Marshal(map[string]string{
		"name":  "Aniket",
		"email": "ani@xyz.com",
	})
	resBody := bytes.NewBuffer(pbody)
	res, err = http.Post("https://postman-echo.com/post", "application/json", resBody)

	//Handling Error
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}