package views

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var cwd, _ = os.Getwd()

/*
	Load(name string, konten map[string]string, w http.ResponseWriter)
	Parameter:
	- name, name of file views
	- konten, data to be sent into views
	- w, http.ResponseWriter from controllers parameter
*/
func Load(name string, konten map[string]string, w http.ResponseWriter) {
	t, err := template.ParseFiles(filepath.Join(cwd, "web/../views/"+name))
	if err != nil {
		fmt.Println(err.Error())
		fmt.Fprintf(w, "Error %s ", err)
	} else {
		t.Execute(w, konten)
	}
}
