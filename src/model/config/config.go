package config

import (
	"time"
)

// Public Public配置
type Public struct {
	Server Server `yaml:"Server"`
	Log    Log    `yaml:"Log"`
	Rule   Rule   `yaml:"Rule"`
	App    App    `yaml:"App"`
	Limit  Limit  `yaml:"Limit"`
}

// Private Private配置
type Private struct {
	Postgresql Postgresql `yaml:"Postgresql"`
}

type Postgresql struct {
	DriverName string `yaml:"DriverName"`
	SourceName string `yaml:"SourceName"`
}

type Rule struct {
	UsernameLenMax int `yaml:"UsernameLenMax"`
	UsernameLenMin int `yaml:"UsernameLenMin"`
}

type Server struct {
	RunMode               string        `yaml:"RunMode"`
	Address               string        `yaml:"Address"`
	ReadTimeout           time.Duration `yaml:"ReadTimeout"`
	WriteTimeout          time.Duration `yaml:"WriteTimeout"`
	DefaultContextTimeout time.Duration `yaml:"DefaultContextTimeout"`
}

type Log struct {
	Level         string `yaml:"Level"`
	LogSavePath   string `yaml:"LogSavePath"`
	LowLevelFile  string `yaml:"LowLevelFile"`
	LogFileExt    string `yaml:"LogFileExt"`
	HighLevelFile string `yaml:"HighLevelFile"`
	MaxSize       int    `yaml:"MaxSize"`
	MaxAge        int    `yaml:"MaxAge"`
	MaxBackups    int    `yaml:"MaxBackups"`
	Compress      bool   `yaml:"Compress"`
}

type App struct {
	Name      string `yaml:"Name"`
	Version   string `yaml:"Version"`
	StartTime string `yaml:"StartTime"`
	MachineID int64  `yaml:"MachineID"`
}

type Limit struct {
	IPLimit IPLimit `yaml:"IPLimit"`
}

type IPLimit struct {
	Cap     int64 `yaml:"Cap"`
	GenNum  int64 `yaml:"GenNum"`
	GenTime int64 `yaml:"GenTime"`
	Cost    int64 `yaml:"Cost"`
}
