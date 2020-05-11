package notification

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"encoding/json"
	"alterCenter/config"
)

//腾讯短信接口sha256编码
func getSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//腾讯短信接口消息格式
type Mobiles struct{
	Mobile string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

type TXmessage struct {
	Ext string `json:"ext"`
	Extend string `json:"extend"`
	Params []string `json:"params"`
	Sig string `json:"sig"`
	Sign string `json:"sign"`
	Tel []Mobiles `json:"tel"`
	Time int `json:"time"`
	Tpl_id int `json:"tpl_id"`
}

//腾讯短信子程序
func Duanxin(text string,mobile string)(string)  {
	strAppKey:=config.MessageConfig.Appkey
	tpl_id:=config.MessageConfig.Tplid
	sdkappid:=config.MessageConfig.Sdkappid
	//腾讯短信接口算法部分
	//mobile格式:"15395105573,16619875573"
	TXmobile:=Mobiles{}
	TXmobiles:=[]Mobiles{}
	mobiles:=strings.Split(mobile,",")
	for _,m:=range mobiles {
		TXmobile.Mobile=m
		TXmobile.Nationcode="86"
		TXmobiles=append(TXmobiles,TXmobile )
	}
	strRand := "7226249334"
	strTime := strconv.FormatInt(time.Now().Unix(),10)
	intTime,_:=strconv.Atoi(strTime)
	tplId,_ := strconv.Atoi(tpl_id)
	sig := getSha256Code("appkey="+strAppKey+"&random="+strRand+"&time="+strTime+"&mobile="+mobile)
	TXurl:="https://yun.tim.qq.com/v5/tlssmssvr/sendmultisms2?sdkappid="+sdkappid+"&random="+strRand
	u := TXmessage{
		Ext:"",
		Extend:"",
		Params:[]string{text},
		Sig:sig,
		Tel:TXmobiles,
		Time:intTime,
		Tpl_id:tplId,
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
	res,err  := client.Post(TXurl, "application/json", b)
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