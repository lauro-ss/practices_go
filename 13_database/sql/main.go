package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	envs, err := SourceEnv()
	if err != nil {
		log.Fatalln(err)
	}
	//"postgresql://username:password@localhost:5432/database_name"
	v := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v",
		envs["USER"], envs["PASSWORD"], envs["HOST"], envs["PORT"], envs["DATABASE"])
	fmt.Println(v)
	// v := fmt.Sprintf("user=%vpassword=%vhost=%vport=%vsslmode=disablepool_max_conns=10database=%v",
	// 	envs["USER"], envs["PASSWORD"], envs["HOST"], envs["PORT"], envs["DATABASE"])
	// fmt.Println(v)
	conn, err := pgx.Connect(context.Background(), returnUrl(v))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(context.Background())

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
