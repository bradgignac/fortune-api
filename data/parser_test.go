package data

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, &ParserTestSuite{})
}

func (s *ParserTestSuite) TestEmptyData() {
	reader := strings.NewReader("")
	db, err := Parse(reader)

	s.Equal(0, db.Count())
	s.Nil(err)
}

func (s *ParserTestSuite) TestSingleLineFortune() {
	reader := strings.NewReader(`hi there`)
	db, err := Parse(reader)
	list := db.List()

	s.Equal(db.Count(), 1)
	s.Equal(list[0].Data, "hi there")
	s.Nil(err)
}

func (s *ParserTestSuite) TestMultLineFortune() {
	reader := strings.NewReader(`hi there
  - taco from trello
`)
	db, err := Parse(reader)
	list := db.List()

	s.Equal(db.Count(), 1)
	s.Equal(list[0].Data, "hi there\n  - taco from trello")
	s.Nil(err)
}

func (s *ParserTestSuite) TestMultipleFortunes() {
	reader := strings.NewReader(`hi there
%
one more
  - with attribution
%
last one`)
	db, err := Parse(reader)
	list := db.List()

	data := make([]string, db.Count())
	for i, f := range list {
		data[i] = f.Data
	}

	s.Len(list, 3)
	s.Contains(data, "hi there")
	s.Contains(data, "one more\n  - with attribution")
	s.Contains(data, "last one")
	s.Nil(err)
}
