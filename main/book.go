package main
import "fmt"
type Book struct {
	ID int
	Title string
	Author string
	YearPublished int
}
func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n" +
			"Published:\t%v\n",b.Title, b.Author, b.YearPublished)
}
var books = []Book{
	{
		ID:			1,
		Title: 		"The Hitchhiker's Guide to the Galaxy",
		Author:		"Douglas Adams",
		YearPublished:1979,
	},
	{
		ID:			2,
		Title: 		"The Hobbit",
		Author:		"J.R.R. Tolkien",
		YearPublished:1937,
	},
	{
		ID:			3,
		Title: 		"Non-Probable",
		Author:		"Adam Fawer",
		YearPublished:2000,
	},
	{
		ID:			4,
		Title: 		"Notes From Underground",
		Author:		"Dostoyevski",
		YearPublished:1985,
	},
	{
		ID:			5,
		Title: 		"Chess",
		Author:		"Stefan Zweig",
		YearPublished:1942,
	},
	{
		ID:			6,
		Title: 		"1984",
		Author:		"George Orwell",
		YearPublished:1995,
	},
	{
		ID:			7,
		Title: 		"Animal Farm",
		Author:		"George Orwell",
		YearPublished:1987,
	},
	{
		ID:			8,
		Title: 		"Empathy",
		Author:		"Adam Fawer",
		YearPublished:1996,
	},
	{
		ID:			9,
		Title: 		"Kiralık Konak",
		Author:		"Yakup Kadri Karaosmanoğlu",
		YearPublished:1975,
	},
	{
		ID:			10,
		Title: 		"Araba Sevdası",
		Author:		"Recaizade Mahmut Ekrem",
		YearPublished:1875,
	},
}
