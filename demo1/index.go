package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Response Price BTC-THB
// type Response struct {
// 	Market string  `json:"THB_BTC"`
// 	Price  float64 `json:"last"`
// }

type tickerResponse struct {
	ThbBtc struct {
		ID            int     `json:"id"`
		Last          float64 `json:"last"`
		LowestAsk     float64 `json:"lowestAsk"`
		HighestBid    float64 `json:"highestBid"`
		PercentChange float64 `json:"percentChange"`
		BaseVolume    float64 `json:"baseVolume"`
		QuoteVolume   float64 `json:"quoteVolume"`
		IsFrozen      int     `json:"isFrozen"`
		High24Hr      float64 `json:"high24hr"`
		Low24Hr       float64 `json:"low24hr"`
	} `json:"THB_BTC"`
}
type book struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{
		ID:     "1",
		Name:   "Harry Potter",
		Author: "J.K. Rowling",
		Price:  15.9,
	},
	{
		ID:     "2",
		Name:   "One Piece",
		Author: "Oda Eiichirō",
		Price:  2.99,
	},
	{
		ID:     "3",
		Name:   "demon slayer",
		Author: "koyoharu gotouge",
		Price:  2.99,
	},
}

func main() {
	router := gin.Default()
	router.GET("/getPriceThbBtc", getPriceThbBtc)
	router.GET("/books", getbooks)
	router.Run()
}

//	func getPriceThbBtc(c *gin.Context) {
//	    c.JSON(http.StatusOK, books)
//	}
func getPriceThbBtc(c *gin.Context) {

	response, err := http.Get("https://api.bitkub.com/api/market/ticker?sym=THB_BTC")

	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		os.Exit(1)
	} else {
		responseData, err := ioutil.ReadAll(response.Body)
		response.Body.Close()

		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			log.Fatal(err)

		} else {
			//var resp Response;
			fmt.Println(string(responseData))

			var result tickerResponse
			if err := json.Unmarshal(responseData, &result); err != nil {
				fmt.Println("Can not unmarshal JSON")
			}

			fmt.Println(result)

			c.JSON(http.StatusOK, result)
			return
		}
		c.JSON(http.StatusNotFound, "data not found")
	}

}

func getbooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}
