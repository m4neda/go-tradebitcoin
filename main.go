package main

import (
	"fmt"
	"go-tradebitcoin/config"
	"go-tradebitcoin/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
