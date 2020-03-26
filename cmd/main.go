package main

import (
	"net/http"
	"os"
  "Binfa/API/controller"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

  fs := http.FileServer(http.Dir("css"))
  mux.Handle("/css/", http.StripPrefix("/css/", fs))

	mux.HandleFunc("/", controller.Menu)
  mux.HandleFunc("/QuienesSomos", controller.Somos)
  mux.HandleFunc("/Contacto", controller.Contacto)
	mux.HandleFunc("/categoria/{categoria}/", controller.Catego)
	mux.HandleFunc("/supersecret", controller.Supersecret)
	mux.HandleFunc("/productos", controller.Productos)

	http.ListenAndServe(":"+port, mux)
}
