package main

import "fmt"

// Declarar structure Books
type Books struct {
	title, autor, subject string
	bookID                int
}

func main() {
	book1 := Books{title: "Go", autor: "Santiago", subject: "Learning", bookID: 1020}
	book2 := Books{title: "Python", autor: "Santiago", subject: "Learning", bookID: 1021}

	fmt.Printf("Book 1 info: %s, %s, %s, %d\n", book1.autor, book1.subject, book1.title, book1.bookID)
	fmt.Printf("Book 2 info: %s, %s, %s, %d\n", book2.autor, book2.subject, book2.title, book2.bookID)
}


