package models

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	GoCredentials struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"gocredentials"`
}
