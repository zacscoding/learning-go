package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"log"
	"net/http"
	"strconv"
	strings "strings"
)

type Filter struct {
	res []string
}

func main() {
	bindQueryString()
	// bindQueryStrings()
	// bindCustomValidator()
	//bindHeader()
}

func tests(f Filter) []string {
	f.res = append(f.res, "a")
	f.res = append(f.res, "b")

	return f.res
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
			"p1": p.P1,
			"p2": p.P2,
			"p3": p.P3,
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
		P1 string `form:"p1" binding:"max=4,omitempty"`
		P2 string `form:"p2" binding:"omitempty"`
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
			fmt.Println("failed to bind error")
			h := gin.H{}
			for i, err := range err.(validator.ValidationErrors) {
				m := make(map[string]interface{})
				m["namespace"] = err.Namespace()
				m["field"] = err.Field()
				m["structNamespace"] = err.StructNamespace()
				m["structField"] = err.StructField()
				m["actualTag"] = err.ActualTag()
				m["kind"] = err.Kind()
				m["type"] = err.Type()
				m["value"] = err.Value()
				m["param"] = err.Param()
				h["err-"+strconv.Itoa(i)] = m
			}
			c.JSON(http.StatusBadRequest, h)
			return
		}

		fmt.Println("Check params")
		fmt.Println("==> p1 :", p.P1)
		fmt.Println("==> p2 :", p.P2)

		c.JSON(http.StatusOK, gin.H{
			"p1": p.P1,
			"p2": p.P2,
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.Run(":3500")
}

/*
curl -XGET http://localhost:3500/bind-params?p1=aaaa&p2=1
curl -XGET http://localhost:3500/bind-params?p1=aa
*/
func bindQueryStrings() {
	r := gin.Default()

	type Params struct {
		P1 []string `form:"p1" binding:"omitempty,max=2,dive,startswith=aa|startswith=aA,len=4"`
		P2 []uint   `form:"p2" binding:"gt=0,dive,required"`
		P3 string   `form:"filter,default=all"`
	}
	// form:"filter,default=all"
	//"actualTag": "oneof",
	//        "field": "P3",
	//        "kind": 24,
	//        "namespace": "Params.P3",
	//        "param": "all klay ft",
	//        "structField": "P3",
	//        "structNamespace": "Params.P3",
	//        "type": {},
	//        "value": "a"

	r.GET("/bind-params", func(c *gin.Context) {
		var p Params
		err := c.ShouldBind(&p)
		if err != nil {
			fmt.Println("failed to bind error")
			h := gin.H{}
			for i, err := range err.(validator.ValidationErrors) {
				m := make(map[string]interface{})
				m["namespace"] = err.Namespace()
				m["field"] = err.Field()
				m["structNamespace"] = err.StructNamespace()
				m["structField"] = err.StructField()
				m["actualTag"] = err.ActualTag()
				m["kind"] = err.Kind()
				m["type"] = err.Type()
				m["value"] = err.Value()
				m["param"] = err.Param()
				h["err-"+strconv.Itoa(i)] = m
			}
			c.JSON(http.StatusBadRequest, h)
			return
		}

		fmt.Println("Check params")
		b, _ := json.Marshal(p)
		fmt.Println(string(b))

		//for _, s := range p.P1 {
		//	fmt.Println("==>", s)
		//}

		c.JSON(http.StatusOK, gin.H{
			"p1": p.P1,
			"p2": p.P2,
			"p3": p.P3,
		})
	})
	r.Run(":3500")
}

// x-test-header1:{key1}:{key1Value}:{key2}:{key2Value}
// x-test-header2:{key3}:{key3Value}:{key4}:{key4Value}
type headers struct {
	Header1Values string `header:"x-test-header1" binding:"omitempty,required"`
	Key1          string `binding:"-"`
	Key2          string `binding:"-"`
	Header2Values string `header:"x-test-header2" binding:"omitempty,required"`
	Key3          string `binding:"-"`
	Key4          string `binding:"-"`
}

func (h *headers) parse() error {
	if h.Header1Values != "" {
		pairs := strings.Split(h.Header1Values, ":")
		if len(pairs)%2 != 0 {
			return errors.New(fmt.Sprintf("invalid header1 key/value pairs:%s", h.Header1Values))
		}
		for i := 0; i < len(pairs)-1; i += 2 {
			key := pairs[i]
			value := pairs[i+1]

			switch key {
			case "key1":
				h.Key1 = value
			case "key2":
				h.Key2 = value
			}
		}
	}

	if h.Header2Values != "" {
		pairs := strings.Split(h.Header2Values, ":")
		if len(pairs)%2 != 0 {
			return errors.New(fmt.Sprintf("invalid header2 key/value pairs:%s", h.Header2Values))
		}
		for i := 0; i < len(pairs)-1; i += 2 {
			key := pairs[i]
			value := pairs[i+1]

			switch key {
			case "key3":
				h.Key3 = value
			case "key4":
				h.Key4 = value
			}
		}
	}
	return nil
}

func bindHeader() {
	r := gin.Default()

	r.GET("/bind-headers", func(c *gin.Context) {
		h := headers{}
		err := c.ShouldBindHeader(&h)
		if err != nil {
			c.Error(err)
			return
		}
		err = h.parse()
		if err != nil {
			c.Error(err)
			return
		}
		fmt.Println("header1")
		fmt.Println("key1:", h.Key1)
		fmt.Println("key2:", h.Key2)
		fmt.Println("header2")
		fmt.Println("key3:", h.Key3)
		fmt.Println("key4:", h.Key4)
		fmt.Println(h)
		c.JSON(http.StatusOK, gin.H{
			"result": true,
		})
	})
	r.Run(":3500")
}
