package config

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

const (
	Production = "production"
)

func BootStrap(env string, property ...Property) iris.Configurator {

	envFile := ".env.yaml"
	if len(env) > 0 {
		envFile = fmt.Sprintf("env.%s.yaml", env)
	}

	c, err := parseYAML(envFile)
	if err != nil {
		panic(err)
	}

	config := c.Iris
	config.Other = make(map[string]interface{})

	c.Register(&config.Other)

	for _, p := range property {
		properties := p(envFile)
		for k, v := range properties {
			config.Other[k] = v
		}
	}

	return iris.WithConfiguration(config)
}


func ReadYAML(filename string) ([]byte, error) {
	yamlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	// read the raw contents of the file
	data, err := ioutil.ReadFile(yamlAbsPath)
	if err != nil {
		return nil, err
	}

	return data, nil
}


func parseYAML(filename string) (*Env, error) {

	c := DefaultEnv()

	data, err := ReadYAML(filename)
	if err != nil {
		return &c, fmt.Errorf("parse yaml: %w", err)
	}

	// put the file's contents as yaml to the default configuration(c)
	if err := yaml.Unmarshal(data, &c); err != nil {
		return &c, fmt.Errorf("parse yaml: %w", err)
	}
	return &c, nil
}
