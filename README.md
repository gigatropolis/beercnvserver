# beercnvserver
## Web Services for converting BeerXML/BeerJSON formats between each other.

The project supplies binaries for different platforms that run a simple web server that can convert BeerXML and BeerJSON formats between each other. 

 Example from postXML1to2.go to send simple POST web request for XML1 to XML2
 ```go
 
 	xmlFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

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
```


