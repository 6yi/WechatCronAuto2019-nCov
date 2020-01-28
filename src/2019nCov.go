package main

import (
	"bufio"
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

type users struct {
	appid string
	secret string
	templateID string
	openid []string
	provinceName []string
}


func main(){
	gui()
}

func gui(){
	var appid string
	var secret string
	var number int
	fmt.Println("\n2019新型肺炎疫情微信公众号定时推送\n")
Here:
	fmt.Println("请设定appID和secret:\n")
	fmt.Printf("appID:")
	fmt.Scanln(&appid)
	fmt.Printf("secret:")
	fmt.Scanln(&secret)
	token:=getAPI(appid,secret)
	if token.AccessToken=="" {
		fmt.Printf("appid或者secret设置错误")
		goto Here
	}
	var templateID string
	fmt.Println("请设定推送模板templateID:")
	fmt.Scanln(&templateID)
	us:=new(users)
	us.templateID=templateID
	us.secret=secret
	us.appid=appid
	fmt.Println("设定推送人数: ")
	fmt.Scanln(&number)
	var openid string
	var provinceName string
	us.openid=make([]string,number)
	us.provinceName=make([]string,number)
	for key := 0; key < number; key++ {
		fmt.Println("第",key+1,"位用户的openID:")
		fmt.Scanln(&openid)
		us.openid[key]=openid
		fmt.Println("第",key+1,"位用户需要推送的省份,请写全称如(广东省):")
		fmt.Scanln(&provinceName)
		us.provinceName[key]=provinceName
	}
	def_spec := "0 0 0,9,15,20 * ?"
	var spec string

	fmt.Printf("定点推送时间:(请按照linuxCron语法   按回车默认  0 0 0,9,15,20 * ? 每日0时 9时 15时 20时推送)\n")

	Scanf(&spec)
	if spec=="" {
		spec=def_spec
	}
	fmt.Println("推送开始")
	c := cron.New()
	c.AddFunc(spec, func() {
		for key := 0; key < number; key++ {
			sendMsg(us.appid,us.secret,us.openid[key],us.templateID,us.provinceName[key])
		}
	})
	c.Start()
	select{}

}

func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
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
	strs2:=(regexp.MustCompile("<span style=\"color: rgb\\(65, 105, 226\\);\">.*?</span>").FindAll(body,-1))
	all:=make([]string,4)
	for key,y:=range strs2{
		x:=strings.Split(string(y),"<span style=\"color: rgb(65, 105, 226);\">")[1]
		z:=strings.Split(x,"</span>")[0]
		all[key]=z
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	data:=new(Msg)
	json.Unmarshal(strs,data)
	data.AllConfirmedCount=all[0]
	data.AllSuspectedCount=all[1]
	data.AllDeadCount=all[2]
	data.AllCuredCount=all[3]
	return &reMsg{state:1,Msg:*data}
}

func getSendMsg(msg *reMsg) string {
	var str string
	str="全国"+"\t确诊:"+msg.AllConfirmedCount+"\t死亡:"+msg.AllDeadCount+"\t疑似:"+msg.AllSuspectedCount+"\t治愈:"+msg.AllCuredCount+"\n"+msg.ProvinceName+
		"\t确诊:"+strconv.Itoa(msg.ConfirmedCount)+
		"\t死亡:"+strconv.Itoa(msg.DeadCount)+
		"\t疑似:"+strconv.Itoa(msg.SuspectedCount)+
		"\t治愈:"+strconv.Itoa(msg.CuredCount)+
		"\n\n"

	for _,city:=range msg.Cities{
		str=str+city.CityName+"\t确诊:"+
			strconv.Itoa(city.ConfirmedCount)+
			"\t死亡:"+strconv.Itoa(city.DeadCount)+
			"\t疑似:"+strconv.Itoa(city.SuspectedCount)+
			"\t治愈:"+strconv.Itoa(city.CuredCount)+"\n"
	}
	return str
}


type reMsg struct {
	state int
	Msg
}

type Msg struct {
	AllConfirmedCount string
	AllDeadCount string
	AllCuredCount string
	AllSuspectedCount string
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