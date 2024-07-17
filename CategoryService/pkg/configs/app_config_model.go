package configs

import (
	Logger "github.com/ewinjuman/go-lib/logger"
)

type Configuration struct {
	Apps   Apps
	Logger Logger.Options
}

type Apps struct {
	Name     string
	HttpPort int
	GrpcPort int
}
