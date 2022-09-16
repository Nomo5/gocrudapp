package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/posts", handleGetList)
	http.HandleFunc("/posts/", handleRequest)

	server.ListenAndServe()
}

//bookの一覧を取得
func handleGetList(w http.ResponseWriter, r *http.Request) {
	var err error 

	books, err := getBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	output, err := json.MarshalIndent(&books, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Write(output)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error 

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//指定されたidのbookを取得
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return 
	}

	book, err := retrieve(id)
	if err != nil {
		return 
	}

	output, err := json.MarshalIndent(&book, "", "\t")
	if err != nil {
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return 
}

//bookの新規作成
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	contentLength := r.ContentLength
	contentBody := make([]byte, contentLength)

	r.Body.Read(contentBody)

	var book Book 
	json.Unmarshal(contentBody, &book)
	err = book.create()
	if err != nil {
		return 
	}
	w.WriteHeader(200)
	return 
}

//bookの更新
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return 
	}

	book, err := retrieve(id)
	if err != nil {
		return 
	}

	contentLength := r.ContentLength
	contentBody := make([]byte, contentLength)
	
	r.Body.Read(contentBody)

	json.Unmarshal(contentBody, &book)
	err = book.update()
	if err != nil {
		return 
	}
	w.WriteHeader(200)
	return 
}

//bookの削除
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return 
	}

	book, err := retrieve(id)
	if err != nil {
		return 
	}

	err = book.delete()
	if err != nil {
		return 
	}
	w.WriteHeader(200)
	return 
}