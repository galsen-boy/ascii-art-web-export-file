package main

import (
	"strconv"
	"io"
	"os"
	"net/http"
	"text/template"
)
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
	if req.Method != "POST" {
		error405(res)
		return
	}

	if req.URL.Path != "/ascii-art" {
		error404(res)
		return
	}
	tpl, er := template.ParseFiles("./templates/home.page.tmpl")
	if er !=nil {
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

	_, err := os.ReadFile(ban + ".txt")
	if err != nil {
		error500(res)
		return
	}
	if !isPrintable(txt) {
		error400(res)
		return
	}

	if err != nil {
		error500(res)
	}
	res.WriteHeader(200)
	tpl.Execute(res, ConvertStr(txt, ban))

	// processus pour telecharger le texte ascii-art dans un fichier
	
	out := ConvertStr(txt, ban)
	fich1 := os.WriteFile("asciiResult.txt", []byte(out), 0644)
	if fich1 != nil {
	error500(res)
	}
	fich2 := os.WriteFile("asciiResult.doc", []byte(out), 0644)
	if fich2 != nil {
		error500(res)
	}
	fich3 := os.WriteFile("asciiResult.pdf", []byte(out), 0644)
	if fich3 != nil {
		error500(res)
	}
}

	
	func download(res http.ResponseWriter, req *http.Request) {

		formatType := req.FormValue("fileformat")
	
		f, err := os.Open("asciiResult." + formatType)
		if err !=nil{
			error500(res)
			return
		}

		defer f.Close()
		file, _ := f.Stat()
		fsize := file.Size()
		sfSize := strconv.Itoa(int(fsize))
		res.Header().Set("Content-Disposition", "attachment; filename=asciiresult."+formatType)
		res.Header().Set("Content-Type", "text/html")
		res.Header().Set("Content-Length", sfSize)
		io.Copy(res, f)
		
	}




	