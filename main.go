package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
)

type File struct {
    Name string
    Size int64
}

type Drive struct {
    Name string
    Size int64
    files []File
}
type User struct {
    Username string
    Password string
    drives []Drive
}

var templates *template.Template

func main() {
    // Zmienne środowiskowe

    // WAŻNE DO ROBIENIA FILE EXPLORERA!

    fmt.Println("Server is running...")
    // TO-DO
    // uniwersalny handler do plików html który przetwarza nazwy z funkcji HandleFunc i wpisuje je do szablonu
    // znajdując plik html o tej nazwie 
    templates,_ = template.ParseGlob("src/templates/*.html")
    //handlers

    //HandleFunc waits for a function, where Handle waits for a Handler 
    http.HandleFunc("/content", contentHandler) 
    http.HandleFunc("/login", loginHandler) 
    http.HandleFunc("/style", styleHandler) 

    // Obsługa plików statycznych

    handler := http.FileServer(http.Dir("src"))
    log.Fatal(http.ListenAndServe(":8000", handler))  // entering nil implicitly uses DefaultServeMux
}

func styleHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("src/output.html"))
    tmpl.Execute(w, nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "r.URL.Path = %q\n", r.URL.Path)
    tmpl := template.Must(template.ParseFiles("src/index.html"))
    tmpl.Execute(w, nil)
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
    files := map[string][]File{ // Ściąganie plików z BD
        "Files": {
            {Name: "zdjecie.png", Size: 1024},
            {Name: "tutorial.txt", Size: 2048},
        },
    }
    htmlStr := fmt.Sprintf( "{{ range .Files }} <p class='text-white'> {{ .Name }} - {{ .Size }} </p> {{ end }}" )
    tmpl, _ := template.New("content").Parse( htmlStr )
    tmpl.Execute(w, files)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w,"login.html", nil)
}
// Using a database
//import (
//	"github.com/deliskyxd/myfilemanager/database"
//	"github.com/gofiber/fiber/v2"
//)
//
//func main(){
//    database.Connect()
//    PORT := "8500"
//    app := fiber.New()
//    app.Listen(":" + PORT)
//}

