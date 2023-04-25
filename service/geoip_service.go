package service

import (
	"fmt"
	"io/ioutil"
	"log"

	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/oschwald/geoip2-golang"
)

// 需要满足以上要求
type Response struct {
	Country   string  `json:"country"`
	Province  string  `json:"province"`
	City      string  `json:"city"`
	ISP       string  `json:"isp"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TimeZone  string  `json:"timezone"`
}

// 第三方
type ThirdResponse struct {
	Address   string  `json:"address"`
	City      string  `json:"city"`
	Ip        string  `json:"ip"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GeoipInfor(para string, typev int, requesturl string) ThirdResponse {
	var res ThirdResponse
	if typev == 2 {
		db, err := geoip2.Open("/go/src/GeoLite2-City.mmdb")
		//db, err := geoip2.Open("F:\\Project-cz\\go\\workspace\\src\\tools_go\\GeoLite2-City.mmdb")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		// If you are using strings that may be invalid, check that ip is not nil
		ip := net.ParseIP(para)
		record, err := db.City(ip)
		if err != nil {
			log.Fatal(err)
		}
		var build strings.Builder
		if record.Country.Names != nil {
			build.WriteString(record.Country.Names["zh-CN"])
		}

		if record.Subdivisions != nil {
			build.WriteString(record.Subdivisions[0].Names["zh-CN"])
		}

		if record.City.Names != nil {
			build.WriteString(record.City.Names["zh-CN"])
			res.City = record.City.Names["zh-CN"]
		}

		if record.Location.Longitude != 0 {
			res.Latitude = record.Location.Latitude
			res.Longitude = record.Location.Longitude
		}
		res.Address = build.String()
	} else {
		if requesturl == "" {
			// 配置一个默认的 请求地址第三方免费开放API
			requesturl = "http://xxx.xxxxx.xxxx/xxx.xxxx"
		}
		name := url.Values{}
		name.Set("ip", para)
		param := name.Encode()
		url := fmt.Sprintf("%s?%s", requesturl, param)
		resp, err := http.Get(url)
		if err != nil {
			//log.Println(err)
			log.Fatal(err)
			//return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		res.Address = mahonia.NewDecoder("gbk").ConvertString(strings.TrimSpace(string(body[:])))

	}
	//db, err := geoip2.Open("/go/src/GeoLite2-City.mmdb")
	// db, err := geoip2.Open("F:\\Project-cz\\go\\workspace\\src\\tools_go\\GeoLite2-City.mmdb")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// // If you are using strings that may be invalid, check that ip is not nil
	// ip := net.ParseIP(para)
	// record, err := db.City(ip)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var res Response
	// if record.City.Names != nil {
	// 	res.City = record.City.Names["zh-CN"]
	// }

	// if record.Subdivisions != nil {
	// 	res.Province = record.Subdivisions[0].Names["zh-CN"]
	// }

	// if record.Country.Names != nil {
	// 	res.Country = record.Country.Names["zh-CN"]
	// }

	// if record.Location.Longitude != 0 {
	// 	res.Latitude = record.Location.Latitude
	// 	res.Longitude = record.Location.Longitude
	// 	res.TimeZone = record.Location.TimeZone
	// }

	res.Ip = para
	return res
}
