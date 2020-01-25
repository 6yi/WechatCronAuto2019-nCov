# WechatCronAuto2019-nCov

# 2019新型肺炎疫情 微信定时推送

  具体使用步骤见[python微信天气定时提醒](https://github.com/6yi/WechatAutoWeather/blob/master/README.md),把里面的openID和模板ID填了就完事
  
  
# 命令行版本
  按步骤填好appid等直接跑就行,已放出windows可执行exe文件,linux版本自行编译
  
  ```cmd
  # 交叉编译
  SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build 2019nCov.go
  ```
  
 <img src='https://github.com/6yi/WechatCronAuto2019-nCov/blob/master/src/demo.png'/> 
  
  数据来源:丁香医生
