package handler

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }

	// if r.Method != http.MethodGet {
	// 	w.Header().Add("Allow", http.MethodGet)
	// 	m := fmt.Sprintf("%v method not allowed", http.StatusMethodNotAllowed)
	// 	http.Error(w, m, http.StatusMethodNotAllowed)
	// 	return
	// }
	fmt.Println(r.URL.Path)
	w.Write([]byte("Home"))
}
