package middleware

import(
  pp "github.com/pkmngo-odi/pogo-protos"
  "github.com/ur0/pokeserver/models"
  "fmt"
)

type AuthError struct {}

func (e AuthError) Error() string {
  return "An authentication error occurred"
}

func GetPlayerFromRequestEnvelope(req *pp.RequestEnvelope) (*models.Player) {
  // A RequestEnvelope either contains a AuthInfo message or an AuthTicket field
  authInfo := req.AuthInfo
  authTicket := req.AuthTicket
  if authInfo != nil {
    fmt.Println("Found auth info")
    // Authenticate by parsing and verifying the JWT
  } else {
    fmt.Println("No auth info")
    // Authenticate using AuthTicket
  }
  fmt.Println(authTicket)
  return &models.Player{}
}
