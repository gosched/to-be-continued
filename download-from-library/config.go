package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	// config := &Config{MaxGoroutines: 3}
	// config.Store()

	// config := &Config{}
	// config.Load("config.json")
	// fmt.Printf("%+v\n", config)
}

// Config .
type Config struct {
	MaxGoroutines int `json:"maxGoroutines"` // maximum number of goroutines
}

// Load .
func (c *Config) Load(name string) error {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}

// Store .
func (c *Config) Store(name string) {
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}
