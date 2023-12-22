package main

import "fmt"

type DBConnectionInterface interface {
	connect() string
}

type MySQLConnection struct{}

func (mc *MySQLConnection) connect() string {
	return "connect database"
}

type PasswordReminder struct {
	dbConnection DBConnectionInterface
}

func NewPasswordReminder(dbConnection DBConnectionInterface) *PasswordReminder {
	return &PasswordReminder{dbConnection: dbConnection}
}

func main() {
	mysqlConnection := &MySQLConnection{}
	passwordReminder := NewPasswordReminder(mysqlConnection)
	fmt.Println(passwordReminder.dbConnection.connect())
}
