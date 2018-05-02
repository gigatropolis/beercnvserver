package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"net/url"
	//"time"
	//"bytes"
	//"github.com/gigatropolis/beercnv"
)

func main() {

	//beer := BeerXml{}
	//beer2 := beercnv.BeerXml2{}

	//filename := "Recipies\\xml\\include-hops.xml"
	filename := "Recipies\\xml\\include-hops.xml"
	//filename := "Recipies\\xml\\nhc_2015.xml"

	xmlFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	/*
		p := xml.NewDecoder(xml1)
		p.CharsetReader = charset.NewReaderLabel
		err = p.Decode(&beer)

		if err != nil {
			panic(err)
		}

		buf := bytes.Buffer
		err = xml.Unmarshal(buf, &beer)

		if err != nil {
			panic(err)
		}
	*/
	resp, err := http.Post("http://localhost:8080/convertxml1to2", "", xmlFile)

	if err != nil {
		fmt.Printf("server:%s POST error:%s\n", "localhost", err.Error())
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("server:%s Read error:%s\n", "localhost", err.Error())
	}

	fmt.Printf("%s\n", body)

}
