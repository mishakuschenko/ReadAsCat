package lib

import (
	"io/ioutil"
)

func ReadFile(path string) string {
	cont, err := ioutil.ReadFile(path)
	if err != nil {
		print("file not found!\n")
	}
	return string(cont)
} 


