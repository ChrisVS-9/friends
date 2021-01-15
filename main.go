package main

import (
	"log"
	"github.com/ChrisVS-9/friends/handlers"
	"github.com/ChrisVS-9/friends/bd"
)
	
func main() {
	if bd.ChequeoConnection()==0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}