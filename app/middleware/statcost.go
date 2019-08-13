package middleware

func StacCost() gin.HandlerFunc{
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		log.Println(cost)
	}
}