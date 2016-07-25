package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	ne "github.com/pkmngo-odi/pogo-protos/networking_envelopes"
	"io/ioutil"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	requestEnvelope := &ne.RequestEnvelope{}

	body, _ := ioutil.ReadAll(r.Body)

	proto.Unmarshal(body, requestEnvelope)

	for _, request := range requestEnvelope.Requests {
		// A RequestEnvelope has a bunch of requests. We need to handle all of them, get their responses
		// and shoehorn them into a response, marshall them and send them.
		switch request.RequestType.String() {
		case "GET_PLAYER_PROFILE":
			// Something
		default:
			fmt.Println("Requests of type", request.RequestType.String(), "aren't implemented yet")
		}
	}
}

func RunServer() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
