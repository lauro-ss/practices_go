package main

import (
	"fmt"
	"log"
)

func main() {
	envs, err := SourceEnv()
	if err != nil {
		log.Fatalln(err)
	}
	//"postgres://username:password@localhost:5432/database_name"
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		envs["USER"], envs["PASSWORD"], envs["HOST"], envs["PORT"], envs["DATABASE"])
	fmt.Println(url)
	// conn, err := pgx.Connect(context.Background(), url)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer conn.Close(context.Background())

}
