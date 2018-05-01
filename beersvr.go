package main

import (
	//"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	//"net/url"
	"os"
	//"reflect"
	//"strconv"
	//"time"
	//"../xml/beerxml"
	"github.com/gigatropolis/beercnv"
)

// --------------------------------------------------------------------

type ConvData struct {
	FileName   string
	StrConvert string
	XML2       string
}

var mux map[string]func(http.ResponseWriter, *http.Request)

type BeerServerHandler struct {
}

func (h *BeerServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("URL = ", r.URL.String())
	if h, ok := mux[r.URL.EscapedPath()]; ok {
		h(w, r)
		return
	}

	fmt.Fprintf(w, "URL Not Found: %s", r.URL.String())
}

func convertxml1to2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		body, err := ioutil.ReadAll(r.Body)
		//defer resp.Body.Close()

		if err != nil {
			fmt.Fprintf(w, "Unable to Read post body")
			return
		}

		fmt.Printf("body: %s\n", body)

		beer2 := beercnv.BeerXml2{}
		err = beercnv.ConvertXML1to2(r.Body, &beer2)

		if err != nil {
			fmt.Fprintf(w, "error: %v\n", err)
			return
		}

		output, err := xml.MarshalIndent(beer2, "  ", "   ")

		if err != nil {
			fmt.Fprintf(w, "error: %v\n", err)
			return
		}

		fmt.Fprintf(w, string(output))

	}
}

func convertxml1to2file(w http.ResponseWriter, r *http.Request) {

	data := ConvData{StrConvert: "Upload BeerXML 1.0 file to be converted to BeerXML 2.x format"}

	if r.Method == "GET" {

		t, _ := template.New("UploadBeerXML1.tpl").ParseFiles("templates/UploadBeerXML1.tpl")
		template.Must(t.Clone())
		t.Execute(w, data)

	} else {
		file, header, err := r.FormFile("beerxml1file")

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		defer file.Close()

		filename := "downloads/xml1/" + header.Filename
		out, err := os.Create(filename)
		if err != nil {
			fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		defer out.Close()

		// write the content from POST to the file

		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		file.Seek(0, 0)

		data.StrConvert = "File uploaded successfully: " + header.Filename

		beer2 := beercnv.BeerXml2{}

		err = beercnv.ConvertXML1to2(file, &beer2)

		if err != nil {
			fmt.Fprintf(w, "error: %v\n", err)
			return
		}

		output, err := xml.MarshalIndent(beer2, "  ", "   ")

		if err != nil {
			fmt.Fprintf(w, "error: %v\n", err)
			return
		}

		xml2, err := os.Create("downloads/xml2/" + header.Filename)
		if err != nil {
			fmt.Fprintf(w, "Unable to create XML2 file for writing. Check your write access privilege")
			return
		}

		fmt.Fprintf(xml2, "%s", output)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		data.FileName = header.Filename
		data.XML2 = string(output)

		tXml2, _ := template.New("UploadedBeerXML1.tpl").ParseFiles("templates/UploadedBeerXML1.tpl")
		template.Must(tXml2.Clone())
		tXml2.Execute(w, data)

	}
}

func main() {

	server := http.Server{
		Addr:    ":8080",
		Handler: &BeerServerHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = convertxml1to2file
	mux["/convertxml1to2file"] = convertxml1to2file
	mux["/convertxml1to2"] = convertxml1to2

	server.ListenAndServe()

}
