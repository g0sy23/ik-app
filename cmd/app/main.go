package main

import
(
	"github.com/g0sy23/ik-app/internal/app"
	// swagger "github.com/arsmn/fiber-swagger/v2"
	// _ "github.com/g0sy23/ik-app/docs"
)

const(
	configPath = "configs"
	configName = "main"
)

// @title           IK App API
// @version         1.0
// @description     API server for Irina Kairatovna web application.

// @host      localhost:8080
// @BasePath  /
func main()	{
	ik_app.Run(configPath, configName)
}
