package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pp "github.com/pkmngo-odi/pogo-protos"
	"github.com/ur0/pokeserver/handlers"
	"io/ioutil"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	requestEnvelope := &pp.RequestEnvelope{}

	body, _ := ioutil.ReadAll(r.Body)

	proto.Unmarshal(body, requestEnvelope)
	// Get the request id to send back
	requestID := requestEnvelope.RequestId

	// This holds the returns of all handler functions
	responses := make([][]byte, len(requestEnvelope.Requests))

	for i, request := range requestEnvelope.Requests {
		// A RequestEnvelope has a bunch of requests. We need to handle all of them, get their responses
		// and shoehorn them into a response, marshall them and send them.
		switch request.RequestType.String() {
		case "GET_PLAYER":
			responses[i] = handlers.GetPlayer(request)
		default:
			fmt.Println("Requests of type", request.RequestType.String(), "aren't implemented yet")
		}
	}

	responseEnvelope := pp.ResponseEnvelope{
		StatusCode: 0,
		RequestId:  requestID,
		Returns:    responses,
	}

	response, _ := proto.Marshal(&responseEnvelope)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(response[:])
}

func RunServer() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
