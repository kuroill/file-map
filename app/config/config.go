package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	password string
	rootDir  string
)

func GetConfig() Config {
	flag.Parse()

	if rootDir == "" {
		log.Fatal("root_dir is required. Please provide it using the -root flag.")
	}

	path := "../app/config"
	if os.Getenv("GIN_MODE") == "release" {
		path = "/app"
	}

	file, err := os.Open(filepath.Join(path, "config.yaml"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	c := Config{}
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	if rootDir != "" {
		c.User.RootDir = rootDir
	}
	if password != "" {
		c.User.Password = password
	}
	return c
}

func init() {
	flag.StringVar(&password, "password", "", "password, eg: -password your_password")
	flag.StringVar(&rootDir, "root", "", "root directory, eg: -root /path/to/root_dir")
}

type Config struct {
	Server Server `yaml:"server"`
	User   User   `yaml:"user"`
}

type Server struct {
	HTTP HTTP `yaml:"http"`
}

type HTTP struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type User struct {
	SecretKey string `yaml:"secret_key"`
	Password  string `yaml:"password"`
	RootDir   string `yaml:"root_dir"`
	Mnt       string `yaml:"mnt"`
}
