package main

import (
	"fmt"
	"go-tradebitcoin/bitflyer"
	"go-tradebitcoin/config"
	"go-tradebitcoin/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
