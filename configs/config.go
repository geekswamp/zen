package configs

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"path/filepath"
	"runtime"
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
		Name string `mapstructure:"name"`
		Host string `mapstructure:"host"`
		Port uint32 `mapstructure:"port"`
	} `mapstructure:"app"`

	Postgres struct {
		Address  string `mapstructure:"address"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Port     uint32 `mapstructure:"port"`
		SSLMode  string `mapstructure:"sslmode"`
		Timezone string `mapstructure:"timezone"`

		Base struct {
			MaxOpenConn     int           `mapstructure:"max-open-conn"`
			MaxIdleConn     int           `mapstructure:"max-idle-conn"`
			ConnMaxLifeTime time.Duration `mapstructure:"conn-max-life-time"`
			ConnMaxIdleTime time.Duration `mapstructure:"conn-max-idle-time"`
		} `mapstructure:"base"`
	} `mapstructure:"postgres"`

	Password struct {
		Pepper string `mapstructure:"pepper"`

		Argon2 struct {
			Memory      uint32 `mapstructure:"memory"`
			Iterations  uint32 `mapstructure:"iterations"`
			Parallelism uint8  `mapstructure:"parallelism"`
			SaltLength  uint32 `mapstructure:"salt-length"`
			KeyLength   uint32 `mapstructure:"key-length"`
		} `mapstructure:"argon2"`
	} `mapstructure:"password"`

	JWT struct {
		PubKeyPath  string `mapstructure:"pub-key-path"`
		PrivKeyPath string `mapstructure:"priv-key-path"`
	} `mapstructure:"jwt"`
}

func init() {
	var reader io.Reader

	_, filename, _, _ := runtime.Caller(0)
	configPath := filepath.Dir(filename) + "/"
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
