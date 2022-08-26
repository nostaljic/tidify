package repository

const (
	DB_INFRA_ERROR        = 0
	RECORD_NOT_FOUND      = 1
	RECORD_ALREADY_EXISTS = 2
)

type DBError struct {
	ErrorCode    int
	ErrorMessage string
}

func CreateDBError(errorCode int, errorMessage string) *DBError {
	dbError := new(DBError)
	dbError.ErrorCode = errorCode
	dbError.ErrorMessage = errorMessage
	return dbError
}
