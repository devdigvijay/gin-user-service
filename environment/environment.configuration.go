package environment

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	App      app      `yaml:"app"`
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
}

type app struct {
	Name string `yaml:"name" env:"APP_NAME" env-default:"main"`
	Env  string `yaml:"env" env:"APP_ENV" env-default:"dev"`
	Port int    `yaml:"port" env:"APP_PORT" env-default:"8080"`
}

type server struct {
	ReadTimeout     string `yaml:"readTimeout" env:"SERVER_READ_TIMEOUT" env-default:"2s"`
	WriteTimeout    string `yaml:"writeTimeout" env:"SERVER_WRITE_TIMEOUT" env-default:"5s"`
	ShutDownTimeout string `yaml:"shutDownTimeout" env:"SERVER_SHUTDOWN_TIMEOUT" env-default:"5s"`
}

type database struct {
	Host     string `yaml:"host" env:"DATABASE_HOST" env-default:"localhost"`
	Port     int    `yaml:"port" env:"DATABASE_PORT" env-default:"5432"`
	Name     string `yaml:"name" env:"DATABASE_NAME" env-default:"pgsql"`
	Username string `yaml:"username" env:"DATABASE_USERNAME" env-default:"admin"`
	Password string `yaml:"password" env:"DATABASE_PASSWORD" env-default:"admin"`
	SSL      bool   `yaml:"ssl" env:"DATABASE_SSL" env-default:"true"`
}

//go:embed configuration/*.yaml
var configuration embed.FS

func Load(env string) (*Configuration, error) {
	var cfg Configuration

	envFile := fmt.Sprintf("configuration/environment.%s.yaml", env)

	data, err := configuration.ReadFile(envFile)
	if err != nil {
		return nil, err
	}

	if err := cleanenv.ParseYAML(bytes.NewReader(data), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
