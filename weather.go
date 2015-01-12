package weather

type Client struct {
	Response Response
}

type Response struct {
	PublicTime        string             `json:"publictime`
	Link              string             `json:"link"`
	Forecasts         []Forecast         `json:"forecast"`
	Location          Location           `json:"location"`
	PinpointLocations []PinpointLocation `json:"pinpointLocations"`
}

type Location struct {
	City       string `json:"city"`
	Area       string `json:"area"`
	Prefecture string `json:"prefecture"`
}

type Description struct {
	Text       string `json:"text"`
	PublicTime string `json:"publictime"`
}

type Forecast struct {
	DataLabel string `json:"datalabel"`
	Telop     string `json:"telop"`
	Data      string `json:"data"`
	Image     Image  `json:"image"`
}

type Temperature struct {
	Min string `json:"min"`
	Max Max    `json:"max"`
}

type Max struct {
	Celsius    int32     `json:"celsius"`
	Fahrenheit complex64 `json:"fahrenheit"`
}

type Image struct {
	Width  int    `json:"width"`
	Url    string `json:"url"`
	Title  string `json:"title"`
	Height int32  `json:"height"`
}

type Copyright struct {
	Provider Provider `json:"provider"`
	Link     string   `json:"link"`
	Title    string   `json:"title"`
	Image    Image    `json:"image"`
}

type Provider struct {
	Link string `json:"link"`
	Name string `json:"name"`
}
type PinpointLocation struct {
	Link string `json:"link"`
	Name string `json:"name"`
}
