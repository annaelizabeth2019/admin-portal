package config

var config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Port                       string
	Secret                     string
	AccessTokenExpireDuration  int
	RefreshTokenExpireDuration int
	LimitCountPerRequest       float64
}

type DatabaseConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

func Setup() error {
	configuration := Configuration{
		Server:   getServerConfig(),
		Database: getDbConfig(),
	}
	config = &configuration
	return nil
}

func getServerConfig() ServerConfiguration {
	return ServerConfiguration{
		// TBD
	}
}

func getDbConfig() DatabaseConfiguration {
	return DatabaseConfiguration{
		// TBD
	}
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	if config != nil {
		return config
	}
	err := Setup()
	if err != nil {
		panic(err)
	}
	return config
}
