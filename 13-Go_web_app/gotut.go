package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")


type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	keywords []string `xml:"url>news>keywords"`
	Location []string `xml:"url>loc"`
}

func main() {
	var s SitemapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	
	xml.Unmarshal(bytes, &s)  //unmarshal into memory address of s

	//fmt.Println(s.Locations)
	fmt.Printf("Here %s some %s", "are", "variables")
	
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)  //get response 
		bytes, _ := ioutil.ReadAll(resp.Body)  
		xml.Unmarshal(bytes, &n)

		fmt.Println(n)

	}
}