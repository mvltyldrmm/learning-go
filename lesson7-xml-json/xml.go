package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main(){
	file, err := os.Open("sites.xml")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	data ,err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	v:= ObjectSites{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(v)
	
}

type ObjectSites struct {
	XMLName xml.Name `xml:"sites"`
	Version string `xml:"version,attr"`
}

type site struct {
	XMLname xml.Name `xml:"sites"`
	Name string `xml:"name"`
	Url string `xml:"url"`
}