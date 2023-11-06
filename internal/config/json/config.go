package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Server   Server     `json:"server"`
	Decstop  Decstop    `json:"Decstop"`
	Keyboard []Keyboard `json:"Keyboard"`
}

type Decstop struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
}

type Keyboard struct {
	Name  string `json:"Name"`
	Value []int  `json:"Value"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func New() Config {
	return Config{}
}

func ParseConfig(path string, c *Config) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("can't open config file: ", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}
}

func KeyboardSetToMap(keyboardSet []Keyboard) map[int]string {
	keyboardSetMap := make(map[int]string)

	for _, set := range keyboardSet {
		for _, value := range set.Value {
			keyboardSetMap[value] = set.Name
		}
	}

	return keyboardSetMap
}
