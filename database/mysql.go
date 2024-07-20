package database

var connection string

// Vid: 52_init <Initialization>
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
