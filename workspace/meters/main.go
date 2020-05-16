package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	gin.SetMode(gin.DebugMode)
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

	r.Use(gin.Logger())

	reporter := NewReporter("http://localhost:8000/meters")

	reportCallSpan := func(c *gin.Context) {
		reporter.Send(ApiCallSpan{
			Url:      c.Request.URL.String(),
			FullPath: c.FullPath(),
			Headers:  c.Request.Header.Clone(),
		})
	}

	r.GET("/members", func(c *gin.Context) {
		reportCallSpan(c)

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
		reportCallSpan(c)

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
	r := gin.Default()
	r.Use(gin.Logger())

	r.POST("/meters", func(c *gin.Context) {
		var spans []ApiCallSpan
		if err := c.ShouldBindJSON(&spans); err != nil {
			fmt.Println("failed to bind spans. ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid span",
			})
			return
		}

		fmt.Println("Success to receive meters :", len(spans))
		for _, span := range spans {
			b, _ := json.Marshal(span)
			fmt.Println(string(b))
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	log.Fatal(r.Run(":8000"))
}
