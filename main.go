package main

import (
	"pkg/configs"
	"pkg/routes"

	"github.com/gin-gonic/gin"
)

// Just MVC Pattern -> Should Change Repository Pattern Later....
// Model Should Add After DB Connected
func main() {
	r := gin.Default()
	routes.SetupRouter(r)
	configs.ConnectDB()
	configs.ConnectDB2()
	r.Run(":80")
}
