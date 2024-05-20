package main
import(
	"fmt"
	"CloudSystem/database"
	"CloudSystem/routes"
)


func main(){
fmt.Println("Project start")
    // Initialize the database connection pool
    database.Init()

	router := routes.RegisterRoutes()
	router.Run(":8080") 
}