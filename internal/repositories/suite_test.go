package repositories

import (
	"github.com/fmiskovic/new-amz/internal/testcontainers"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RepositoryTestSuite struct {
	suite.Suite
	testDb *testcontainers.TestDB
}

func (s *RepositoryTestSuite) SetupSuite() {
	var err error
	s.testDb, err = testcontainers.SetUpDb()
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *RepositoryTestSuite) TearDownSuite() {
	s.testDb.Shutdown()
}

func TestRepositorySuite(t *testing.T) {
	if testing.Short() {
		return
	}
	suite.Run(t, new(RepositoryTestSuite))
}
