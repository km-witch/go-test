package main

import (
	"flag"
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
	// PORT Setting
	PORT := flag.String("port", "8080", "Write Your PORT")
	flag.Parse()
	PORT_RE := ":" + *PORT

	r.Run(PORT_RE)
}
