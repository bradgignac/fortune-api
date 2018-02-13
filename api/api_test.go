package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradgignac/fortune-api/db"
	"github.com/stretchr/testify/suite"
)

var client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

var fortunes = []string{
	"one fish, two fish",
	"red fish, blue fish",
}

type APITestSuite struct {
	suite.Suite
}

func initializeTestServer(fortunes []string) *httptest.Server {
	db := db.NewDatabase(fortunes)
	api := NewHandler(db)
	return httptest.NewServer(api)
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, &APITestSuite{})
}

func (s *APITestSuite) TestListsAllFortunes() {
	ts := initializeTestServer(fortunes)
	defer ts.Close()

	url := fmt.Sprintf("%s/fortunes", ts.URL)
	res, err := client.Get(url)

	s.Require().Nil(err)
	s.Equal(200, res.StatusCode)
	s.Equal("application/json", res.Header.Get("Content-Type"))

	var data []db.Fortune
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&data)

	s.Require().Nil(err)
	s.Len(data, len(fortunes))
	for _, fortune := range data {
		s.Contains(fortunes, fortune.Data)
	}
}

func (s *APITestSuite) TestGetReturnsFortune() {
	ts := initializeTestServer(fortunes)
	defer ts.Close()

	id := db.ComputeID(fortunes[0])
	url := fmt.Sprintf("%s/fortunes/%s", ts.URL, id)
	res, err := client.Get(url)

	s.Require().Nil(err)
	s.Equal(200, res.StatusCode)
	s.Equal("application/json", res.Header.Get("Content-Type"))

	defer res.Body.Close()
	var data db.Fortune
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&data)

	s.Require().Nil(err)
	s.Equal(id, data.ID)
	s.Equal(fortunes[0], data.Data)
}

func (s *APITestSuite) TestGetReturns404() {
	ts := initializeTestServer(fortunes)
	defer ts.Close()

	url := fmt.Sprintf("%s/fortunes/invalid", ts.URL)
	res, err := client.Get(url)

	s.Require().Nil(err)
	s.Equal(404, res.StatusCode)

	res.Body.Close()
}

func (s *APITestSuite) TestRandomRedirectsToFortune() {
	ts := initializeTestServer(fortunes)
	defer ts.Close()

	url := fmt.Sprintf("%s/random", ts.URL)
	res, err := client.Get(url)

	s.Require().Nil(err)
	s.Equal(302, res.StatusCode)
	s.Contains(res.Header.Get("Location"), "/fortunes")
}

func (s *APITestSuite) TestRandomWithEmptyDB() {
	ts := initializeTestServer([]string{})
	defer ts.Close()

	url := fmt.Sprintf("%s/random", ts.URL)
	res, err := client.Get(url)

	s.Require().Nil(err)
	s.Equal(404, res.StatusCode)
}
