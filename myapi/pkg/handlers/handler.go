package handlers

import (
	"fmt"
	"io"
	"net/http"
)

const InvalidMethod = "Invalid request method."

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Hello Home Page\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Post article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Article list\n")
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	articleID := 1
	resString := fmt.Sprintf("Article detail for article ID %d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Post nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Post comment...\n")
}
