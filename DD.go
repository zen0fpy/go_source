package go_source

import (
	"io/ioutil"
	"os"
)

func ProcessInput(f *os.File) error {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	return processBytes(b)
}
