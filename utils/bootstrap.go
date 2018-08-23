package utils

func init() {

	//load configuration from file
	loadConfig()

	// Initialize private/public keys for JWT authentication
	initKeys()

	// Initialize Logger objects with Log Level
	setLogLevel(Level(AppConfig.LogLevel))

}
