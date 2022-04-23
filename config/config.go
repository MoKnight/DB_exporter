/*
 * @Description:
 * @version:
 * @Author: MoonKnight
 * @Date: 2022-02-17 10:23:25
 * @LastEditors: MoonKnight
 * @LastEditTime: 2022-02-21 16:06:29
 */

package config

import (
	"fmt"
	"io/ioutil"

	"github.com/prometheus/common/log"
	"gopkg.in/yaml.v2"
)

// Load attempts to parse the given config file and return a Config object.
func Load(configFile string) (*Config, error) {
	log.Infof("Loading configuration from %s", configFile)
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	c := Config{ConfigFile: configFile}
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return nil, err
	}

	fmt.Println(c.DSN())

	return &c, nil
}

type Config struct {
	Dsn        string   `yaml:"data_sourse_name"`
	Querys     []string `yaml:"queries"`
	ConfigFile string   //file name
}

func (c *Config) DSN() string {
	return c.Dsn
}

func (c *Config) QUERYS() []string {
	return c.Querys
}
