package author

import "fmt"

type Author struct{
	Name string
	Contact string
}

func NewAuthor(name,contact string) *Author{
	return &Author{Name: name, Contact: contact}
}

func (a *Author) WriteChapter(chapterTitle string,content string){
	fmt.Printf("Author %s is writing chapter titled '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

func (a *Author) ReviewChapter(chapterTitle string,content string){
	fmt.Printf("Author %s is reviewing chapter titled '%s'\n", a.Name, chapterTitle)
	fmt.Println(content)
}

func (a *Author) FinalizeChapter(chapterTitle string){
	fmt.Printf("Author %s is finalizing chapter titled '%s'\n", a.Name, chapterTitle)
}