package rest

import (
	"book-management/pkg/apis"
	"book-management/pkg/book-cli/cmd/pkg/options"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// handleBookModifications handles the book modifications on the path /api/v1/books/
func (s *BookServer) handleBookModifications(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("received new request. URL: %v Method: %v\n", req.URL, req.Method)

	switch req.Method {
	case options.Create.String():
		s.CreateBook(res, req)
		break
	}
}

// CreateBook parses the request body and passes the object to the database driver
func (s *BookServer) CreateBook(res http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		http.Error(res, apis.NewError(http.StatusBadRequest, fmt.Errorf("error while reading request body: %v", err)).JSON(), http.StatusBadRequest)
		return
	}

	book := &apis.Book{}
	err = json.Unmarshal(reqBody, book)

	if err != nil {
		http.Error(res, apis.NewError(http.StatusBadRequest, fmt.Errorf("error while unmarshaling request body: %v", err)).JSON(), http.StatusBadRequest)
		return
	}

	msg, err := s.db.CreateBook(book)

	if err != nil {
		http.Error(res, apis.NewError(http.StatusInternalServerError, fmt.Errorf("error while creating book: %v", err)).JSON(), http.StatusInternalServerError)
		return
	}
	res.Write([]byte(apis.NewSuccess(msg).JSON()))
}
