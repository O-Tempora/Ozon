package config

type Config struct {
	Host   string `yaml:"HOST"`
	Port   int    `yaml:"PORT"`
	DbPort int    `yaml:"DB_PORT"`
	DbName string `yaml:"DB_NAME"`
	DbUser string `yaml:"DB_USER"`
	DbPass string `yaml:"DB_PASS"`
}