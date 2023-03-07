package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"encoding/base64"
     
	"github.com/fouita/abtesting-api/src/data"
	"github.com/gin-gonic/gin"
)

func engine() *gin.Engine {
	//roots
	router := gin.New()
	router.Use(parseJwtMd())
	router.Use(CORSMd())

	r := router.Group("/data")
	r.POST("/get-analytics/:id", data.GetAnalytics)
	r.POST("/get-filter", data.GetFilter)
	r.POST("/create-test", data.CreateABTest)
	r.POST("/get-ABtest/:id", data.GetABTest)

	return router
}

func main() {
	//server
	r := engine()
	r.Use(gin.Logger())

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func CORSMd() gin.HandlerFunc {
	return func(c *gin.Context) {
		//origin := c.GetHeader("Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
//JWT
func parseJwtMd() gin.HandlerFunc {
	type User struct {
		Name  string
		Email string
		Uid   string
		Role  string
		Ckey  string
	}

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			c.Next()
			return
		}
		s := strings.Split(token, ".")[1]
		s1 := strings.Replace(s, "-", "+", -1)
		bs := strings.Replace(s1, "_", "/", -1)

		str, err := DecodeSegment(bs)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var user User
		json.Unmarshal(str, &user)

		c.Request.Header.Set("X-UID", user.Uid)
		c.Request.Header.Set("X-Email", user.Email)
		c.Request.Header.Set("X-Name", user.Name)
		c.Request.Header.Set("X-Role", user.Role)
		c.Request.Header.Set("X-CKEY", user.Ckey)
		c.Next()
	}
}

func DecodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}
