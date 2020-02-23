package spider

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func GenerateTxtBook(book Book) {
	filename := fmt.Sprintf("%v.txt", book.BookName)
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		// 创建文件失败处理
		fmt.Println("创建文件失败！")
	} else {
		content := ""
		for i, _ := range book.Directories {
			content += fmt.Sprintf("%v.%v%v",i,book.Directories[i].Name, "\n")
			content += fmt.Sprintf("%v%v", book.Directories[i].Content, "\n")
		}
		content = strings.Replace(content, "\n", "\n\n", -1)
		_, err = f.Write([]byte(content))
		if err != nil {
			// 写入失败处理
			fmt.Println("写入文件失败！")
		}
		fmt.Println("写入文件成功，3秒后本窗口关闭...")
		time.Sleep(3000)
		os.Exit(0)
	}
}
