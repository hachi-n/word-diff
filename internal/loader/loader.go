package loader

import (
	"fmt"
	"io/ioutil"
	"os"
)

func LoadWithExitOnFail(fileName string) []byte {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return b
}
