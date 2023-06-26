package main

import (
	"strconv"
	"io"
	"os"
	"net/http"
	"text/template"
)
var check1 bool
var check2 bool
// Starting to find the html page
func StartPage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" || req.Method != "GET" {
		error404(res)
		return
	} else {
		pageTemplate, err := template.ParseFiles("./templates/home.page.tmpl")
		if err != nil {
			error500(res)
			return
		}
		res.WriteHeader(200)
		pageTemplate.Execute(res, nil)
	}
}

// Form data after submiting
func SubmitTing(res http.ResponseWriter, req *http.Request) {
	if req.FormValue("sub")=="sub" {
		check1=true
	}
	
	if req.Method != "POST" {
		error405(res)
		return
	}

	if req.URL.Path != "/ascii-art" {
		error404(res)
		return
	}

	tpl, er := template.ParseFiles("./templates/home.page.tmpl")
	if er != nil {
		error500(res)
		return
	}

	req.ParseForm()

	if req.FormValue("banner") != "standard" &&
		req.FormValue("banner") != "thinkertoy" &&
		req.FormValue("banner") != "shadow" {
		error500(res)
		return
	}

	txt := req.FormValue("input")
	ban := req.FormValue("banner")

	if !isPrintable(txt) {
		error400(res)
		return
	}

	out := ConvertStr(txt, ban)

	// Afficher le résultat dans le template
	res.WriteHeader(200)
	tpl.Execute(res, out)
	data := os.MkdirAll("data", 0755)
	if data != nil {
		error500(res)
		return
	}
	// Enregistrer le texte ASCII dans un fichier
	data = os.WriteFile("./data/result.txt", []byte(out), 0644)
	if data != nil {
		error500(res)
		return
	}
	data = os.WriteFile("./data/result.exe", []byte(out), 0644)
	if data != nil {
		error500(res)
		return
	}
	data = os.WriteFile("./data/result.pdf", []byte(out), 0644)
	if data != nil {
		error500(res)
		return
	}
}

func download(res http.ResponseWriter, req *http.Request) {
	if req.FormValue("telecharge")=="telecharge" {
		check2=true
	}
	if check1 && check2 {
	formatType := req.FormValue("fileformat")

	filename := "./data/result." + formatType

	// Ouvrir le fichier à télécharger
	f, err := os.Open(filename)
	if err != nil {
		error500(res)
		return
	}
	defer f.Close()

	// Récupérer les informations sur le fichier
	fileInfo, err := f.Stat()
	if err != nil {
		error500(res)
		return
	}
	fileSize := fileInfo.Size()

	// Définition des en-têtes
	res.Header().Set("Content-Disposition", "attachment; filename=result."+formatType)
	res.Header().Set("Content-Type", "text/plain")
	res.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))
	f.Seek(0, io.SeekStart)
	// Transmission du contenu du fichier
	io.Copy(res, f)
	}else{
		error400(res)
	}
}

