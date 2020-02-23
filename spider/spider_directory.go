package spider

import (
	"fmt"
	"github.com/gocolly/colly"
)

const MAIN = "http://www.ebtang.com"

func DirectorySpider(book Book) {

	//directories:=make([]Directory,0)

	startUrl := fmt.Sprintf("http://www.ebtang.com/book/%v/directory", book.BookId)

	//解析页面新闻条目收集器
	pageCollector := colly.NewCollector()

	//解析页面新闻条目
	newsItemSelectorStr := "#directoryList > li > b:nth-child(1)"

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
		name, exist := e.DOM.Attr("d-title")
		if !exist {
			fmt.Printf("d-title属性不存在！")
		}
		directoryId, exist := e.DOM.Attr("d-id")
		if !exist {
			fmt.Printf("d-id属性不存在！")
		}

		directory := Directory{BookId: book.BookId, DirectoryId: directoryId, Name: name, Url: e.Request.AbsoluteURL(directoryId)}
		DetialSpider(book, directory)
	})

	pageCollector.OnScraped(func(response *colly.Response) {
	})

	pageCollector.OnError(func(response *colly.Response, e error) {
		fmt.Println("pageCollector.OnError: ", e)
	})

	pageCollector.UserAgent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	pageCollector.Visit(startUrl)

}
