package handlers

import "net/http"

func HandleAllHttpFuncs() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/article", PostArticleHandler)
	http.HandleFunc("/article/list", ArticleListHandler)
	http.HandleFunc("/article/1", ArticleDetailHandler)
	http.HandleFunc("/article/nice", PostNiceHandler)
	http.HandleFunc("/comment", PostCommentHandler)
}
