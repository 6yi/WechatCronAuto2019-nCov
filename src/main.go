package main

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func main(){
	spec := "0 0 0,9,15,20 * * ?"
	c := cron.New()
	c.AddFunc(spec, func() {
		sendMsg("o3UX4s-Yk2vJNG9zxOmCDOsp_sjhU","bzLFye1Rt5zo2dmEWzuR5ax4B506-cg","广东省")
	})
	c.Start()
	select{}
}


func sendMsg(openID string,templateID string,provinceName string){
	pmsg:=getMsg(provinceName)
	var msg string
	if pmsg.state==0{
		msg="数据出错"
	}else{
		msg=getSendMsg(pmsg)
	}
	fmt.Println(msg)
	api := getAPI("wx89b4bd5cab8789", "4ad85249b589ead128f7c4")
	u:="https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+api.AccessToken
	ms:=&MS{
		Value:msg,
		Color:"#173177",
	}
	body:=&Body{
		Touser:openID,
		TemplateID:templateID,
		URL:"https://3g.dxy.cn/newh5/view/pneumonia?scene=2&clicktime=1579579384&enterid=1579579384&from=groupmessage&isappinstalled=0",
		Data: struct{ MS }{MS: *ms},
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b,err:=json.Marshal(body)
	client := http.Client{}
	resp, err := client.Post(u, "application/json", bytes.NewBuffer(b))
	if err!=nil {

	}
	defer resp.Body.Close()

	bo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(bo))
}

type Body struct {
	Touser string `json:"touser"`
	TemplateID string `json:"template_id"`
	URL string `json:"url"`
	Data struct {
		MS
	} `json:"data"`
}
type MS struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

func getSendMsg(msg *reMsg) string {
	var str string
	str=msg.ProvinceName+
		"\t确诊:"+strconv.Itoa(msg.ConfirmedCount)+
		"死亡:"+strconv.Itoa(msg.DeadCount)+
		"治愈:"+strconv.Itoa(msg.CuredCount)+
		"\n\n"

	for _,city:=range msg.Cities{
		str=str+city.CityName+"\t确诊:"+
			strconv.Itoa(city.ConfirmedCount)+
			"\t死亡:"+strconv.Itoa(city.DeadCount)+
			"\t治愈:"+strconv.Itoa(city.CuredCount)+"\n"
	}
	return str
}


func getAPI(appid string, secret string) *token {
	url:="https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&&"+"appid="+appid+"&&secret="+secret
	resp,erro:=http.Get(url)
	if erro!=nil{
		fmt.Println("erro")
	}
	body,erro:=ioutil.ReadAll(resp.Body)
	if erro!=nil {
		fmt.Println("erro")
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	token:=new(token)
	json.Unmarshal(body,token)
	return token
}

func getMsg(provinceName string) *reMsg{
	resp,erro:=http.Get("https://3g.dxy.cn/newh5/view/pneumonia?scene=2&clicktime=1579579384&enterid=1579579384&from=groupmessage&isappinstalled=0")
	if erro!=nil{
		return &reMsg{state:0}
	}
	body,erro:=ioutil.ReadAll(resp.Body)
	if erro!=nil {
		return &reMsg{state:0}
	}
	strs:=(regexp.MustCompile("\\{\"provinceName\":\""+provinceName+"\".*?\\}\\]\\}").FindAll(body,-1))[0]
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	data:=new(Msg)
	json.Unmarshal(strs,data)
	return &reMsg{state:1,Msg:*data}
}


type reMsg struct {
	state int
	Msg
}

type Msg struct {
	ProvinceName string `json:"provinceName"`
	ProvinceShortName string `json:"provinceShortName"`
	ConfirmedCount int `json:"confirmedCount"`
	SuspectedCount int `json:"suspectedCount"`
	CuredCount int `json:"curedCount"`
	DeadCount int `json:"deadCount"`
	Comment string `json:"comment"`
	Cities []struct {
		CityName string `json:"cityName"`
		ConfirmedCount int `json:"confirmedCount"`
		SuspectedCount int `json:"suspectedCount"`
		CuredCount int `json:"curedCount"`
		DeadCount int `json:"deadCount"`
	} `json:"cities"`
}


type token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
}
