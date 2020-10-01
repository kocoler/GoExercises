package redirect

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func RedirectToMain(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://localhost:9090/home")
	log.Println("跳转")
	time.Sleep(100 * time.Microsecond)
	return
}
