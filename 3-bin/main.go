package main

import (
	"bin/api"
	"bin/bins"
	"bin/config"
	"bin/storage"
	"fmt"
)

func main() {
	myStorage := storage.NewFileStorage()

	bin := bins.NewBin("123", false, "test")
	binList := bins.NewBinList([]bins.Bin{*bin})

	err := myStorage.SaveBins(binList.Bins, "bins.json")
	if err != nil {
		fmt.Println("Error:", err)
	}

	cfg := config.GetConfig()
	client := api.NewClient(cfg.Key)
	fmt.Println(client)
}
