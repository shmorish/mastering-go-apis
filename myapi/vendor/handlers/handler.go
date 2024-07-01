package handlers

import (
	"fmt"
	"io"
	"net/http"
)

const InvalidMethod = "Invalid request method."

func homeHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Hello Home Page\n")
}

func postArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Post article...\n")
}

func articleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Article list\n")
}

func articleDetailHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	articleID := 1
	resString := fmt.Sprintf("Article detail for article ID %d\n", articleID)
	io.WriteString(w, resString)
}

func postNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Post nice...\n")
}

func postCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, InvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "Post comment...\n")
}
