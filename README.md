
### easyvideo

简易的视频上传下载服务(目前仅用在本地服务器或电脑上)



### 启动服务

由于只在本地使用，暂时偷懒不用 Nginx 了(


#### 数据库

修改后端配置中的数据库用户名和密码

Mysql 中新建 easyvideo 数据库，创建的数据表的工作交给后端框架来做

#### 前端


```bash
$ cd front_end
$ npm run serve
```


#### flask   

如果使用 windows，需要在 `__init__.py` 中添加如下内容 

```py
import pymysql
pymysql.install_as_MySQLdb()
```


```bash
$ cd flask_back_end

# 创建数据表
$ flask shell
>>> from flask_back_end import db
>>> db.create_all()

$ flask run -h 0.0.0.0
```


### 访问

访问`服务器地址:前端启动时的端口`，然后在管理界面设置后端服务器地址
     
