package config

type Config struct {
	Logger `json:"logger"`
	Tables `json:"tables"`
}

type Tables struct {
	CountUsers int    `json:"countUsers" default:"1"`
	TablePath1 string `json:"tablePath1"`
	TablePath2 string `json:"tablePath2"`
}

// Logger содержит параметры логгера
type Logger struct {
	LogLevel        string `mapstructure:"logLevel"`
	LogFileEnable   bool   `mapstructure:"logFileEnable"`
	LogStdoutEnable bool   `mapstructure:"logStdoutEnable"`
	LogPath         string `mapstructure:"logPath"`
	MaxSize         int    `mapstructure:"maxSize"`
	MaxAge          int    `mapstructure:"maxAge"`
	MaxBackups      int    `mapstructure:"maxBackups"`
	RewriteLog      bool   `mapstructure:"rewriteLog"`
}
