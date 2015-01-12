package weather

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

func New() *Client {
    var c Client
    return &c
}

func (c *Client) SendRequest(city string) *Response {

    c.request(city)
    return &c.Response
}

func (c * Client)request(city string)error{
    url := "http://weather.livedoor.com/forecast/webservice/json/v1?"
    var dist Response

    res,err := http.Get(url + city) ;
    if err != nil {
	return err
    }
    b, err := ioutil.ReadAll(res.Body)
    if err != nil {
	return err
    }

    err = json.Unmarshal(b,&dist)
    if err != nil {
	return err
    }

    c.Response = dist
    return nil
}
