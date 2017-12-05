package helper

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func LayoutFiles() []string {
	files, err := filepath.Glob("templates/layouts/*.html")
	if err != nil {
		log.Panic(err)
	}
	return files
}

func LayoutFilesIncludes() []string {
	filesInc, err := filepath.Glob("templates/layouts/includes/*.html")
	if err != nil {
		log.Panic(err)
	}
	return filesInc
}

type View struct {
	Index  Page
	Show   Page
	New    Page
	Create Page
	Edit   Page
	Update Page
	Delete Page
}

type Page struct {
	Template *template.Template
	Layout   string
}

func GetAuth(r *http.Request) map[string]string {
	authSS := GetSession("AuthSession", r)
	var authJson = make(map[string]string)
	err := json.Unmarshal([]byte(authSS), &authJson)
	if err != nil {
		fmt.Println("Helper View: GetAuther", err)
	}
	return authJson
}

func (self *Page) Render(w http.ResponseWriter, r *http.Request, data interface{}) error {

	authSession := GetAuth(r)
	// fmt.Println(len(authSession))
	sessionData := map[string]string{
		"AuthName": authSession["name"],
	}
	result := map[string]interface{}{
		"Data":        data,
		"PrivateData": sessionData,
	}
	// fmt.Println(result)
	return self.Template.ExecuteTemplate(w, self.Layout, result)
}
