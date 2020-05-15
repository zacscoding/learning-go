package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Meters struct {
	ApiName     string `json:"api_name"`
	Url         string `json:"url"`
	RequestBody string `json:"request_body"`
}

func main() {
	go startMetersServer()

	startApiServer()
}

type Member struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func startApiServer() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	r.Use(gin.Logger())

	r.GET("/members", func(c *gin.Context) {
		// TODO : add measures
		var members []Member
		count := rand.Intn(10)
		for i := 0; i < count; i++ {
			members = append(members, Member{
				Name: "member" + strconv.Itoa(i),
				Age:  rand.Intn(50),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"members": members,
		})
	})

	r.POST("/articles", func(c *gin.Context) {
		// TODO : add measures
		var article Article
		if err := c.ShouldBind(&article); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid article",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"article": article,
		})
	})

	log.Fatal(r.Run(":3000"))
}

func startMetersServer() {
	return
	//r := gin.Default()
	//
	//r.PUT("/meters", func(c *gin.Context) {
	//
	//})
	//
	//log.Fatal(r.Run(":3000"))
}
