package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string     `yaml:"storage_path"`
	HTTPServer   `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath==""{
		// go run main.go --config=config.yaml
		flags:=flag.String("config","","path to the configuration file")
		// command-line flag:
		// when running argument is passed it is called flags.
		flag.Parse()
		configPath= *flags

		if configPath==""{
			log.Fatal("config path is not set")
		}
	}
		

		if _,err:=os.Stat(configPath);os.IsNotExist(err){
			log.Fatalf("config file does not exist: %s",configPath)
		}

		var cfg Config
		err:=cleanenv.ReadConfig(configPath,&cfg)
// Reads the config file,

// Converts the YAML data into Go variables,

// Stores everything inside cfg.
		if err!= nil{
			log.Fatalf("can not read config file: %s",err.Error())
		}

		return &cfg
	}

