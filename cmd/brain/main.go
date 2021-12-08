package main

import (
	"fmt"

	"github.com/m-nny/coco-verse/internal/app"
	"github.com/m-nny/coco-verse/internal/server"
)

func main() {
	appConfig := app.LoadAppConfig()

	router := server.NewRouter()

	address := fmt.Sprintf("%s:%d", appConfig.Server.Host, appConfig.Server.Port)

	router.Run(address)

	fmt.Printf("Server is listening on %v\n", address)
}
