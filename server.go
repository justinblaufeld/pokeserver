package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	pp "github.com/pkmngo-odi/pogo-protos"
	"github.com/ur0/pokeserver/handlers"
	"github.com/ur0/pokeserver/handlers/middleware"
	"io/ioutil"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	requestEnvelope := &pp.RequestEnvelope{}

	body, _ := ioutil.ReadAll(r.Body)

	proto.Unmarshal(body, requestEnvelope)
	// Get the request id to send back
	requestID := requestEnvelope.RequestId

	currentPlayer := middleware.GetPlayerFromRequestEnvelope(requestEnvelope)

	log.WithFields(log.Fields{"RequestID": requestID}).Info("Incoming request")

	// This holds the returns of all handler functions
	responses := make([][]byte, len(requestEnvelope.Requests))

	for i, request := range requestEnvelope.Requests {
		// A RequestEnvelope has a bunch of requests. We need to handle all of them, get their responses
		// and shoehorn them into a response, marshall them and send them.
		log.WithFields(log.Fields{"RequestID": requestID, "RequestType": request.RequestType.String()}).Debug("Handling request")
		switch request.RequestType.String() {
		case "GET_PLAYER":
			responses[i] = handlers.GetPlayer(request, currentPlayer)
		default:
			log.WithFields(log.Fields{"RequestID": requestID, "RequestType": request.RequestType.String()}).Error("Unsupported request type")
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
	log.WithFields(log.Fields{"RequestID": requestID}).Debug("Request complete")
}

func RunServer() {
	log.SetLevel(log.DebugLevel) // Maybe remove this for prod envs
	log.WithFields(log.Fields{"package": "main"}).Info("PokeServer starting")
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
