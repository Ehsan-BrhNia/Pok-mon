package apiclients

import (
	"encoding/json"
	"log"
	"net/http"
)

func Api(link string, Address interface{}) {
	response, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(response.Body).Decode(Address)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
}







