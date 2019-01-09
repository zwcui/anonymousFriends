/*
@Time : 2019/1/8 下午4:10 
@Author : zwcui
@Software: GoLand
*/
package models

//逆地理编码API服务地址
const AmapRegeoUrl = "https://restapi.amap.com/v3/geocode/regeo"

//天气查询API服务地址
const AmapWeatherUrl = "https://restapi.amap.com/v3/weather/weatherInfo"

//key
const AmapKey = "0dbdcdfff6045db91d251d163604c115"

//签名私钥
const AmapPrivateKey = "c6e172df0c712772136195af95f1af89"

//-------------------amap------------高德逆地理编码--------返回结构体------------------
type RegeoResponse struct {
	Status						string					`description:"status" json:"status"`
	Info						string					`description:"info" json:"info"`
	Infocode					string					`description:"infocode" json:"infocode"`
	Regeocode					Regeocode				`description:"regeocode" json:"regeocode"`
}

type Regeocode struct {
	FormattedAddress			string					`description:"formatted_address" json:"formatted_address"`
	AddressComponent			AddressComponent		`description:"addressComponent" json:"addressComponent"`
	Pois						[]Poi					`description:"pois" json:"pois"`
	Roads						[]Road					`description:"roads" json:"roads"`
	Roadinters					[]Roadinter					`description:"roadinters" json:"roadinters"`
	Aois						[]Aoi					`description:"aois" json:"aois"`
}

type AddressComponent struct {
	Country						string					`description:"country" json:"country"`
	Province					string					`description:"province" json:"province"`
	City						string					`description:"city数组" json:"city"`
	Citycode					string					`description:"citycode" json:"citycode"`
	District					string					`description:"district" json:"district"`
	Adcode						string					`description:"adcode" json:"adcode"`
	Township					[]string				`description:"township" json:"township"`
	Towncode					[]string				`description:"towncode" json:"towncode"`
	//Neighborhood				Neighborhood			`description:"neighborhood" json:"neighborhood"`
	Building					Building				`description:"building" json:"building"`
	//StreetNumber				StreetNumber			`description:"streetNumber" json:"streetNumber"`
	//BusinessAreas				[]string				`description:"businessAreas" json:"businessAreas"`
}

type Neighborhood struct {
	Name						[]string				`description:"name" json:"name"`
	Type						[]string				`description:"type" json:"type"`
}

type Building struct {
	Name						[]string				`description:"name" json:"name"`
	//Type						string					`description:"type" json:"type"`
}

type StreetNumber struct {
	Street						string					`description:"street" json:"street"`
	Number						string					`description:"number" json:"number"`
	Location					string					`description:"location" json:"location"`
	Direction					string					`description:"direction" json:"direction"`
	Distance					string					`description:"distance" json:"distance"`
}

type BusinessArea struct {
	Location					string					`description:"location" json:"location"`
	Name						string					`description:"name" json:"name"`
	Id							string					`description:"id" json:"id"`
}

type Poi struct {
	Id							string					`description:"id" json:"id"`
	Name						string					`description:"name" json:"name"`
	Type						string					`description:"type" json:"type"`
	Tel							string					`description:"tel" json:"tel"`
	Direction					string					`description:"direction" json:"direction"`
	Distance					string					`description:"distance" json:"distance"`
	Location					string					`description:"location" json:"location"`
	Address						string					`description:"address" json:"address"`
	Poiweight					string					`description:"poiweight" json:"poiweight"`
	Businessarea				string					`description:"businessarea" json:"businessarea"`
}

type Road struct {
	Id							string					`description:"id" json:"id"`
	Name						string					`description:"name" json:"name"`
	Direction					string					`description:"direction" json:"direction"`
	Distance					string					`description:"distance" json:"distance"`
	Location					string					`description:"location" json:"location"`
}

type Roadinter struct {
	Direction					string					`description:"direction" json:"direction"`
	Distance					string					`description:"distance" json:"distance"`
	Location					string					`description:"location" json:"location"`
	FirstId						string					`description:"first_id" json:"first_id"`
	FirstName					string					`description:"first_name" json:"first_name"`
	SecondId					string					`description:"second_id" json:"second_id"`
	SecondName					string					`description:"second_name" json:"second_name"`
}

type Aoi struct {
	Id							string					`description:"id" json:"id"`
	Name						string					`description:"name" json:"name"`
	Adcode						string					`description:"adcode" json:"adcode"`
	Location					string					`description:"location" json:"location"`
	Area						string					`description:"area" json:"area"`
	Distance					string					`description:"distance" json:"distance"`
	Type						string					`description:"type" json:"type"`
}

//-------------------amap------------高德逆地理编码--------返回结构体------------------

//-------------------amap------------高德天气查询--------返回结构体------------------
type WeatherResponse struct {
	Status						string					`description:"status" json:"status"`
	Count						string					`description:"count" json:"count"`
	Info						string					`description:"info" json:"info"`
	Infocode					string					`description:"infocode" json:"infocode"`
	Lives						[]Live					`description:"lives" json:"lives"`
}

type Live struct {
	Province					string					`description:"province" json:"province"`
	City						string					`description:"city" json:"city"`
	Adcode						string					`description:"adcode" json:"adcode"`
	Weather						string					`description:"weather" json:"weather"`
	Temperature					string					`description:"temperature" json:"temperature"`
	Winddirection				string					`description:"winddirection" json:"winddirection"`
	Windpower					string					`description:"windpower" json:"windpower"`
	Humidity					string					`description:"humidity" json:"humidity"`
	Reporttime					string					`description:"reporttime" json:"reporttime"`
}

//-------------------amap------------高德天气查询--------返回结构体------------------