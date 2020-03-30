package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func (book Books) printBook() {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func main() {
	var p *int
	if p == nil {
		fmt.Printf("p pointer is nil: %x\n", p)
	}
	fmt.Printf("p pointer is not nil: %x\n", p)

	book := Books{title: "hoge", author: "huga", subject: "sub", book_id: 1}
	book.printBook()
	printBook(&book)
}
