package main

import (
	"fmt"
	"log"

	"github.com/kavirajk/bad-patterns/models"
)

func main() {
	models.InitModel()
	albums, err := models.GetAllAlbums()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(albums)
}
