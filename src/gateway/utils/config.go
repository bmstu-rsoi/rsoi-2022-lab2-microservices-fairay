package utils

type Configuration struct {
	LogFile         string `json:"log_file"`
	Port            uint16 `json:"port"`
	FlightsEndpoint string `json:"flights-endpoint"`
	PrivilegesEndpoint string `json:"privileges-endpoint"`
}

var (
	Config Configuration
)

// TODO: returnable error
func InitConfig() {
	Config = Configuration{
		"logs/server.log",
		8080,

		"http://flights:8080",
		"http://privileges:8080",
	}
}
