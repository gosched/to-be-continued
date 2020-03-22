package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ConfigWeatherService()

	// 一般天氣預報
	// 今明 36 小時天氣預報
	// service1 := ServiceChannel(1)
	// response1 := <-service1
	// if response1.OK {
	// 	fmt.Println(string(response1.Data))
	// }
	// fmt.Println()

	// 鄉鎮天氣預報
	// 臺中市未來2天天氣預報
	service2 := ServiceChannel(2)
	response2 := <-service2
	if response2.OK {
		fmt.Println(string(response2.Data))
		// file, err := os.Create("weather.json")
		// if err != nil {
		// 	log.Fatalln()
		// }
		// defer file.Close()
		// file.Write(response2.Data)
	}
	fmt.Println()

	// 鄉鎮天氣預報
	// 臺中市未來1週天氣預報
	// service3 := ServiceChannel(3)
	// response3 := <-service3
	// if response3.OK {
	// 	fmt.Println(string(response3.Data))
	// }
	// fmt.Println()

}

// https://opendata.cwb.gov.tw/index

// CWB-156B0091-9D2C-4D20-BBAD-A338B04AED4E
var weatherServiceURL map[int]string
var authorizationToken = "?Authorization=CWB-156B0091-9D2C-4D20-BBAD-A338B04AED4E"
var apiURL = "https://opendata.cwb.gov.tw/api/v1/rest/datastore/"

// ConfigWeatherService .
func ConfigWeatherService() {
	weatherServiceURL = make(map[int]string)
	weatherServiceURL[1] = apiURL + "F-C0032-001" + authorizationToken
	weatherServiceURL[2] = apiURL + "F-D0047-073" + authorizationToken
	weatherServiceURL[3] = apiURL + "F-D0047-075" + authorizationToken
}

// WeatherServiceResponse .
type WeatherServiceResponse struct {
	OK    bool
	Error error
	Data  []byte
}

// ServiceChannel .
func ServiceChannel(id int) <-chan *WeatherServiceResponse {
	c := make(chan *WeatherServiceResponse)
	go func() {
		for {
			r := &WeatherServiceResponse{}
			url := weatherServiceURL[id]

			response, err := http.Get(url)
			if err != nil {
				r.OK = false
				r.Data = nil
				c <- r
				continue
			}

			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				r.OK = false
				r.Data = nil
				c <- r
				continue
			}

			r.OK = true
			r.Data = data
			c <- r
		}
	}()
	return c
}
