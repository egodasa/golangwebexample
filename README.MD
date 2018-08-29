### Example of Simple Web Apps Using Golang
This example using MVC concept. Below is the structure of this web apps :
```
/root
----- main.go
----- controllers/
---------- files_of_controllers.go
----- views/
---------- files_of_views.html
----- models/
---------- files_of_models.go
----- routess/
---------- routes.go //All controllers will be registered here
----- helper/
---------- views.go // This included helper method for views 
```

### How to create a new page?
1. Make a file inside controllers folder, ex: `mahasiswa.go`. Below is the content of `mahasiswa.go`
```
package controllers

import (
	"net/http"
	"strings"
	"web/helper/views"
	"web/models"
)

func DaftarMahasiswa(w http.ResponseWriter, r *http.Request) {
	views.Load("daftar-mahasiswa.html", map[string]string{
		"Title":   "Judul Halaman",
		"Content": ""Kontent Halaman,
	}, w)
}
```
2. Create new files in `views/` folder. We name it `daftar-mahasiswa.html`. Below is the content of that file:
```
<html>
    <head>
        <title>{{.Title}}</title>
    </head>
    <body>
        {{.Content}}
    </body>
</html>
```
3. Open `routes/routes.go` and write the code like below :
```
package routes

import (
	"fmt"
	"net/http"
	"web/controllers" // we import controllers
)

func Routes() {
	http.HandleFunc("/mahasiswa", controllers.DaftarMahasiswa) // this is the controllers we made before
	//http.HandleFunc("/other_url", controllers.otherControllers) insert other controllers here 
}
```
**Note: All controllers must be registered inside `Routes()` method. So, You should**
4. Open terminal and cd into root directory. And run `go run main.go`.
5. Open `http://127.0.0.1:8080/mahasiswa`
6. That's it.

### Notes
**This is for personal documentation only. Don't use this structure. You can create Your own framework or use another framework such as `Beego` or other.**