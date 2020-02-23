package spider

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

const jsonUrl = "http://www.ebtang.com/book/readbook/%v/%v/more?randomString=%v"

var RANDOMSTRING string

func DetialSpider(book Book, directory Directory) {
	startUrl := directory.Url

	pageCollector := colly.NewCollector()
	jsonCollector := colly.NewCollector()

	newsItemSelectorStr := "#randomString"

	pageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("Error: ", e)
	})

	pageCollector.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
	})

	pageCollector.OnHTML(newsItemSelectorStr, func(e *colly.HTMLElement) {
		randomString, exist := e.DOM.Attr("value")
		if !exist {
			fmt.Printf("value属性不存在！")
		}

		//fmt.Println(randomString)
		RANDOMSTRING = randomString
		totalJsonUrl := fmt.Sprintf(jsonUrl, directory.BookId, directory.DirectoryId, randomString)
		jsonCollector.Visit(totalJsonUrl)
	})

	pageCollector.OnScraped(func(response *colly.Response) {
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("pageCollector.OnError: ", e)
	})

	jsonCollector.OnResponse(func(response *colly.Response) {
		jsonBytes := response.Body
		//fmt.Println(string(jsonBytes))
		detail := Detail{}
		err := json.Unmarshal(jsonBytes, &detail)
		if err != nil {
			fmt.Println("json Unmarshal failed!")
		}

		directory.BookId = book.BookId
		directory.DirectoryId=strconv.Itoa(detail.DetailChapter.ChapterId)
		directory.Url = response.Request.AbsoluteURL(strconv.Itoa(detail.DetailChapter.ChapterId))
		directory.Content = detail.DetailChapter.Content
		directory.Name = detail.DetailChapter.Name

		book.AuthorImg = detail.Book.AuthorImg
		book.AuthorName = detail.Book.AuthorName
		book.BookImg = detail.Book.BookImg
		book.BookName = detail.Book.BookName
		book.Desc = detail.DetailChapter.Desc
		book.Directories = append(book.Directories, directory)

		fmt.Printf("章节：%v - %v下载完毕\n", directory.DirectoryId, directory.Name)
		if detail.NextChapterId == -1 {
			//fmt.Println(book)
			GenerateTxtBook(book)
		} else {
			totalJsonUrl := fmt.Sprintf(jsonUrl, directory.BookId, detail.NextChapterId, RANDOMSTRING)
			jsonCollector.Visit(totalJsonUrl)
		}
	})

	pageCollector.UserAgent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	pageCollector.Visit(startUrl)

}
