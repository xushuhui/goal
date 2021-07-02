package middleware

import (
	"fmt"
	"github.com/xushuhui/goal/internal"

	"log"

	"net/http"
	"runtime"
	"strings"
)

func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(4, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTrace:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}
func Recovery() internal.HandlerFunc {
	return func(c *internal.Context) error {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
		}()

		c.Next()
		return nil
	}
}
