package main

import (
	"bin/bins"
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
}
