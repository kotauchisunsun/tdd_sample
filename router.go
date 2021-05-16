package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func MakeSampleRouter(c *DataBaseConfig) *mux.Router {
	r := mux.NewRouter()
	db, err := sql.Open("mysql", c.BuildDataSourcename())
	
	if err != nil {
		log.Fatal(err)
	}

    r.HandleFunc("/", func (w http.ResponseWriter,r *http.Request){

		byteArray, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
	
		reqBody := &RequestCreateArticle{}
		err = json.Unmarshal(byteArray,reqBody)
		if err != nil {
			log.Fatal(err)
		}

		ins, err := db.Prepare("INSERT INTO article(id,title,text) VALUES(?,?,?)")
		if err != nil {
			log.Fatal(err)
		}

		id := uuid.New().String()
		if err != nil {
			log.Fatal(err)
		}

		title := reqBody.Title
		text  := reqBody.Text

		_, err = ins.Exec(id, title, text)
		if err != nil {
			log.Fatal(err)
		}

		respBody := &ResponseCreateArticle{
			ID: id,
			Title: title,
			Text: text,
		}

		resp,err := json.Marshal(respBody)
		if err != nil{
			log.Fatal(err)
		}
		w.Write(resp)
	})

	return r
}