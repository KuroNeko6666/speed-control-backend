package config

const (
	ServerHost = "0.0.0.0"
	ServerPort = ":8000"

	DatabaseHost = "178.128.209.164"
	DatabasePort = ":3306"
	DatabaseUser = "dikdik"
	DatabasePass = "Kurniawan757!"
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
