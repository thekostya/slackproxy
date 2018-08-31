package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Listen: :8808\n")
	err := http.ListenAndServe(":8808", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("New Request\n")
	bodyStr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		bodyStr = []byte("something happens")
	}
	defer r.Body.Close()

	fmt.Printf("With body: %s\n", bodyStr)
	data := make(map[string]string)
	data["text"] = string(bodyStr)
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", "https://hooks.slack.com/services/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", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	client := &http.Client{}
	client.Do(req)
}
