package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/zacscoding/learning-go/workspace/testings/learning-mocks/repo/mocks"
	"testing"
)

func TestRepoForMocking(t *testing.T) {
	repo := new(mocks.Repo)

	// Tests GetNames()
	repo.On("GetNames").Return([]string{"a", "b"}, nil)

	res, err := repo.GetNames()

	assert.Nil(t, err)
	assert.Equal(t, res, []string{"a", "b"})


	repo.On("PutName", "a").Return(true, nil)
	repo.On("PutName", "b").Return(false, nil)

	b, err := repo.PutName("a")
	assert.True(t, b)
	assert.Nil(t, err)
	repo.AssertCalled(t, "PutName","a")
	repo.AssertNotCalled(t, "PutName", "b")

	b, err = repo.PutName("b")
	assert.False(t, b)
	assert.Nil(t, err)
	repo.AssertCalled(t, "PutName","a")
	repo.AssertCalled(t, "PutName", "b")
}