package config

type Config struct {
	Debug      bool
	HostDB     string
	PortDB     int
	UserDB     string
	PasswordDB string
	NameDB     string
	SchemaName string
	TableName  string
}
