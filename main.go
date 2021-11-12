package main

import (
	"log"

	"github.com/Mickeyrr/twittor/bd"
	"github.com/Mickeyrr/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
