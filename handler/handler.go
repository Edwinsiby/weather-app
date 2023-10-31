package handler

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/resty.v1"
)

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func Weather(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	city := c.PostForm("city")

	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":     city,
			"appid": apiKey,
		}).
		Get("https://api.openweathermap.org/data/2.5/weather")

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": "An error occurred while fetching weather data",
		})
		return
	}

	if resp.StatusCode() != 200 {
		c.JSON(resp.StatusCode(), gin.H{
			"error": "Failed to fetch weather data",
		})
		return
	}

	var weatherData WeatherData

	if err := json.Unmarshal(resp.Body(), &weatherData); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": "Error processing weather data",
		})
		return
	}

	c.HTML(200, "weather.html", gin.H{
		"city":        city,
		"temperature": weatherData.Main.Temp,
		"description": weatherData.Weather[0].Description,
	})
}
