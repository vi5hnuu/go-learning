package main

import (
	"project-1/author"
)

func main() {
	// Create a new author
	author := author.NewAuthor("John Doe", "john.doe@example.com")

	// Author writes a chapter
	author.WriteChapter("Chapter 1", "This is the content of chapter 1.")

	// Author reviews a chapter
	author.ReviewChapter("Chapter 1", "This is the content of chapter 1.")

	// Author finalizes a chapter
	author.FinalizeChapter("Chapter 1")
}