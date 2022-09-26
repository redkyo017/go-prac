package main

import "testing"

func TestCreateNewArticle(t *testing.T) {
	originalLength := len(getAllArticles())

	a, err := createNewArticle("New test title", "New test content")

	allArticles := getAllArticles()
	newLength := len(allArticles)

	if err != nil || newLength != originalLength+1 ||
		a.Title != "New test title" || a.Content != "New test content" {

		t.Fail()
	}
}
