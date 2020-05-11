package notification

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Mark struct {
	Content string `json:"content"`
}
type WXMessage struct {
	Msgtype string `json:"msgtype"`
	Markdown Mark `json:"markdown"`
}

func WeiXin(text,WXurl string)(string)  {
	u := WXMessage{
		Msgtype:"markdown",
		Markdown:Mark{Content:text},
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	log.SetPrefix("[DEBUG 2]")
	log.Println(b)
	tr :=&http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	//res,err := http.Post(Ddurl, "application/json", b)
	//resp, err := http.PostForm(url,url.Values{"key": {"Value"}, "id": {"123"}})
	client := &http.Client{Transport: tr}
	res,err  := client.Post(WXurl, "application/json", b)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()
	result,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(result)
}