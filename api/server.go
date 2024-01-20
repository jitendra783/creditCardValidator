package api

import (
	"fmt"
	"log"
)

func Init() {
	port := "8080"
	fmt.Println("port ", port)

	route := Routes()
	log.Println("listening and serve !")
	route.Run(":8081")
	log.Println("listening and serve !")
}
