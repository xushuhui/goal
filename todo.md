### TODO

- [x] mysql
- [x] redis
- [x] validate 验证信息
- [x] json web token
- [x] cors
- [x] 自定义响应
- [x] 错误码和提示信息
- [x] 统一错误处理
- [x] log 日志
- [x] config 配置
- [x] 请求参数签名
- [x] tests
- [x] 文件上传
- [ ] ~~alarm 服务器错误警告（邮箱，钉钉）(业务层)~~
- [ ] ~~消息中心~~(业务层)

- [ ] websocket server
- [ ] tcp server

- [ ] 热更新
- [ ] 开发工具

//func main() {
//	r := goal.Default()
//
//	r.GET("/", func(c *Context) error {
//		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
//		return nil
//	})
//	r.GET("/hello", func(c *Context) error {
//
//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
//		return nil
//	})
//
//	r.GET("/hello/:name", func(c *Context) error {
//
//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//		return nil
//	})
//	v1 := r.Group("/v1")
//	v1.Use(Logging())
//	{
//		v1.GET("/", func(c *Context) error {
//			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
//			return nil
//		})
//
//		v1.GET("/hello", func(c *Context) error {
//			c.JSON(http.StatusOK, H{"x": "x"})
//
//			return nil
//		})
//		v1.GET("/hello/:name", func(c *Context) error {
//			// expect /hello/geektutu
//			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//			return nil
//		})
//	}
//	r.Run(":8000")
//}
