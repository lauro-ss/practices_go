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
		err := db.QueryRow(r.Context(), "select name, emoji from animal where id = $1;", id).Scan(&name, &emoji)
		if err != nil {
			fmt.Println(err)
			http.NotFound(w, r)
			return
		}
		w.Write([]byte(name + " " + emoji))
	}
	return http.HandlerFunc(fn)
}
