package singlton

import (
	"sync"
)

type SingleTon interface {
	AddOne() int
}

type singlton struct {
	count int
}

var instance *singlton

func GetInstance() *singlton {
	if instance == nil {
		instance = new(singlton)
	}
	return instance
}

func (s *singlton) AddOne() int {
	s.count += 1
	return s.count

}

type Config struct {
	name string
}

var configInstance *Config
var configOnce sync.Once

func GetConfigInstance() *Config {
	if configInstance == nil {
		configOnce.Do(func() {
			configInstance = &Config{
				name: "Surya",
			}
		})
	}
	return configInstance
}

func (c *Config) SetName(name string) {
	c.name = name
}

func (c *Config) GetName() string {
	return c.name
}
