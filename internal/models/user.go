package models

import (
	"database/sql"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func UserExists(sql *sql.DB, user string, pass string) (string, error) {
	var userId string
	query := "select id from users where username = $1 and pass = $2"
	rows, err := sql.Query(query, user, pass)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&userId); err != nil {
			log.Println(err)
			return "", err
		}
	}

	return userId, nil
}

func GetUserName(sql *sql.DB, userID string) (string, error) {
	var userName string
	query := "select username from users where id = $1"
	rows, err := sql.Query(query, userID)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&userName); err != nil {
			log.Println(err)
			return "", err
		}
	}

	return userName, nil
}
