package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var natoMap = map[string]string{
	"a": "Alpha", "b": "Bravo", "c": "Charlie", "d": "Delta",
	"e": "Echo", "f": "Foxtrot", "g": "Golf", "h": "Hotel",
	"i": "India", "j": "Juliett", "k": "Kilo", "l": "Lima",
	"m": "Mike", "n": "November", "o": "Oscar", "p": "Papa",
	"q": "Quebec", "r": "Romeo", "s": "Sierra", "t": "Tango",
	"u": "Uniform", "v": "Victor", "w": "Whiskey", "x": "X-ray",
	"y": "Yankee", "z": "Zulu",
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current working directory:", cwd)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Serve files from the public directory
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", fileServerHandler)
	http.HandleFunc("/nato/", natoHandler)

	fmt.Println("Server is running on port " + port)
	http.ListenAndServe(":"+port, nil)
}

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	// Serve index.html for the root path
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "public/index.html")
		return
	}

	// Construct the path to the file in the public directory
	filePath := "public" + r.URL.Path

	// Check if the file exists and is not a directory
	if info, err := os.Stat(filePath); err == nil && !info.IsDir() {
		http.ServeFile(w, r, filePath)
	} else {
		// File not found or is a directory, return 404
		http.NotFound(w, r)
	}
}

func natoHandler(w http.ResponseWriter, r *http.Request) {
	word := strings.ToLower(r.URL.Query().Get("word"))
	var natoWords []string
	for _, char := range word {
		if char < 'a' || char > 'z' {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		natoWord, exists := natoMap[string(char)]
		if !exists {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		natoWords = append(natoWords, natoWord)
	}
	response := strings.Join(natoWords, "<br />")
	fmt.Fprint(w, response)
}
