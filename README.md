# WechatCronAuto2019-nCov

# 2019新型肺炎疫情 微信定时推送

  具体使用步骤见[python微信天气定时提醒](https://github.com/6yi/WechatAutoWeather/blob/master/README.md),把里面的openID和模板ID填了就完事
  
  
# [命令行版本](https://github.com/6yi/WechatCronAuto2019-nCov/releases)
  按步骤填好appid等直接跑就行,已放出windows可执行exe文件,linux版本自行编译
  
  <img src='http://59.110.173.180/usr/uploads/2020/01/152454484.png'/> 
 
  ```cmd
  # 交叉编译
  SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build 2019nCov.go
  ```
  
 <img src='https://github.com/6yi/WechatCronAuto2019-nCov/blob/master/src/demo.png'/> 
  
  
  
  
  # 教程

  首先注册公众测试号 ：https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login
  记录APPID以及appsecret
  
  <img src='http://59.110.173.180/usr/uploads/2020/01/924286650.jpg'/> 
 
  扫码关注获取openID
  
  <img src='http://59.110.173.180/usr/uploads/2020/01/2046595370.jpg'/> 
  
  创建模板
  
  <img src='http://59.110.173.180/usr/uploads/2020/01/2790200253.jpg'/> 
  
  建议下载gui版本或者json版本,把相对应的数据填上去就OK了
  
  数据来源:丁香医生
