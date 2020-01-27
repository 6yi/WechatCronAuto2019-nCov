package main

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/robfig/cron"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)
func main(){
	buf,err:=ioutil.ReadFile("2019nCov.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "配置文件不存在!")
		return
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	data:=new(ini)
	json.Unmarshal(buf,data)
	c:=cron.New()
	fmt.Println("开始推送")
	c.AddFunc(data.Spec, func() {
		for _,user:=range data.User{
			sendMsg(data.AppID,data.Secret,user.OpenID,data.TemplateID,user.ProvinceName)
		}
	})
	c.Start()
	select {

	}
}


func sendMsg(appid string,secret string,openID string,templateID string,provinceName string){
	pmsg:=getMsg(provinceName)
	var msg string
	if pmsg.state==0{
		msg="数据出错"
	}else{
		msg=getSendMsg(pmsg)
	}
	fmt.Println(msg)
	api := getAPI(appid, secret)
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
	str="全国"+"\t确诊:"+msg.AllConfirmedCount+"\t死亡:"+msg.AllDeadCount+"\t治愈:"+msg.AllCuredCount+"\n"+msg.ProvinceName+
		"\t确诊:"+strconv.Itoa(msg.ConfirmedCount)+
		"\t死亡:"+strconv.Itoa(msg.DeadCount)+
		"\t治愈:"+strconv.Itoa(msg.CuredCount)+
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
	strs2:=(regexp.MustCompile("<span style=\"color: #4169e2\">.*?</span>").FindAll(body,-1))
	all:=make([]string,4)
	for key,y:=range strs2{
		x:=strings.Split(string(y),"<span style=\"color: #4169e2\">")[1]
		z:=strings.Split(x,"</span>")[0]
		all[key]=z
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	data:=new(Msg)
	json.Unmarshal(strs,data)
	data.AllConfirmedCount=all[0]
	data.AllDeadCount=all[2]
	data.AllCuredCount=all[3]
	return &reMsg{state:1,Msg:*data}
}


type reMsg struct {
	state int
	Msg
}

type Msg struct {
	AllConfirmedCount string
	AllDeadCount string
	AllCuredCount string
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

type ini struct {
	Spec string `json:"spec"`
	AppID string `json:"appID"`
	Secret string `json:"secret"`
	TemplateID string `json:"templateID"`
	User []struct {
		OpenID string `json:"openID"`
		ProvinceName string `json:"provinceName"`
	} `json:"user"`
}

type token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
}

