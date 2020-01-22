package tools

import (
	"encoding/json"
	"io/ioutil"
)

// Read - read file and unmarshal to destination
func Read(filepath string, destination interface{}) error {

	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &destination)

	if err != nil {
		return err
	}

	return nil
}
