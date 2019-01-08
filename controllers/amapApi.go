/*
@Time : 2019/1/8 下午4:00 
@Author : zwcui
@Software: GoLand
*/
package controllers
//
//import (
//	"net/http"
//	"io/ioutil"
//	"fmt"
//	"strconv"
//	"anonymousFriends/models"
//	"encoding/json"
//)
//
//func geocode(longitude float64, latitude float64) (err error) {
//	location := "location=" + strconv.FormatFloat(longitude, 'f', 6, 64) + "," + strconv.FormatFloat(latitude, 'f', 6, 64)
//
//	url := models.RegeoUrl + "?" + location + "&key=" + models.AmapKey
//
//
//
//	client := &http.Client{}
//	request, _ := http.NewRequest("GET", url, nil)
//	response, err := client.Do(request)
//	if err != nil {
//
//	}
//
//	defer response.Body.Close()
//
//	if response.StatusCode == 200 {
//
//
//		body, err := ioutil.ReadAll(response.Body)
//		if err != nil {
//
//		}
//
//		response := hwTokenResponse{}
//		err = json.Unmarshal(body, &response)
//		if err != nil {
//			return
//		}
//
//
//
//
//	} else {
//
//
//
//
//	}
//
//
//
//
//
//}
