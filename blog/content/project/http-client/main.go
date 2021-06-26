package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetExporterMetrics(rawUrl string, paramMap map[string]string, headMap map[string]string) (text){
	Url, err := url.Parse(rawUrl)
	if err != nil {
		return
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = GetParam(paramMap).Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		fmt.Println("GET request failed :", err)
		return err
	}

	req.Header.Add("name", "zhaofan")
	req.Header.Add("age", "3")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("错误:发送请求", err)
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func GetParam(paramMap map[string]string) *url.Values {
	params := url.Values{}
	if paramMap == nil || len(paramMap) == 0 {
		return &params
	}
	for k, v := range paramMap {
		params.Set(k, v)
	}
	return &params
}

