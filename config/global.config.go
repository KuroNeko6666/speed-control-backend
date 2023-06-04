package config

const (
	ServerHost = "localhost"
	ServerPort = ":8000"

	DatabaseHost = "localhost"
	DatabasePort = ":3306"
	DatabaseUser = "root"
	DatabasePass = "mareca"
	DatabaseName = "speed_control_database"

	StorageHost = ""
	StoragePort = ""

	SecretKeyApp = "SP3EDC0ntR01"
)

func DatabaseDSN() string {
	option := "?charset=utf8mb4&parseTime=True&loc=Local"
	return DatabaseUser + ":" + DatabasePass + "@tcp(" + DatabaseHost + DatabasePort + ")/" + DatabaseName + option
}

func ServerAddress() string {
	return ServerHost + ServerPort
}
