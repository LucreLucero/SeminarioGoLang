package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

//Db config
type DbConfig struct {
	Type   string `yaml:"type"` //toma la decision de como levantar la db
	Driver string `yaml:"driver"` //driver para levantar mi database
	Conn   string `yaml:"conn"`   //connection string
}

//Config
type Config struct {
	DB      DbConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

//LoadConfig
func LoadConfig(filename string) (*Config, error) { //deberia devolver un config
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var c = &Config{}             //es un puntero a config
	err = yaml.Unmarshal(file, c) //si no hubo un error lo tengo que parsear para meterlo en el yaml---pkg.go.dev
	if err != nil {
		return nil, err
	}
	return c, nil
}
