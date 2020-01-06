package main

import (
	"GoServerTemplate/internal/config"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"time"
)

var templates = template.Must(template.ParseFiles("./templates/home.html", "./templates/gallery.html"))
var validPath = regexp.MustCompile("^/(home|gallery)/([a-zA-Z0-9]+)$")

//Page is a struct
type Page struct {
	Title       string
	RequestInfo ReqInfo
	Image       string
	PhotoList   []string
}

//ReqInfo is a struct
type ReqInfo struct {
	CreatedStr  string
	ElapasedStr string
}

func generateHome(title string) (*Page, error) {

	var reqinfo ReqInfo
	start := time.Now()
	reqinfo.ElapasedStr = time.Since(start).String()
	reqinfo.CreatedStr = time.Now().Format("Mon Jan _2 15:04:05 2006")
	return &Page{RequestInfo: reqinfo, Title: title}, nil
}

func generateGallery(title string) (*Page, error) {

	var reqinfo ReqInfo
	start := time.Now()
	reqinfo.ElapasedStr = time.Since(start).String()
	reqinfo.CreatedStr = time.Now().Format("Mon Jan _2 15:04:05 2006")

	file, err := os.Open("./addons/images")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders

	return &Page{RequestInfo: reqinfo, Title: title, PhotoList: list}, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := generateHome(title)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "home", p)
}

func galleryHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := generateGallery(title)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "gallery", p)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		fmt.Println(ip)
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	serverConfig := config.Init()
	http.HandleFunc("/home/", makeHandler(homeHandler))
	http.HandleFunc("/gallery/", makeHandler(galleryHandler))
	http.Handle("/addons/", http.StripPrefix("/addons/", http.FileServer(http.Dir("addons"))))

	fmt.Println("Reporting is Running....")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", serverConfig.Server.Port), nil))

}
