## goal
go web framework

## [Guide](https://blog.phpst.cn)

### Installation

```sh

go get github.com/xushuhui/goal
```

### Example

```go
package main

import (
    "net/http"
    "github.com/xushuhui/goal"
 
)

func main() {
	r := goal.Default()

	r.GET("/", func(c *Context) error {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
		return nil
	})
	r.GET("/hello", func(c *Context) error {

		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		return nil
	})
    r.Run(":8000")
}


```
- config


### 微信公众号
![扫码关注](https://tvax4.sinaimg.cn/large/a616b9a4gy1grl9d1rdpvj2076076wey.jpg)
