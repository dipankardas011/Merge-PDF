package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! 🐳☸️🚀👍🏼🥳✅ %s", time.Now())
}

func MergePdf() error {
	cmd := exec.Command("qpdf", "--empty", "--pages", "./uploads/00.pdf", "./uploads/01.pdf", "--", "./uploads/resrelt.pdf")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error in MergePDF", err)
	}
	return err
}

func getPort() string {
	port := os.Getenv("PORT")
	fmt.Printf("ENV{Port}: %v\n", port)
	if port == "" {
		return ":8080"
	}
	return ":" + port
}

func main() {
	uploadedStat = false
	err := os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/downloads", DownloadFile)
	http.HandleFunc("/", html)
	http.HandleFunc("/pdf/clear", clearExistingpdfs)
	http.HandleFunc("/css/styles", CSSFileAccess)
	http.HandleFunc("/html/about", AboutHTMLAccess)

	http.ListenAndServe(getPort(), nil)
}

func CSSFileAccess(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/style.css")
}

func AboutHTMLAccess(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/About.html")
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "uploads/resrelt.pdf")
	}
}

func html(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "web/index.html")
	}
}

func clearExistingpdfs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("./templates/upload.html")

	var x templateStat
	if err != nil {
		x = templateStat{
			Status: "Internal Server error 501 ⚠️",
		}
	}

	cmd := exec.Command("rm", "-Rf", "./uploads/")
	err = cmd.Run()

	if err != nil {
		x = templateStat{
			Status: "CRITICAL ERROR 503 ❌",
		}
	}
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		panic(err)
	} else {
		x = templateStat{
			Status: fmt.Sprintf("Cleared the data!!✅\t%s", time.Now()),
		}
	}
	t.Execute(w, x)
}
