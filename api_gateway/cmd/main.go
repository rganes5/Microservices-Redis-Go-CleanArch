package main

import (
	"api_gateway/pkg/config"
	"api_gateway/pkg/di"
	"log"
)

//	@title			X-TENTIONCREW
//	@version		2.0
//	@description	MICROSERVICE API BUILD USING GOLANG, REDIS, POSTGRESSQL, REST API following Clean Architecture.

//	@contact
// name: Ganesh
// url: https://github.com/rganes5
// email: ganeshraveendranit@gmail.com

//	@license
// name: MIT
// url: https://opensource.org/licenses/MIT

//	@host	localhost:3000

// @Basepath	/
// @Accept		json
// @Produce	json
// @Router		/ [get]

func main() {
	//swag init -g cmd/main.go
	c, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("Failed to load config:", configerr)
	}
	server, dierr := di.InitializeAPI(c)
	if dierr != nil {
		log.Fatal("Failed to initialize server", dierr)
	}
	server.Start()
}
