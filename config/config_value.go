package config

type AppConfig struct {
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}
