# blog学习

1. [学习连接](https://eddycjy.com/posts/go/gin/2018-02-11-api-01/)


## 项目开始 思考

1. 配置文件处理
2. db句柄处理
3. API错误处理
4. 分页等公共逻辑处理
5. 路由不要放在main， 放在专门的routers 文件


****

项目相关工具

1. 读取配置[go-ini/ini](https://github.com/go-ini/ini)


项目相关库
1. github.com/unknwon/com  工具库
2. github.com/go-ini/ini    配置库
3. log  日志库
4. github.com/jinzhu/gorm [gorm 官网文档](https://gorm.io/docs/)
5. github.com/astaxie/beego/validation 表单验证

**** 

## 定义接口
1. 获取标签列表 get （"/tags"）
2. 新建标签列表 post（"/tags"）
3. 更新指定标签列表 put（"/tags/:id"）
4. 删除指定标签列表 delete（"/tags/:id"）