package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

func (suite *ExampleTestSuite) SetupSuite() {
	fmt.Println("SetupSuite() is called")
}

func (suite *ExampleTestSuite) SetupTest() {
	fmt.Println("SetupTest() is called")
	suite.VariableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite() is called")
}

func (suite *ExampleTestSuite) TearDownTest() {
	fmt.Println("TearDownTest() is called")
}

func (suite *ExampleTestSuite) TestExample() {
	fmt.Println("TestExample() is called")
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func (suite *ExampleTestSuite) TestExample2() {
	fmt.Println("TestExample2() is called")
	fmt.Println("Name ::", suite.T().Name())
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func tempCall() (string, error) {
	return "", errors.New("force error")
}
