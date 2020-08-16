package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//WeatherInfoJSON is a struct expression.
type WeatherInfoJSON struct {
	Desc   string
	Status float64
	Data   DataObject
}

//DataObject is a struct expression.
type DataObject struct {
	WenDu     string
	GanMao    string
	Forecast  []ForecastsObject
	Yesterday YesterdayObject
	Aqi       string
	City      string
}

//ForecastsObject is a struct expression.
type ForecastsObject struct {
	FengXiang string
	FengLi    string
	High      string
	Type      string
	Low       string
	Date      string
}

//YesterdayObject is a struct expression.
type YesterdayObject struct {
	Fl   string
	Fx   string
	High string
	Type string
	Low  string
	Date string
}

//GetWeather is used to get weather.
func GetWeather(city string) {

	str := "http://wthrcdn.etouch.cn/weather_mini?city=" + city

	resp, err := http.Get(str) //使用get方法访问
	if err != nil {
		fmt.Println("json err: ", err)
	}

	defer resp.Body.Close()
	input, err1 := ioutil.ReadAll(resp.Body) //读取流数据
	if err1 != nil {
		fmt.Println("json err1: ", err1)
	}

	var jsonWeather WeatherInfoJSON
	err2 := json.Unmarshal(input, &jsonWeather) //解析json数据
	if err2 != nil {
		fmt.Println("json err2: ", err2)
	}

	if len(jsonWeather.Desc) != 0 { //判断有无解析数据

		fmt.Println(jsonWeather.Desc, jsonWeather.Status, jsonWeather.Data.City)
		fmt.Printf("今日温度：%s 提醒：%s\n", jsonWeather.Data.WenDu, jsonWeather.Data.GanMao)
		fmt.Print("从今日起最近五天的天气情况：\n")

		for i := 0; i < 5; i++ {

			fmt.Printf("日期：%s 最低温：%s 最高温：%s ", jsonWeather.Data.Forecast[i].Date, jsonWeather.Data.Forecast[i].Low, jsonWeather.Data.Forecast[i].High)
			fmt.Printf(" 天气类型：%s 风力：%s 风向：%s\n", jsonWeather.Data.Forecast[i].Type, jsonWeather.Data.Forecast[i].FengXiang, jsonWeather.Data.Forecast[i].FengLi)

		}

	}
}

func main() {

	GetWeather("龙泉驿")

}
