package db

import "database/sql"

func GetSingleInt(row *sql.Row) (int, error) {
	var val int
	if err := row.Scan(&val); err != nil {
		return 0, err
	}
	return val, nil
}

func GetSingleString(row *sql.Row) (string, error) {
	var val string
	if err := row.Scan(&val); err != nil {
		return "", err
	}
	return val, nil
}
