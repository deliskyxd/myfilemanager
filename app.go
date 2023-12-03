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


func main() {
    // Zmienne środowiskowe
    handler := http.FileServer(http.Dir(".")) // Obsługa plików

    fmt.Println("Server is running...")
    index := func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("src/index.html"))
        files := map[string][]File{ // Ściąganie plików z BD
            "Files": {
                {Name: "zdjecie.png", Size: 1024},
                {Name: "tutorial.txt", Size: 2048},
            },
        }
        tmpl.Execute(w, files)
    }
    login := func(w http.ResponseWriter, r *http.Request) {

        htmlStr := fmt.Sprintf("<form action='index' method='post'> <label for='username' style='color: white;'>Użytkownik:</label><br> <input type='text' id='username' name='username' ><br> <label for='password' style='color: white;' >Hasło:</label><br> <input type='password' id='password' name='password'><br><br> <input type='submit' value='Zaloguj' style='background-color: white;'> </form>")
        tmpl, _ := template.New("login").Parse(htmlStr)
        tmpl.Execute(w, nil)
    }
    //HandleFunc czeka na funkcje, gdzie Handle czeka na handler
    http.HandleFunc("/login", login) // Obsługiwanie loginu
    http.HandleFunc("/index", index) // Obsługiwanie index.html
    http.Handle("/", handler) // Obsługiwanie zapytań

    log.Fatal(http.ListenAndServe(":8000", nil)) 
}
