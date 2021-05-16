package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T){
	r := MakeSampleRouter(BuildDatabaseConfigFromEnvironment())
	testServer := httptest.NewServer(r)
	defer testServer.Close()

	req := &RequestCreateArticle{
		Title: "title",
		Text: "text",
	}

	reqBody,err := json.Marshal(req)
	assert.NoError(t,err)

	resp, err := http.Post(testServer.URL,"application/json",bytes.NewBuffer(reqBody))
	assert.NoError(t,err)
	defer resp.Body.Close()
  
	byteArray, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t,err)

	respBody := &ResponseCreateArticle{}
	err = json.Unmarshal(byteArray,respBody)
	assert.NoError(t,err)

	assert.Equal(t,"title",respBody.Title)
	assert.Equal(t,"text",respBody.Text)

	_,err = uuid.Parse(respBody.ID)
	assert.NoError(t,err,respBody.ID)
}