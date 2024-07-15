package config

type Config struct {
	Port int

	MySQL MySQL
}

type MySQL struct {
	Username string
	Password string
	Address  string
	DBName   string
}
