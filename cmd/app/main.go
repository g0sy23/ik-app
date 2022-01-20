package main

import "github.com/g0sy23/ik-app/internal/app"

const(
	configPath = "configs"
	configName = "main"
)

func main()	{
	ik_app.Run(configPath, configName)
}
