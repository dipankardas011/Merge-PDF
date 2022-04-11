package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! 🐳☸️🚀👍🏼🥳✅ %s", time.Now())
}

func MergePdf() {
	cmd := exec.Command("qpdf", "--empty", "--pages", "./uploads/00.pdf", "./uploads/01.pdf", "--", "./uploads/resrelt.pdf")
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error in MergePDF", err)
	}
}

func main() {

	port := os.Getenv("PORT")
	fmt.Println(port)
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

	http.ListenAndServe(":"+port, nil)
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
	cmd := exec.Command("rm", "-Rf", "./uploads/")
	err := cmd.Run()
	if err != nil {
		log.Println("Error in deleting Existing PDFs", err)
	}
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Cleared the data!!✅\t%s", time.Now())
}
