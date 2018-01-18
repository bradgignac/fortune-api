package fortune

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var fortunes = []string{
	"one fish, two fish",
	"red fish, blue fish",
}

type DatabaseTestSuite struct {
	suite.Suite
}

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, &DatabaseTestSuite{})
}

func (s *DatabaseTestSuite) TestListReturnsAllFortunes() {
	db := NewDatabase(fortunes)
	list := db.List()

	data := make([]string, db.Count())
	for i, f := range list {
		data[i] = f.Data
	}

	s.Len(list, len(fortunes))
	for _, f := range fortunes {
		s.Contains(data, f)
	}
}

func (s *DatabaseTestSuite) TestGetReturnsFortuneWithGivenID() {
	id := ComputeID(fortunes[0])
	db := NewDatabase(fortunes)
	get, err := db.Get(id)

	s.Nil(err)
	s.Equal(id, get.ID)
	s.Equal(fortunes[0], get.Data)
}

func (s *DatabaseTestSuite) TestGetReturnsErrorForMissingID() {
	db := NewDatabase([]string{})
	get, err := db.Get("missing")

	s.Nil(get)
	s.Equal(err, MissingFortuneError)
}

func (s *DatabaseTestSuite) TestReturnsRandomID() {
	db := NewDatabase(fortunes)
	random, err := db.Random()

	ids := make([]string, db.Count())
	for i, f := range db.List() {
		ids[i] = f.ID
	}

	s.Nil(err)
	s.Contains(ids, random)
}

func (s *DatabaseTestSuite) TestRandomForEmptyDatabase() {
	db := NewDatabase([]string{})
	random, err := db.Random()

	s.Equal(random, "")
	s.Equal(err, EmptyDatabaseError)
}

func (s *DatabaseTestSuite) TestCountReturnsDatabaseLength() {
	fortunes := []string{
		"one fish, two fish",
		"red fish, blue fish",
		"one more for good measure",
	}
	db := NewDatabase(fortunes)

	s.Equal(3, db.Count())
}
