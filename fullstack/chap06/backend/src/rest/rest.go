package rest

import "github.com/gin-gonic/gin"

func RunAPI(address string) error {
	// Gin 엔진
	r := gin.Default()
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
	return r.Run(address)
}


