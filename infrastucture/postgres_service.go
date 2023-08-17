package infrastucture

import (
	"database/sql"
	"errors"
	"log"
)

type DatabaseService struct {
	DB *sql.DB
}

func NewDatabaseService(connectionString string) *DatabaseService {
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatalf("Error while establishing database connection : %s \n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error while pinging database : %s \n", err)
	}
	return &DatabaseService{
		DB: db,
	}
}

func NewDatabaseServiceWithDriverName(connectionString string, driverName string) *DatabaseService {
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		log.Fatalf("Error while establishing database connection : %s \n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error while pinging database : %s \n", err)
	}
	return &DatabaseService{
		DB: db,
	}
}

func (d *DatabaseService) CallStoredProcedure(procedureName string, params ...interface{}) (*string, *string) {
	var result sql.NullString
	err := d.DB.QueryRow(procedureName, params...).Scan(&result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, nil
		}
	}

	if result.Valid {
		return &result.String, nil
	} else {
		return nil, nil
	}
}
