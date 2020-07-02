package yamlcustom

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// ConfigSSH struct for yaml SSH config
type ConfigSSH struct {
	UserID     string `yaml:"userid"`
	PrivateKey string `yaml:"privatekey"`
	PrivateKeyProd string `yaml:"privatekeyprod"`
}

type Config struct {
	Conf []ConfigSSH `yaml:"conf"`
}

// ParseYAML parse yaml config file
func ParseYAML() Config {
	filename, _ := filepath.Abs("config.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}