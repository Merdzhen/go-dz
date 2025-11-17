package storage

import (
	"bin/bins"
	"encoding/json"
	"fmt"
	"os"
)

type Storage interface {
	SaveBins(bins []bins.Bin, fileName string) error
	ReadBins(fileName string) ([]bins.Bin, error)
}

type FileStorage struct{}

func NewFileStorage() Storage {
	return &FileStorage{}
}

func (fs *FileStorage) SaveBins(bins []bins.Bin, fileName string) error {
	data, err := json.Marshal(bins)
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileStorage) ReadBins(fileName string) ([]bins.Bin, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var bins []bins.Bin
	err = json.Unmarshal(data, &bins)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bins, nil
}
