package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

//Config represents the Project configuration
type Config struct {
	Settings     SettingsType `json:"settings"`
	Environments Environments `json:"environments"`
}

type SettingsType struct {
	BuildDir string `json:"build_dir"`
	Filemode string `json:"filemode"`
}
type Variables map[string]string

type Environment struct {
	Variables `json:"Variables"`
}

type Environments map[string]Environment

func NewConfigFromFile(f string) (Config, error) {

	// open file
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		return Config{}, err
	}
	c, err := readConfigFromString(string(raw))
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

func readConfigFromString(s string) (Config, error) {
	cdec := json.NewDecoder(strings.NewReader(s))

	var config Config
	if err := cdec.Decode(&config); err != nil {
		return Config{}, err
	}

	// for _, t := range config.Environments {
	// 	for k, v := range t.Variables {
	// 		fmt.Printf("%s: %s \n", k, v)
	// 	}
	// }
	return config, nil
}
