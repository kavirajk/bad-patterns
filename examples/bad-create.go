package main

import (
	"fmt"
	"log"

	"github.com/kavirajk/bad-patterns/models"
)

func main() {
	models.InitModel()
	p := models.Picture{AlbumId: 99, Caption: "Learning a lot!! #golang"}
	if err := p.Save(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}
