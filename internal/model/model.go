package model

import "time"

type User struct {
	ID       string `json:"-" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"uniqueIndex;not null"`
	Token    string `json:"token" gorm:"-"`
	IsActive bool   `json:"is_active"`
}

type App struct {
	ID         string `json:"-" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"uniqueIndex;not null"`
	UserID     string `json:"-" gorm:"index;not null"`
	Token      string `json:"token" gorm:"-"`
	IsActive   bool   `json:"is_active"` // NOTICE: can't be redeclared in environment from "inactive" to "active"
	IsVersions bool   `json:"is_versions"`
	PubKey     string `json:"-"`
}

type Env struct {
	ID         string `gorm:"primaryKey"`
	Name       string
	AppID      string
	IsActive   int8
	IsVersions int8
	PubKey     string
}

type Config struct {
	ID        string `gorm:"primaryKey"`
	AppID     string
	EnvID     string
	Format    string
	Data      string
	Version   uint
	CreatedAt time.Time
}
