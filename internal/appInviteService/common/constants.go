package common

import "time"

var (
	SERVER_PORT       string
	ENABLE_PROFILING  bool
	DB_NAME           string
	DB_DRIVER         string
	MYSQL_USERNAME    string
	MYSQL_PASSWORD    string
	CONNECTION_STRING string
	RETRY_SLEEP       time.Duration

	ContentTypeKey   = "Content-Type"
	ContentTypeValue = "application/json; charset=UTF-8"

	GenerateTokenPath = "/service/generate"

	GetUserQuery = "SELECT * FROM users where username='%v' and password='%v';"
)

const (
	TOKEN_ERROR = "Failed to generate token, please try again later."
)
