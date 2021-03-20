package main

import (
	"service/pkg/Utils"
)

func main() {
	Utils.InitLogger()
	defer Utils.Logger.Sync()
}
