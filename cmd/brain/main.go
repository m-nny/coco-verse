package main

import (
	"fmt"

	"github.com/m-nny/coco-verse/internal/app"
	"github.com/m-nny/coco-verse/internal/server"
)

func main() {
	context := app.NewContext()

	router := server.NewRouter(context)

	address := context.Config.Server.GetBindAddress()

	router.Run(address)

	fmt.Printf("Server is listening on %v\n", address)
}
