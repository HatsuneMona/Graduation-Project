package main

import "service/Utils/log"

func main() {
	log.InitLogger()
	defer log.Logger.Sync()
}
