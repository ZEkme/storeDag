package main

import (
	"log"

	dag "github.com/Way-Flare/dagestan-backend"
)

func main() {
	srv := new(dag.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
