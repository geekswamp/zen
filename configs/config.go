package configs

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/geekswamp/zen/pkg/env"
	"github.com/geekswamp/zen/pkg/file"
	"github.com/spf13/viper"
)

var (
	//go:embed config-dev.yaml
	devConfig []byte

	//go:embed config-pro.yaml
	proConfig []byte

	config = new(Config)
)

type Config struct {
	App struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"app"`

	Postgres struct {
		Address  string `mapstructure:"address"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`

		Base struct {
			MaxOpenConn     int       `mapstructure:"max-open-conn"`
			MaxIdleConn     int       `mapstructure:"max-idle-conn"`
			ConnMaxLifeTime time.Time `mapstructure:"conn-max-life-time"`
			ConnMaxIdleTime time.Time `mapstructure:"conn-max-idle-time"`
		} `mapstructure:"base"`
	} `mapstructure:"postgres"`

	Argon2 struct {
		Pepper      string `mapstructure:"pepper"`
		Memory      int    `mapstructure:"memory"`
		Iterations  int    `mapstructure:"iterations"`
		Parallelism int    `mapstructure:"parallelism"`
	} `mapstructure:"argon2"`
}

func init() {
	var reader io.Reader

	configPath := "./configs"
	prefixFile := "config-"

	switch env.Active().Value() {
	case env.Pro:
		reader = bytes.NewReader(proConfig)
	default:
		reader = bytes.NewReader(devConfig)
	}

	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(reader); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.SetConfigName(prefixFile + env.Active().Value())
	viper.AddConfigPath(configPath)

	configFile := configPath + prefixFile + env.Active().Value() + ".yaml"
	_, ok := file.IsExist(configFile)
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(_ fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() Config {
	return *config
}
