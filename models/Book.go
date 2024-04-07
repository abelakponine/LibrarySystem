package models

type BookModel struct {
	Id            uint
	Title         string
	Author        string
	DatePublished string
	YearPublished uint
	SerialNumber  string
}

func GetBookList() *[]BookModel {
	return &Library
}
func AddToLibrary(Library *[]BookModel, book BookModel) *[]BookModel {
	newList := append(*Library, book)
	return &newList
}
func CreateBook(book *BookModel) *[]BookModel {
	Library = append(Library, *book)
	return &Library
}
