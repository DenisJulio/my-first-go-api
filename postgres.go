package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "denis"
	password = "password"
	dbname   = "myDb"
)

type DBClient struct {
	DB *sql.DB
}

type Message struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	Language string `json:"language"`
}

func (d *DBClient) Close() error {
	return d.DB.Close()
}

func (d *DBClient) GetAllMessages(lang string) ([]Message, error) {
	var rows *sql.Rows
    var err error

    if lang != "" {
        rows, err = d.DB.Query("SELECT id, message, language FROM messages WHERE language ILIKE $1", lang)
    } else {
        rows, err = d.DB.Query("SELECT id, message, language FROM messages")
    }

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.Message, &msg.Language)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func connectToDB() *DBClient {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database!")
	return &DBClient{DB: db}
}


