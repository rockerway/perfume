package database

import (
	"database/sql"
	"perfume/packages/exception"

	_ "github.com/go-sql-driver/mysql"
)

const dbUsername string = "root"
const dbPassword string = "password"
const dbName string = "go"
const dbHost string = "172.18.0.2"

// Model struct
type Model struct {
	db *sql.DB
}

// Init is the method that connect to target
func (model *Model) Init() *sql.DB {
	var err error
	model.db, err = sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":3306)/"+dbName)
	recoder.Write(err)
	recoder.Write(model.db.Ping())
	return model.db
}

// Query is the method that select statement
func (model *Model) Query(statement string) *sql.Rows {
	rows, err := model.db.Query(statement)
	recoder.Write(err)
	return rows
}

// InitModel to init model
func InitModel() *Model {
	return &Model{db: &sql.DB{}}
}

func (model *Model) Close() {
	model.db.Close()
}
