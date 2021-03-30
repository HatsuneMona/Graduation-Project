package main

import (
	"service/Router"
	"service/pkg/Utils"
)

func main() {
	Utils.InitLogger()
	defer Utils.Logger.Sync()
	Router.InitRouter()
}
