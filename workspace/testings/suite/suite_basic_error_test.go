package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleAssertSuite struct {
	suite.Suite
}

func TestExampleAssertSuite(t *testing.T) {
	suite.Run(t, new(ExampleAssertSuite))
}

func (s *ExampleAssertSuite) TestExample1() {
	assert.Truef(s.T(), false, "invalid result, want : %v, got : %v", true, false)
	//s.T().Errorf("invalid result, want : %v, got %v", true, false)
	s.T().Fail()
	//s.Truef(returnBool(false), "invalid result, want : %v, got : %v", true, false)
}

func returnBool(b bool) bool {
	return b
}

func returnError(msg string) error {
	return errors.New(msg)
}
