package config

type DB struct {
	Enable   bool
	Host     string
	Post     int
	User     string
	Password string
	DBName   string
	SSLMode  bool
}
