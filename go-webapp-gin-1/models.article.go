package main

import (
	"errors"
	"testing"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticles() []article {
	return articleList
}

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v := range alist {
		article := articleList[i]
		if v.Content != article.Content || v.ID != article.ID || v.Title != article.Title {
			t.Fail()
			break
		}
	}
}

func GetArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
