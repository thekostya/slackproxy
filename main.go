package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("port", "80")
	viper.SetEnvPrefix("slackproxy")
	port := viper.GetString("port")
	host := viper.GetString("host")

	http.HandleFunc("/", handler)
	fmt.Printf("Listen: %s:%s\n", host, port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
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
