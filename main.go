package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"pkg/configs"
	"pkg/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Just MVC Pattern -> Should Change Repository Pattern Later....
// Model Should Add After DB Connected
func main() {
	// Initialzie Log
	now := time.Now().Format("2006_01_02")
	fpLog, err := os.OpenFile(filepath.Join("logs", now+".txt"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[LOG] ")
	log.SetOutput(io.MultiWriter(fpLog, os.Stdout))
	log.Println("Log Initialized")

	// Initialize gin
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

	log.Println("System Shutdown")
}
