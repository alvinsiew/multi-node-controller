package yamlcustom

import (
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// ConfigSSH struct for yaml SSH config
type ConfigSSH struct {
	UserID         string `yaml:"userid"`
	PrivateKey     string `yaml:"privatekey"`
	PrivateKeyProd string `yaml:"privatekeyprod"`
}

// Config struct for mnc config
type Config struct {
	Conf []ConfigSSH `yaml:"conf"`
}

// ParseYAML parse yaml config file
func ParseYAML() Config {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	filename, _ := filepath.Abs(usr.HomeDir + "/.aws/config.yml")
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
