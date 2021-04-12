package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"encoding/json"
	//"net/url"
	"strings"
)

/*type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}*/

func updateHTML(url string) {

	resp, err := http.Get(url)

	if err != nil {
		//log.Fatalln(err)
		fmt.Printf("Error")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("index.html")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(string(body))

	if err2 != nil {
		log.Fatalln(err2)
	}

}

/*func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

func SendJqueryJs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/functions.js")
}*/

func mainPage(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("main.html"))
	tpl.Execute(w, nil)

	r.ParseForm()
	if r.Form["url"] != nil {
		if !strings.Contains("https://", r.Form["url"][0]) {
			url := "https://" + r.Form["url"][0]
			updateHTML(url)
			/*Block{
				Try: func() {
					updateHTML(url)
					Throw("Error")
				},
				Catch: func(exception Exception) {
					fmt.Printf("Error Caught")
				},
				Finally: func() {
					tpl.Execute(w, nil)
				},
			}.Do()*/
			tpl = template.Must(template.ParseFiles("index.html"))
			tpl.Execute(w, nil)
			fmt.Println("Set to Index")
		}
	}
}

/*func indexHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("index.html"))
	tpl.Execute(w, nil)
}*/

func main() {
	//updateHTML()
	fmt.Println("Restart")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	//fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	//mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", mainPage)
	//mux.HandleFunc("/", SendJqueryJs)
	http.ListenAndServe(":"+port, mux)
}
