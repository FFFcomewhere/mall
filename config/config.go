package config

import "time"

type Mysql struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	TestName    string
}

type Config struct {
	Mysql  Mysql
	Server Server
}

type Server struct {
	RunMode      string
	Host         string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
