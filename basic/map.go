package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "jonathan",
		"city": "jakarta",
	}

	fmt.Println(person)

	// Function pada map
	book := make(map[string]string) // => cara lain membuat map
	book["title"] = "Buku Coding"
	book["author"] = "jonathan"
	book["aa"] = "aa"
	fmt.Println(book)

	delete(book, "aa") // => menghapus data pada map
	fmt.Println(book)
}
