package seed

import (
	"fmt"
	"testing"

	"github.com/ag7if/go-files"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v3"
)

type TestData struct {
	Units   []Unit   `yaml:"units"`
	Members []Member `yaml:"members"`
}

type Unit struct {
	CharterNumber string `yaml:"charter_number"`
	Kind          string `yaml:"kind"`
	Category      string `yaml:"category"`
	Address       string `yaml:"address"`
	City          string `yaml:"city"`
}

type Member struct {
	CAPID     uint   `yaml:"capid"`
	LastName  string `yaml:"last_name"`
	FirstName string `yaml:"first_name"`
}

type SeedTestSuite struct {
	suite.Suite
	data *TestData
}

func (s *SeedTestSuite) SetupSuite() {
	f, err := files.NewFile("test.yaml")
	if err != nil {
		panic(err)
	}

	b, err := f.ReadFile()
	if err != nil {
		panic(err)
	}

	d := new(TestData)
	err = yaml.Unmarshal(b, d)
	if err != nil {
		panic(err)
	}

	s.data = d
}

func (s *SeedTestSuite) TestEnsureUniqueCAPID() {
	capids := make(map[uint]bool)
	for _, m := range s.data.Members {
		_, ok := capids[m.CAPID]
		assert.False(s.T(), ok, fmt.Sprintf("CAPID %d has been duplicated", m.CAPID))
		capids[m.CAPID] = true
	}
}

func TestSeedTestSuite(t *testing.T) {
	suite.Run(t, new(SeedTestSuite))
}
