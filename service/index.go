package service

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	index, err := template.ParseFiles("index.html", "view/chat/head.html")
	if err != nil {
		panic(err)
	}
	index.Execute(c.Writer, "index")

}

func ToRegister(c *gin.Context) {
	t, err := template.ParseFiles("view/user/register.html")
	if err != nil {
		panic(err)
	}
	t.Execute(c.Writer, "register")
}
