package repository

import (
	"src/db"
	"strings"
)

func GetIntParameter(key string) (int, error) {
	sql := `select value_int from od_parameter where name=?`
	return db.GetSingleInt(db.Get().QueryRow(sql, strings.ToLower(key)))
}

func GetIntParameterOrZero(key string) int {
	v, _ := GetIntParameter(key)
	return v
}

func GetStrParameter(key string) (string, error) {
	sql := `select value_int from od_parameter where name=?`
	return db.GetSingleString(db.Get().QueryRow(sql, strings.ToLower(key)))
}

func GetStrParameterOrEmpty(key string) string {
	v, _ := GetStrParameter(key)
	return v
}
