package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

/////////////////////////////////////////////////////////////
// structure

type DBManager struct {
	User         string
	Password     string
	Host         string
	DatabaseName string
	Port         int
	Db           *sql.DB
}

type TableDesc struct {
	Field        string
	Type         string
	Null         string
	Key          string
	Default      string
	Extra        string
}

//////////////////////////////////////////////////////////
// interface

type DMLStatement interface {
	Desc(tableName string) error
}

///////////////////////////////////////////////////////////
// function define

// make teh dsn string
func makeDsn(user string, password string, database string, host string, port int) string {
	return user + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + database + "?charset=utf8&parseTime=true"
}

// create db manager
func GetDBManager(user string, password string, database string, host string, port int) (*DBManager, error) {
	db, err := sql.Open("mysql", makeDsn(user, password, database, host, port))
	if err != nil {
		return nil, err
	}
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	dbManager := &DBManager{
		User:         user,
		Password:     password,
		Host:         host,
		DatabaseName: database,
		Port:         port,
		Db:           db,
	}
	return dbManager, nil
}

func (this *DBManager) Desc(tableName string) []TableDesc, error {
	var desc_myql_sql string = "desc " + tableName
	rows, err := this.Db.Query(sql)
	tableDesc := make([]TableDesc, 0)
	if err != nil {
		return tableDesc, err
	}
	defer rows.Close()
	for rows.Next() {
		var desc MySqlTableDesc
		err := rows.Scan(&desc.Field, &desc.Type, &desc.Null, &desc.Key, &desc.Default, &desc.Extra)
		if err != nil {
			return tableDesc, err
		}
		tableDesc = append(tableDesc, desc)
	}
	err = rows.Err()
	if err != nil {
		return tableDesc, err
	}
	return tableDesc, nil
}