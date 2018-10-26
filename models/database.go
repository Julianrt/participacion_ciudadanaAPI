package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Julianrt/participacion_ciudadana/config"
	//driver mysql
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var debug bool

func init() {
	CreateConnection()
	debug = config.GetDebug()
}

//CreateConnection method
func CreateConnection() {

	if GetConnection() != nil {
		return
	}

	url := config.GetURLDatabase()
	if connection, err := sql.Open("mysql", url); err != nil {
		panic(err)
	} else {
		db = connection
	}
}

//CreateTables method
func CreateTables() {
	createTable("ciudad", ciudadScheme)
}

func createTable(tableName, scheme string) {
	if !existsTable(tableName) {
		Exec(scheme)
	} else {
		truncateTable(tableName)
	}
}

func truncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

func existsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, _ := Query(sql)
	return rows.Next()
}

//Exec method
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil && !debug {
		log.Println(err)
	}
	return result, err
}

//Query method
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil && !debug {
		log.Println(err)
	}
	return rows, err
}

//InsertData method
func InsertData(query string, args ...interface{}) (int64, error) {
	result, err := Exec(query, args...)
	if err != nil {
		return int64(0), err
	}
	id, err := result.LastInsertId()
	return id, err
}

//GetConnection method
func GetConnection() *sql.DB {
	return db
}

//Ping method
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//CloseConnection method
func CloseConnection() {
	db.Close()
}
