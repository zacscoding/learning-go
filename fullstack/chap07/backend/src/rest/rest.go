package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	// Gin 엔진
	r := gin.Default()
	//r.Use(MyCustomMiddleWare(), MyCustomMiddleWare2()) // middleware 추가
	// Handler 생성
	h, _ := NewHandler()

	// 상품목록
	r.GET("/products", h.GetProducts)
	// 프로모션 목록
	r.GET("/promos", h.GetPromos)

	/*
		// 사용자 로그인 POST 요청
		r.POST("/users/signin", h.SignIn)
		// 사용자 추가 POST 요청
		r.POST("/users", h.AddUser)
		// 사용자 로그아웃 POST 요청
		r.POST("/user/:id/signout", h.SignOut)
		// 구매 목록 조회
		r.GET("/user/:id/orders", h.GetOrders)
		// 결제 POST 요청
		r.POST("/users/charge", h.Charge)
	*/
	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("/", h.AddUser)
	}

	// 서버 시작
	// return r.RunTLS(address, "../cert.pem", "../key.pem") // TLS
	return r.Run(address)
}

// Temporary for middleware
// Output if r.Use(MyCustomMiddleWare(), MyCustomMiddleWare2()) // middleware 추가
// ****************Before****************
//v : 123
//==========Before==========
//==========After==========
//****************After****************
//status : 200

func MyCustomMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ================== Before
		// 요청을 처리하기 전에 실행할 코드
		// 예제 변수 설정
		c.Set("v", 123)
		// c.Get("v")를 호출하면 변수 값을 확인할 수 있음
		fmt.Println("****************Before****************")
		v, _ := c.Get("v")
		fmt.Println("v :", v)

		// 요청 처리 로직 실행
		c.Next()

		// ================== After
		status := c.Writer.Status()
		fmt.Println("****************After****************")
		fmt.Println("status :", status)
	}
}

func MyCustomMiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("==========Before==========")
		c.Next()
		fmt.Println("==========After==========")
	}
}
