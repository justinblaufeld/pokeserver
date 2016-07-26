package handlers

import (
	"github.com/golang/protobuf/proto"
	pp "github.com/pkmngo-odi/pogo-protos"
)

func GetPlayer(req *pp.Request) []byte {
	response := pp.GetPlayerResponse{
		Success: true,
	}
	data, _ := proto.Marshal(&response)
	return data
}
