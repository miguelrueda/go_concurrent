package book

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf("Title:\t\t%q\n"+"Author:\t\t%q\n"+"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

func GetBooks() []Book {
	return books
}

var books = []Book{
	{
		ID:            1,
		Title:         "The Hitchhiker's Guide to the Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R. Tokien",
		YearPublished: 1937,
	},
	{
		ID:            3,
		Title:         "A Tale of Two Cities",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	{
		ID:            4,
		Title:         "A Tale of Two Cities 2",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	{
		ID:            5,
		Title:         "A Tale of Two Cities 3",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	{
		ID:            6,
		Title:         "A Tale of Two Cities 4",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	{
		ID:            7,
		Title:         "A Tale of Two Cities 5",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	{
		ID:            8,
		Title:         "A Tale of Two Cities 6",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	{
		ID:            9,
		Title:         "A Tale of Two Cities 7",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	{
		ID:            10,
		Title:         "A Tale of Two Cities 8",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
}
