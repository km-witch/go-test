package main

import (
	"flag"
	"pkg/configs"
	"pkg/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Just MVC Pattern -> Should Change Repository Pattern Later....
// Model Should Add After DB Connected
func main() {
	r := gin.Default()
	routes.SetupRouter(r)
	configs.ConnectDB()
	configs.ConnectDB2()

	// CORS
	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://127.0.0.1:8080", "https://dev2.witchworld.io/", "http://dev-go.witchworld.io", "https://dev-go.witchworld.io"},
			AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
			AllowHeaders:     []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	))

	// PORT Setting
	PORT := flag.String("port", "8080", "Write Your PORT")
	flag.Parse()
	PORT_RE := ":" + *PORT

	r.Run(PORT_RE)
}
