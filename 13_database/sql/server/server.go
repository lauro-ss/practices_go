package server

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitServer(db *pgxpool.Pool) {
	m := http.NewServeMux()

	m.Handle("/animals", GetAnimal(db))

	http.ListenAndServe(":4500", m)
	fmt.Println("Server running")
}

func GetAnimal(db *pgxpool.Pool) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		fmt.Println(id)
		var (
			name  string
			emoji string
		)
		//err := db.QueryRow(r.Context(), "select name, emoji from animal where id = $1;", id).Scan(&name, &emoji)
		//query := "select name, emoji from animal where name = " + id + " ;" //SQL Injection
		//rows, err := db.Query(r.Context(), query)
		rows, err := db.Query(r.Context(), "select name, emoji from animal where name = $1;", id) //No SQL Injection

		if err != nil {
			fmt.Println(err)
			http.NotFound(w, r)
			return
		}

		defer rows.Close()
		if rows.CommandTag().RowsAffected() == 0 {
			http.NotFound(w, r)
			return
		}
		for rows.Next() {
			rows.Scan(&name, &emoji)
			w.Write([]byte(name + " " + emoji))
		}
		if rows.Err() != nil {
			fmt.Println(rows.Err())
			http.NotFound(w, r)
			return
		}

	}
	return http.HandlerFunc(fn)
}
