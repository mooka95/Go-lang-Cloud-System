package main

import (
	"CloudSystem/database"
	"CloudSystem/routes"
	"fmt"
)

// var result = make(map[string]Word)

// var totalFrequencyInAllFiles = 0

func main() {
	fmt.Println("Project start")
	// searchWords([]string{"Qmeu", "LD"})
	// fmt.Println(result)
	// Initialize the database connection pool
	database.Init()

	router := routes.RegisterRoutes()
	router.Run(":8080")
}
