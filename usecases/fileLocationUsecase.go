package usecases

import (
	"backend-rsmata-360/repositories"
	"errors"
)

func GetFilePathByName(name string) (string, error) {
	if name == "" {
		return "", errors.New("file location name is required")
	}

	fileLocation, err := repositories.GetFileLocationByName(name)

	if err != nil{
		return "", err
	}

	if fileLocation.Path == ""{
		return "", errors.New("file location path is empty")
	}

	return fileLocation.Path, nil
}