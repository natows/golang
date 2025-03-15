package main

import (
	"fmt"
	"net/http"
	"io"
	
)

func fetchData() io.ReadCloser{
	url := "https://public.opendatasoft.com/api/explore/v2.1/catalog/datasets/coronavirus-covid-19-pandemic-usa-counties/records?limit=100"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp.Body
}

func main() {
	data := fetchData()
	defer data.Close()
	content, err := io.ReadAll(data)
	fmt.Println(content)
	fmt.Println(err)
}
