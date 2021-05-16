package main

type RequestCreateArticle struct{
	Title string `json: "title"`
	Text  string `json: "text"`
}

type ResponseCreateArticle struct{
	ID    string `json: "ID"`
	Title string `json: "title"`
	Text  string `json: "text"`
}