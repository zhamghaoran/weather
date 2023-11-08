package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

var client http.Client

func init() {
	client = http.Client{}
}
func main() {

	engine := gin.Default()

	engine.GET("/api/district", func(context *gin.Context) {
		keywords := context.Query("keywords")
		key := context.Query("key")
		url := fmt.Sprintf("https://restapi.amap.com/v3/config/district?keywords=%s&subdistrict=1&key=%s", keywords, key)
		log.Println(url)
		res, err := client.Get(url)
		if err != nil {
			log.Println(err)
		}
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(bytes))
		context.String(http.StatusOK, string(bytes))
	})

	engine.GET("/api/weather", func(context *gin.Context) {
		city := context.Query("city")
		url := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?city=%s&key=64f1e991489ad70e8dce7ce266d1661a", city)
		log.Println(url)
		res, err := client.Get(url)
		if err != nil {
			log.Println(err)
		}
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
		}
		context.String(http.StatusOK, string(bytes))

	})
	engine.Run(":8080")

}
