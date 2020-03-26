package controller

import(
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
  "log"
)

var tpl = template.Must(template.ParseFiles("../web/menu.html"))
var tp2 = template.Must(template.ParseFiles("../web/somos.html"))
var tp3 = template.Must(template.ParseFiles("../web/contacto.html"))
var tp4 = template.Must(template.ParseFiles("../web/categorias.html"))
var tp5 = template.Must(template.ParseFiles("../web/admins.html"))
var tp6 = template.Must(template.ParseFiles("../web/productos.html"))


func Menu(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func Somos(w http.ResponseWriter, r *http.Request) {
	tp2.Execute(w, nil)
}

func Contacto(w http.ResponseWriter, r *http.Request) {
	tp3.Execute(w, nil)
}

func Catego(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  categoria:= vars["categoria"]
  log.Println(categoria)
	tp4.Execute(w,categoria)
}

func Supersecret(w http.ResponseWriter, r *http.Request) {
	tp5.Execute(w, nil)
}

func Productos(w http.ResponseWriter, r *http.Request) {
	tp6.Execute(w, nil)
}
