package view

import (
	"github.com/golang/glog"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	//"../../session"
	"log"
	"os"
)

var (
	indexTemplate *template.Template
	templates     *template.Template
)

func Init(templateDir string) {
	if len(templateDir)<=0{
		log.Println("Init templateDir empty")
		return
	}
	//目录不存在
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		log.Println("Init templateDir not exist")
		return
	}

	var allfile []string
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		glog.Errorln(err)
		return
	}
	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".html") {
			allfile = append(allfile,templateDir+fileName)
		}
	}
	templates = template.Must(template.ParseFiles(allfile...))
	indexTemplate = templates.Lookup("index.html")
}

func LoadTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if r.Method == "GET" {
		indexTemplate.Execute(w, nil)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
