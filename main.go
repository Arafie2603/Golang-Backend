package main

import (
	"finpro-golang2/database"
	"finpro-golang2/router"
)

func main() {
	// Menghubungkan ke database
	database.ConnectDB()
	

	// Menginisialisasi rute
	r := router.SetupRouter()

	// Menjalankan server
	err := r.Run(":8081")
	if err != nil {
		panic(err)
	}
}
