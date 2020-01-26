# WechatCronAuto2019-nCov

# 2019新型肺炎疫情 微信定时推送

  使用步骤可参照[python微信天气定时提醒](https://github.com/6yi/WechatAutoWeather/blob/master/README.md)
  
  # 教程

  首先注册公众测试号 ：https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login
  
  记录APPID以及appsecret
  
  <img src='http://59.110.173.180/usr/uploads/2020/01/924286650.jpg'/> 
 
  扫码关注获取openID
  
  <img src='http://59.110.173.180/usr/uploads/2020/01/2046595370.jpg'/> 
  
  创建模板
  
  <img src='http://59.110.173.180/usr/uploads/2020/01/2790200253.jpg'/> 
     
# [命令行版本](https://github.com/6yi/WechatCronAuto2019-nCov/releases)
  按步骤填好appid等直接跑就行,已放出windows可执行exe文件以及linux版本
  
  <img src='http://59.110.173.180/usr/uploads/2020/01/152454484.png'/> 
  
# [json配置版本](https://github.com/6yi/WechatCronAuto2019-nCov/releases)
  在json版本可执行程序目录下创建 2019nCov.json文件,填写格式如下,填完直接运行即可
  ```json
  {
  "spec": "*/10 */1 * * ?",
  "appID": "wxb4b6edasd9",
  "secret":"4ad85249baa128f7c4",
  "templateID": "bzLFye1RtsdEWP4g7ZA_l-QGtzuR5ax4B506-cg",
  "user": [
    {
      "openID": "o3UX4s-YkamCDOsp_sjhU",
      "provinceName": "广东省"
    },
    {
      "openID": "o3UX4s-Yk2asmCDOsp_sjhU",
      "provinceName": "湖北省"
    }
  ]
}
  ```
   
  <img src='https://github.com/6yi/WechatCronAuto2019-nCov/blob/master/src/demo.png'/> 
  
  
  # 数据来源:丁香医生
