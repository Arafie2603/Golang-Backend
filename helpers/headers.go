// helpers/headers.go

package helpers

import (
	"net/http"
	_ "github.com/gin-gonic/gin"
)

func ExtractToken(req *http.Request) string {
	token := req.Header.Get("Authorization")
	return token
}
