package config

type Config struct {
	Server   Server
	Database Database
	Email    Email
}

type Server struct {
	Host string
	Port string
}
type Database struct {
	Host     string
	Password string
	Name     string
	User     string
	Port     string
}

type Email struct {
	Host     string
	Port     string
	Email    string
	Password string
}
