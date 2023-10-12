package main

import (
	"net/http"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/show", showMenu)

	router.Run()
}

func showMenu(c *gin.Context) {
	menu := []string{
		"치킨베이컨볶음밥",
		"김치삼겹살볶음밥",
		"김치치킨볶음밥",
		"해산물볶음밥",
	}
	c.IndentedJSON(http.StatusOK, menu[rand.Intn(len(menu))])
}
