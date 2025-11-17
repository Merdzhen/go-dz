package main

import (
	"bin/bins"
	"bin/storage"
)

func main() {
    myStorage := storage.NewFileStorage()
    
    bin := bins.NewBin("123", false, "test")
		binList := bins.NewBinList([]bins.Bin{*bin})
		
    myStorage.SaveBins(binList.Bins, "data.json")
}
