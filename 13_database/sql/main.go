package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	envs, err := SourceEnv()
	if err != nil {
		log.Fatalln(err)
	}
	//"postgresql://username:password@localhost:5432/database_name"
	v := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v",
		envs["USER"], envs["PASSWORD"], envs["HOST"], envs["PORT"], envs["DATABASE"])
	conn, err := pgxpool.New(context.Background(), returnUrl(v))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	Lines(conn)
}

func Line(c *pgxpool.Pool) {
	var (
		id    string
		name  string
		emoji string
	)
	err := c.QueryRow(context.Background(), "select id, name, emoji from animal;").Scan(&id, &name, &emoji)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id, name, emoji)
}

func Lines(c *pgxpool.Pool) {
	var (
		id    string
		name  string
		emoji string
	)
	rows, err := c.Query(context.Background(), "select id, name, emoji from animal;")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&id, &name, &emoji)
		fmt.Println(id, name, emoji)
	}

	if rows.Err() != nil {
		log.Fatalln(rows.Err())
	}

	fmt.Println("All Rows")
}

func returnUrl(url string) string {
	urlBytes := []byte(url)
	b := make([]byte, len(urlBytes))
	i := 0
	for _, v := range urlBytes {
		if v != 0 {
			b[i] = v
			i++
		}
	}

	return string(b[:bytes.IndexByte(b, 0)])
}
