package repo

import (
	"blynker/internal/model"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type RepoSuite struct {
	suite.Suite
}

func TestRepoRun(t *testing.T) {
	suite.Run(t, new(RepoSuite))
}

func (s *RepoSuite) TestRepo_Save() {
	data := model.Sensor{
		Temperature: 12,
		Light:       12,
		Movement:    false,
		UpdatedAt:   time.Time{},
	}
	testRepo := new(CSVRepo)
	err := testRepo.SaveData(&data)
	s.Require().NoError(err)
}
