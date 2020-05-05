// this package is not working because of `go get` but working if use go mode
// just maintain code snippets for test containers
package containers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"strconv"
	"strings"
	"testing"
	"time"
)

var (
	filePaths  = []string{"mysql-docker-compose.yaml"}
	identifier = strings.ToLower(uuid.New().String())
)

type DatabaseSuite struct {
	suite.Suite
	compose *testcontainers.LocalDockerCompose
	db      *DB
}

func (s *DatabaseSuite) SetupSuite() {
	fmt.Println(">>>> SetupSuite() <<<<")
	fmt.Println("docker-compose up and wait connection")
	s.compose = testcontainers.NewLocalDockerCompose(filePaths, identifier)
	err := s.compose.
		WithCommand([]string{"up", "-d"}).
		Invoke().Error
	if err != nil {
		s.Fail("failed to up docker-compose", err)
	}
	time.Sleep(1 * time.Second)
	// TODO : more efficient way to wait for process
	repeat := 30
	for {
		if repeat <= 0 {
			assert.Fail(s.T(), "failed to connect to database")
		}
		db, err := NewDatabase("root:password@tcp(127.0.0.1:23306)/my_database?charset=utf8&parseTime=True")
		if err == nil {
			s.db = db
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (s *DatabaseSuite) SetupTest() {
	fmt.Println(">>>> SetupTest() <<<<")
	fmt.Println("Drop all table and migrate models")
	s.db.DropTables()
	s.db.AutoMigrate()
}

func (s *DatabaseSuite) TearDownSuite() {
	err := s.compose.Down().Error
	if err != nil {
		s.Fail("failed to down docker-compose", err)
	}
}

func (s *DatabaseSuite) TestSave() {
	// given
	m := &Member{
		Email:    "user1@email.com",
		Username: "user1",
		Password: "password",
	}

	// when
	err := s.db.Save(m)

	// then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), m.ID)
	assert.NotNil(s.T(), m.CreatedAt)
	assert.NotNil(s.T(), m.UpdatedAt)
}

func (s *DatabaseSuite) TestFindByEmail() {
	// given
	m1 := &Member{
		Email:    "user1@email.com",
		Username: "user1",
		Password: "password",
	}
	m2 := &Member{
		Email:    "user2@email.com",
		Username: "user2",
		Password: "password",
	}
	assert.Nil(s.T(), s.db.Save(m1))
	assert.Nil(s.T(), s.db.Save(m2))
	assert.NotNil(s.T(), m1.ID)
	assert.NotNil(s.T(), m2.ID)

	// when
	find, err := s.db.FindByEmail(m1.Email)

	// then
	assert.Nil(s.T(), err)
	assertEqualsMember(s.T(), m1, find)
}

func (s *DatabaseSuite) TestFindByEmailNilIfNotExist() {
	// given
	m1 := &Member{
		Email:    "user1@email.com",
		Username: "user1",
		Password: "password",
	}
	m2 := &Member{
		Email:    "user2@email.com",
		Username: "user2",
		Password: "password",
	}
	assert.Nil(s.T(), s.db.Save(m1))
	assert.Nil(s.T(), s.db.Save(m2))
	assert.NotNil(s.T(), m1.ID)
	assert.NotNil(s.T(), m2.ID)

	// when
	find, err := s.db.FindByEmail("not_exist_email!@email.com")

	// then
	assert.Nil(s.T(), err)
	assert.Nil(s.T(), find)
}

func (s *DatabaseSuite) TestFindAllByUsername() {
	// given
	count := 10
	username := "user"
	for i := 0; i < count; i++ {
		assert.Nil(s.T(), s.db.Save(&Member{
			Email:    "user" + strconv.Itoa(i) + "@email",
			Username: username,
			Password: "password",
		}))
	}
	for i := 0; i < 3; i++ {
		assert.Nil(s.T(), s.db.Save(&Member{
			Email:    "otherUser" + strconv.Itoa(i) + "@email",
			Username: "otherUser",
			Password: "password",
		}))
	}

	// when
	users, err := s.db.FindAllByUsername(username)
	assert.Nil(s.T(), err)
	for i, user := range users {
		// assert username
		assert.Equal(s.T(), username, user.Username)
		// assert order by id desc
		if i != 0 {
			assert.Greater(s.T(), user.ID, users[i-1].ID)
		}
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}

func assertEqualsMember(t *testing.T, expect *Member, actual *Member) {
	assert.Equal(t, expect.ID, actual.ID)
	assert.Equal(t, expect.Email, actual.Email)
	assert.Equal(t, expect.Username, actual.Username)
	assert.Equal(t, expect.Password, actual.Password)
	// TODO : compare time ?

}
