package utils

type Configuration struct {
	LogFile string `json:"log_file"`
	Port    uint16 `json:"port"`
}

var (
	Config Configuration
)

// TODO: returnable error
func InitConfig() {
	Config = Configuration{
		"logs/server.log",
		8080,
	}
}
