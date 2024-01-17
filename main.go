package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
    //"github.com/gofiber/fiber/v2"
    "github.com/deliskyxd/myfilemanager/database"
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


func main() {
    database.Connect()
    fmt.Println("Server is running...")
    // Zmienne środowiskowe

    // WAŻNE DO ROBIENIA FILE EXPLORERA!
    //handler := http.FileServer(http.Dir("src"))
    index := template.Must(template.ParseFiles("src/index.html"))
    //HandleFunc waits for a function, where Handle waits for a Handler 
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        index.Execute(w, nil)
    })

    log.Fatal(http.ListenAndServe(":8080", mux))  // entering nil implicitly uses DefaultServeMux
}

//func contentHandler(w http.ResponseWriter, r *http.Request) {
//    files := map[string][]File{ // Ściąganie plików z BD
//        "Files": {
//            {Name: "zdjecie.png", Size: 1024},
//            {Name: "tutorial.txt", Size: 2048},
//        },
//    }
//    htmlStr := fmt.Sprintf( "{{ range .Files }} <p class='text-white'> {{ .Name }} - {{ .Size }} </p> {{ end }}" )
//    tmpl, _ := template.New("content").Parse( htmlStr )
//    tmpl.Execute(w, files)
//}

// Using a database
//import (
//	"github.com/gofiber/fiber/v2"
//	"github.com/deliskyxd/myfilemanager/database"
//)
//
//func main(){
//    database.Connect()
//    PORT := "8500"
//    app := fiber.New()
//    app.Listen(":" + PORT)
//}

