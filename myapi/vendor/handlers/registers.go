package handlers

import "net/http"

func HandleAllHttpFuncs() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", articleListHandler)
	http.HandleFunc("/article/1", articleDetailHandler)
	http.HandleFunc("/article/nice", postNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)
}
