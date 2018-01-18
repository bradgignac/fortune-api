package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradgignac/fortune-api/fortune"
	"github.com/stretchr/testify/suite"
)

var fortunes = []string{
	"one fish, two fish",
	"red fish, blue fish",
}

type APITestSuite struct {
	suite.Suite
	api *Handler
	db  *fortune.Database
	ts  *httptest.Server
}

func TestAPITestSuite(t *testing.T) {
	db := fortune.NewDatabase(fortunes)
	api := NewHandler(db)
	suite.Run(t, &APITestSuite{db: db, api: api})
}

func (s *APITestSuite) SetupSuite() {
	s.ts = httptest.NewServer(s.api)
}

func (s *APITestSuite) TearDownSuite() {
	s.ts.Close()
}

func (s *APITestSuite) TestListsAllFortunes() {
	url := fmt.Sprintf("%s/fortunes", s.ts.URL)
	res, err := http.Get(url)

	s.Require().Nil(err)
	s.Equal(200, res.StatusCode)
	s.Equal("application/json", res.Header.Get("Content-Type"))

	var data []fortune.Fortune
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
	id := fortune.ComputeID(fortunes[0])
	url := fmt.Sprintf("%s/fortunes/%s", s.ts.URL, id)
	res, err := http.Get(url)

	s.Require().Nil(err)
	s.Equal(200, res.StatusCode)
	s.Equal("application/json", res.Header.Get("Content-Type"))

	defer res.Body.Close()
	var data fortune.Fortune
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&data)

	s.Require().Nil(err)
	s.Equal(id, data.ID)
	s.Equal(fortunes[0], data.Data)
}

func (s *APITestSuite) TestRandomRedirectsToFortune() {
	url := fmt.Sprintf("%s/random", s.ts.URL)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Get(url)

	s.Require().Nil(err)
	s.Equal(302, res.StatusCode)
	s.Contains(res.Header.Get("Location"), "/fortunes")
}
