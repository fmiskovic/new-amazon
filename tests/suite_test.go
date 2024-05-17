package tests

import (
	"github.com/fmiskovic/new-amz/internal/testcontainers"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HandlersTestSuite struct {
	suite.Suite
	testDb *testcontainers.TestDB
}

func (s *HandlersTestSuite) SetupSuite() {
	var err error
	s.testDb, err = testcontainers.SetUpDb()
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *HandlersTestSuite) TearDownSuite() {
	s.testDb.Shutdown()
}

func TestHandlersSuite(t *testing.T) {
	if testing.Short() {
		return
	}
	suite.Run(t, new(HandlersTestSuite))
}
