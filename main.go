package main

import (
	"NovelSpider/spider"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("请输入-雁北堂-图书ID:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	book:=spider.Book{BookId:input.Text()}
	spider.DirectorySpider(book)
}
