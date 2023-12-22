package main

import "fmt"

type MySQLConnection struct{}

func (mc *MySQLConnection) connect() string {
	return "Database connection"
}

type PasswordReminder struct {
	dbConnection *MySQLConnection
}

func NewPasswordReminder(dbConnection *MySQLConnection) *PasswordReminder {
	return &PasswordReminder{dbConnection: dbConnection}
}

func main() {
	mysqlConnection := &MySQLConnection{}

	passwordReminder := NewPasswordReminder(mysqlConnection)

	fmt.Println(passwordReminder.dbConnection.connect())
}
