package cookiesFuncs

import (
	"github.com/gin-gonic/gin"
)

func SetCookies(c *gin.Context, username string,)  {
	/*cookie, err := c.Cookie("gin_cookie")
	fmt.Printf("222")
	if err != nil {
		cookie = "NotSet"
		return err
	}
	fmt.Printf("111")*/
	c.SetCookie("gin_cookie",username,3600,"/","localhost",false,true)
	//fmt.Printf("Cookie value:%s\n",cookie)
}

func CheckCookiesName(c *gin.Context) (string, error) {
	//log.Println("22222")
	cookie, err := c.Request.Cookie("gin_cookie")
	if err != nil {
		return "",err
	}
	return string(cookie.Value), nil
}