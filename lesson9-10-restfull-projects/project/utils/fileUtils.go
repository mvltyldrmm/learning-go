package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(fileName string)(string, error){

	if IsEmpty(fileName) {
		return "", errors.New("File name is empty")
	}

	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	return string(data), nil
}