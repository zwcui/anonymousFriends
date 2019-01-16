/*
@Time : 2019/1/8 下午4:00 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"io/ioutil"
	"strconv"
	"anonymousFriends/models"
	"encoding/json"
	"anonymousFriends/util"
	"errors"
	"sort"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//高德逆地理编码API
func GetRegeocode(longitude float64, latitude float64) (regeoResponse models.RegeoResponse, err error) {
	param := make(map[string]string)
	param["key"] = models.AmapKey
	param["location"] = strconv.FormatFloat(longitude, 'f', 6, 64) + "," + strconv.FormatFloat(latitude, 'f', 6, 64)
	location := "location=" + strconv.FormatFloat(longitude, 'f', 6, 64) + "," + strconv.FormatFloat(latitude, 'f', 6, 64)
	sign := signParam(param)

	url := models.AmapRegeoUrl + "?" + location + "&key=" + models.AmapKey + "&sig="+sign

	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	if err != nil {
		util.Logger.Info("-------client.Do(request)-----err:"+err.Error())
		return regeoResponse, err
	}

	if response != nil {
		defer response.Body.Close()
	}

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			util.Logger.Info("-------body, err := ioutil.ReadAll(response.Body)-----err:"+err.Error())
			return regeoResponse, err
		}

		//util.Logger.Info("-------------------")
		//util.Logger.Info(string(body))
		//util.Logger.Info("-------------------")

		err = json.Unmarshal(body, &regeoResponse)
		if err != nil {
			util.Logger.Info("-------err = json.Unmarshal(body, &regeoResponse)-----err:"+err.Error())
			return regeoResponse, err
		}

	} else {
		util.Logger.Info("-------response.StatusCode != 200-----")
		return regeoResponse, errors.New("response.StatusCode != 200")
	}

	util.Logger.Info("amap:"+strconv.FormatFloat(longitude, 'f', 6, 64) + "," + strconv.FormatFloat(latitude, 'f', 6, 64)+"===>"+regeoResponse.Regeocode.FormattedAddress)

	return regeoResponse, err
}

//根据位置查询实时天气
func GetCurrentWeather(province, city, area string, longitude, latitude float64) (weatherResponse models.WeatherResponse, err error){
	var adcode string
	var location models.Location
	if area != "" {
		location = GetAreaLocation(city, area)
	} else if city != "" {
		location = GetCityLocation(city)
	} else if province != "" {
		location = GetProvinceLocation(province)
	}

	if location.AreaCode == "" {
		regeocodeResponse, err := GetRegeocode(longitude, latitude)
		if err != nil {
			return weatherResponse, err
		}
		adcode = regeocodeResponse.Regeocode.AddressComponent.Adcode
	} else {
		adcode = location.AreaCode
	}

	if adcode == "" {
		util.Logger.Info("-------adcode = ''-----")
		return weatherResponse, errors.New("adcode = ''")
	}

	url := models.AmapWeatherUrl + "?" + "city=" + adcode + "&key=" + models.AmapKey

	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	if err != nil {
		util.Logger.Info("-------client.Do(request)-----err:"+err.Error())
		return weatherResponse, err
	}

	if response != nil {
		defer response.Body.Close()
	}

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			util.Logger.Info("-------body, err := ioutil.ReadAll(response.Body)-----err:" + err.Error())
			return weatherResponse, err
		}


		util.Logger.Info("-------------------")
		util.Logger.Info(string(body))
		util.Logger.Info("-------------------")

		err = json.Unmarshal(body, &weatherResponse)
		if err != nil {
			util.Logger.Info("-------err = json.Unmarshal(body, &weatherResponse)-----err:"+err.Error())
			return weatherResponse, err
		}

	} else {
		util.Logger.Info("-------response.StatusCode != 200-----")
		return weatherResponse, errors.New("response.StatusCode != 200")
	}

	return weatherResponse, err

}








func signParam(param map[string]string) string {
	keys := make([]string, len(param))

	i := 0
	for k := range param {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	strTemp := ""
	for _, key := range keys {
		strTemp = strTemp + key + "=" + param[key] + "&"
	}
	strTemp += models.AmapPrivateKey

	hasher := md5.New()
	hasher.Write([]byte(strTemp))
	md5Str := hex.EncodeToString(hasher.Sum(nil))


	return strings.ToLower(md5Str)
}


