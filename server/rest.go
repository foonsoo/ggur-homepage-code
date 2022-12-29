package server

func GetUsers() {
	db := GetConnector()
	err := db.Ping()
}
