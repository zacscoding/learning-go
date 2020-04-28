package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"log"
	"net/http"
)

func main() {
	//bindQueryString()
	bindCustomValidator()
}

var firstValidator validator.Func = func(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(string)
	if ok {
		if len(val) >= 5 {
			return false
		}
	}
	return true
}

func bindCustomValidator() {
	fmt.Println("=================================")
	fmt.Println("Test custom binding validator :)")
	fmt.Println("=================================")
	r := gin.Default()

	type Params struct {
		P1 string `form:"p1" binding:"required"`
		P2 string `binding:"required"`
		P3 string `binding:"custom1"`
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("custom1", firstValidator)
	}

	r.GET("/custom", func(c *gin.Context) {
		var p Params
		if err := c.ShouldBind(&p); err != nil {
			fmt.Println("Error to bind..")
			fmt.Println(err)
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		fmt.Println("Success to bind..")
		fmt.Println("p.P1 :", p.P1)
		fmt.Println("p.P2 :", p.P2)
		fmt.Println("p.P3 :", p.P3)
		c.JSON(http.StatusOK, gin.H{
			"p1" : p.P1,
			"p2" : p.P2,
			"p3" : p.P3,
		})

	})

	log.Fatal(r.Run(":3000"))
	// Tests
	// http://localhost:3000/custom => 400
	// Error #01: Key: 'Params.P1' Error:Field validation for 'P1' failed on the 'required' tag
	//Key: 'Params.P2' Error:Field validation for 'P2' failed on the 'required' tag
	//Key: 'Params.P3' Error:Field validation for 'P3' failed on the 'custom1' tag


}

func bindQueryString() {
	r := gin.Default()

	type Params struct {
		P1 string `form:"p1" binding:"omitempty"`
		P2 string `form:"p2" binding:"p2b"`
	}

	//type Params struct {
	//			Opt
	//			OwnerAddr string `form:"ownerAddress" binding:"omitempty,eth_addr"`
	//		}

	r.GET("/bind-params", func(c *gin.Context) {
		_, ok := c.GetQuery("p1")
		if ok {
			fmt.Println("has p1 query param")
		} else {
			fmt.Println("has no p1 query param")
		}

		_, ok = c.GetQuery("p2")
		if ok {
			fmt.Println("has p2 query param")
		} else {
			fmt.Println("has no p2 query param")
		}

		var p Params
		err := c.ShouldBind(&p)
		if err != nil {
			fmt.Println("failed to bind", err)
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		fmt.Println("Check params")
		fmt.Println("==> p1 :", p.P1)
		fmt.Println("==> p2 :", p.P2)

		c.JSON(http.StatusOK, gin.H {
			"p1" : p.P1,
			"p2" : p.P2,
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "ping",
		})
	})
	r.Run(":3500")
}

