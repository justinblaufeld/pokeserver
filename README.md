# PokeServer

This is an open-source Pokemon GO server, written in Golang.

## Setting up
Fetch dependencies using `govendor sync`. The server, by default, runs on port `8080`.

## How does this work?
An incoming request is first unmarshaled from ProtoBufs. A handler function is selected from the `handlers` package and is used to handle the response.

## Writing Handlers
Right now, only a few handlers are implemented. Here's how to write your own.

1. Create your file in the `handlers` directory, named `your_request.go` where YourRequest is the name of the request type, as seen in `PokeProtos.NetworkingEnvelopes.RequestEnvelope`. For example, a handler for the `GetPlayer` request would be named `handlers/get_player.go`.
2. Declare a function named `YourRequest` in the file you just created. It must accept a pointer to `PokeProtos.request` and return a `[]byte`, which should be the response, marshalled into ProtoBufs using the correct buffer for your request.
3. Add your function to the `switch-case` statement in `server.go`, take it's response into the `responses` byte slice.
4. That's it.

## License and legal stuff
This software is licensed under the MIT License, as specified in the `LICENSE` file in the root of this repository. This software DOES NOT redistribute copyrighted material or artwork. Pokemon is a registerted trademark of The Pokemon Company.
