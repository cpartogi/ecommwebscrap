package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Product struct {
	Name      string
	ImageLink string
	Price     string
	Rating    int
	StoreName string
}

func main() {
	startTime := time.Now()

	// define file name
	file, err := os.Create("export_mobilephone.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//set header
	headers := []string{"Name", "ImageLink", "Price", "Rating", "StoreName"}
	writer.Write(headers)

	// set user agent and domain
	c := colly.NewCollector(
		colly.AllowedDomains("www.tokopedia.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	fmt.Println("start get data ...")

	// start get content
	c.OnHTML(".css-16vw0vn", func(e *colly.HTMLElement) {
		product := Product{}
		product.Name = e.ChildText(".css-1bjwylw")
		product.ImageLink = e.ChildAttr(".css-1c0vu8l img", "src")
		product.ImageLink = strings.Replace(product.ImageLink, ",", ";", -1)
		product.Price = e.ChildText(".css-o5uqvq")
		product.StoreName = e.ChildText(".css-vbihp9 > span:nth-child(2)")
		spanClass := e.ChildAttrs("img", "class")

		// convert product rating image to number
		product.Rating = 0
		for i := 0; i < len(spanClass); i++ {
			if spanClass[i] != "fade" {
				product.Rating = i
			}
		}

		// convert to string for csv
		strRating := strconv.Itoa(product.Rating)

		// fill data
		row := []string{product.Name, product.ImageLink, product.Price, strRating, product.StoreName}
		writer.Write(row)
	})

	// set pagination for 100+ rows
	for p := 1; p < 19; p++ {
		strP := strconv.Itoa(p)

		urlp := "https://www.tokopedia.com/p/handphone-tablet/handphone?page=" + strP
		err = c.Visit(urlp)

	}

	if err != nil {
		fmt.Println(err)
	}

	endTime := time.Now()

	resTime := endTime.Sub(startTime).Seconds()

	fmt.Println("execution time : ", resTime, " s")
}
