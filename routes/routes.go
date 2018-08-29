package routes

import (
	"fmt"
	"net/http"
	"web/controllers"
)

func Routes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello. This is index page! ", r)
	})
	http.HandleFunc("/mahasiswa", controllers.DaftarMahasiswa)
}
