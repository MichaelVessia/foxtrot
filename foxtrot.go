package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", fileServerHandler) // Serve files
	http.HandleFunc("/nato/", natoHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Go Web Server!")
}

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

var natoMap = map[string]string{
    "a": "Alpha", "b": "Bravo", "c": "Charlie", "d": "Delta",
    "e": "Echo", "f": "Foxtrot", "g": "Golf", "h": "Hotel",
    "i": "India", "j": "Juliett", "k": "Kilo", "l": "Lima",
    "m": "Mike", "n": "November", "o": "Oscar", "p": "Papa",
    "q": "Quebec", "r": "Romeo", "s": "Sierra", "t": "Tango",
    "u": "Uniform", "v": "Victor", "w": "Whiskey", "x": "X-ray",
    "y": "Yankee", "z": "Zulu",
}

func natoHandler(w http.ResponseWriter, r *http.Request) {
	letter := strings.ToLower(r.URL.Query().Get("letter"))
	if len(letter) != 1 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	natoWord, exists := natoMap[letter]
	if !exists {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "<html><body><p>The NATO word for '%s' is: %s</p></body></html>", strings.ToUpper(letter), natoWord)
}
