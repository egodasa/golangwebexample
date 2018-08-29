package controllers

import (
	"net/http"
	"strings"
	"web/helper/views"
	"web/models"
)

func DaftarMahasiswa(w http.ResponseWriter, r *http.Request) {
	konten := strings.Join(models.GetMahasiswa(), " - ")
	views.Load("daftar-mahasiswa.html", map[string]string{
		"Title":   "Judul Halaman",
		"Content": konten,
	}, w)
}
