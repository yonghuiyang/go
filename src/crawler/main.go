package main

import (
	"net/http"
	"log"
	"gopkg.in/xmlpath.v2"
	"fmt"
)

func main() {
	res, err := http.Get("https://wwww.baidu.com/")
	if err != nil {
		log.Fatal(err)
		fmt.Print(err)
	}

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		fmt.Print(err)
	}
	fmt.Println("here")
	root, err := xmlpath.ParseHTML(res.Body)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println("here")
	path, err := xmlpath.Compile("//title")

	if value, ok := path.String(root); ok {
		fmt.Println("Found:", value)
	}
}
