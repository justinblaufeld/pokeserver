package handlers

import (
	"github.com/golang/protobuf/proto"
	pp "github.com/pkmngo-odi/pogo-protos"
	"github.com/ur0/pokeserver/models"
)

func GetPlayer(req *pp.Request, player *models.Player) []byte {
	response := pp.GetPlayerResponse{
		Success: true,
	}
	data, _ := proto.Marshal(&response)
	return data
}
