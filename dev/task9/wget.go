package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	_ "net/url"
	"os"
	"strings"
)

func downloadOnePage(url, fileName string) {
	if fileName == "" {
		parts := strings.Split(url, "/")
		fileName = parts[len(parts)-1]
	}
	fmt.Println("filename = ", fileName)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("err")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Not valid status code: %d", response.StatusCode)
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("cannot create file with name: %s\n", fileName)
	}
	defer file.Close()
	fmt.Println("fileCopy")
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("cannot copy: %v", err)
	}
}

func getAllLinks(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//обрабатываем чтобы можно было делать поиск по тегам
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//срез для результатов поиска ссылок
	var links []string

	//Ищем и записываем все ссылки на странице
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		links = append(links, link)
	})
	return links
}

func downloadAllPages(address string) {
	newLinks := getAllLinks(address)
	fmt.Println("length: ", len(newLinks))
	for _, link := range newLinks {
		pathArr := strings.Split(link, "/")
		outputPath := "test/" + pathArr[len(pathArr)-1]
		if len(outputPath) > 2 {
			continue
		}
		resp, err := http.Get(link)
		if err != nil {
			fmt.Printf("cannot download page for this link: %v\n", err)
		}
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}()
		f, err := os.Create(outputPath)
		if err != nil {
			fmt.Println("Ошибка создания файла")
		}
		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func main() {
	output := flag.String("O", "", "write documents to FILE")
	recursive := flag.Bool("r", false, "specify recursive download")
	flag.Parse()
	if !*recursive {
		downloadOnePage(flag.Arg(0), *output)
		return
	}
	downloadAllPages(flag.Arg(0))
}
