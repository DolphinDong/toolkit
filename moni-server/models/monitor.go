package models

import "gorm.io/gorm"

type HostPort struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;size:255"`
	Host string `json:"host" gorm:"not null;uniqueIndex:host_port;size:100"`
	Port int    `json:"port" gorm:"not null;uniqueIndex:host_port"`
}
