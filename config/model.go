package config

type Config struct {
	Server   Server
	Database Database
	Email    Email
	Redis    Redis
	Midtrans Midtrans
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

type Redis struct {
	Addres   string
	Password string
}

type Midtrans struct {
	Key    string
	IsProd bool
}
