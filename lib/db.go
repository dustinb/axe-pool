package lib

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Bitaxe struct {
	gorm.Model
	MacAddr             string `json:"macAddr"`
	IP                  string
	Hostname            string `json:"hostname"`
	StratumURL          string `json:"stratumURL"`
	StratumPort         int    `json:"stratumPort"`
	StratumUser         string `json:"stratumUser"`
	FallbackStratumURL  string `json:"fallbackStratumURL"`
	FallbackStratumPort int    `json:"fallbackStratumPort"`
	FallbackStratumUser string `json:"fallbackStratumUser"`
}

type Pool struct {
	gorm.Model
	Host     string
	Port     int
	User     string
	Password string
}

var Database *gorm.DB

func Init() {
	var err error
	Database, err = gorm.Open(sqlite.Open("axe-pool.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	Database.AutoMigrate(&Bitaxe{})
	Database.AutoMigrate(&Pool{})
}
