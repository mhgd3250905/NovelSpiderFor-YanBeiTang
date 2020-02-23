package spider

type Book struct {
	BookId      string
	BookName    string
	BookImg     string
	Desc        string
	AuthorName  string
	AuthorImg   string
	Directories []Directory
}

type Directory struct {
	BookId      string
	DirectoryId string
	Name        string
	Url         string
	Content     string
}

type Detail_Book struct {
	BookName   string `json:"name"`
	BookImg    string `json:"bigCoverImage"`
	AuthorName string `json:"authorNick"`
	AuthorImg  string `json:"authorPic"`
}

type Detail_BookChapter struct {
	ChapterId int    `json:"id"`
	Name      string `json:"title"`
	Content   string `json:"content"`
	Desc      string `json:"shortContent"`
}

type Detail struct {
	Book          Detail_Book        `json:"book"`
	DetailChapter Detail_BookChapter `json:"bookChapter"`
	NextChapterId int                `json:"nextId"`
}
