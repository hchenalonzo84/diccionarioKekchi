package main

import (
	"diccionario/db"
	"diccionario/models"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// estructuras de datos
type Usuarios struct {
	UserName string
	Edad     int
}

var templates = template.Must(template.New("T").ParseGlob("template/*.html"))
var templates2 = template.Must(template.New("T").ParseGlob("template/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("template/error/error.html"))

// render template
func ManejoError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}
func RenderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)
	if err != nil {
		ManejoError(rw, http.StatusInternalServerError)

	}
}
func RenderTemplate2(rw http.ResponseWriter, name string, data interface{}) {
	err := templates2.ExecuteTemplate(rw, name, data)
	if err != nil {
		ManejoError(rw, http.StatusInternalServerError)
	}
}

// Handler
func Inicio(rw http.ResponseWriter, r *http.Request) {

	usuario := Usuarios{"Hugo", 39}

	// template.Execute(rw, usuario)
	RenderTemplate(rw, "index.html", usuario)
}
func Registro(rw http.ResponseWriter, r *http.Request) {
	RenderTemplate(rw, "registro.html", nil)

}
func Principal(rw http.ResponseWriter, r *http.Request) {
	RenderTemplate(rw, "main.html", nil)
}

func main() {
	//estableciendo conexion a base de datos
	db.Connect()
	fmt.Println(db.ExistsTable("users"))
	db.CreateTable(models.UserSchema, "users")
	// user := models.CreateUser("Rolando", "rolando123", "rolando@hotmail.com")
	// fmt.Println(user)
	// db.TruncateTable("users")
	// users := models.LIstarUsers()
	// fmt.Println(users)
	users := models.GetUsers(3)
	fmt.Println(users)
	users.Username = "Pedro"
	users.Save()
	fmt.Println(models.LIstarUsers())
	db.Close()
	//db.Ping()
	// Archivos estaticos
	staticFiles := http.FileServer(http.Dir("static"))
	//Mux es una ruta asociada a un handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", Inicio)
	mux.HandleFunc("/registro", Registro)
	mux.HandleFunc("/main", Principal)
	//Mux de static Files
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	// crear servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Print("escuchando en el puerto 3000")
	log.Fatal(server.ListenAndServe())

}
