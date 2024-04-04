package database

type DBOperations interface {
	GetUsers() ([]string, error)
}

var operations DBOperations

func Operations() DBOperations {
	return operations
}
