
### easyvideo

简易的视频上传下载服务(目前仅用在本地服务器或电脑上)





### 配置   

首先修改后端配置:   
数据库配置    
文件保存路径配置   


### 数据库

Mysql 中新建 easyvideo 数据库，创建的数据表的工作交给后端框架来做


### actix-web

```sh
# ubuntu
sudo apt-get install libsqlite3-dev sqlite3
sudo apt install libpq-dev -y

cargo install diesel_cli
```




### flask   

如果使用 windows，需要在 `__init__.py` 中添加如下内容 

```py
import pymysql
pymysql.install_as_MySQLdb()
```


```bash
$ cd flask_back_end
$ pip install -r requirements.txt

# 创建数据表
$ flask shell
>>> from flask_back_end import db
>>> db.create_all()

$ flask run -h 0.0.0.0
```

在项目根目录执行以下命令:   
由于视频文件可能较大，需要将超时时间设置长一点(这里是 200s)    
具体 workers 数量根据自己情况而定  


```
$ gunicorn --workers=13 flask_back_end:app -b 0.0.0.0 -t 200
```

nginx 配置`/etc/nginx/conf.d/nginx.conf`:   

```
server {
    listen 8888;
    server_name 0.0.0.0;
    location / {
        proxy_pass http://127.0.0.1:8000;
		proxy_redirect off;
    }
}
```

注:     
使用 gunicorn + nginx 在 windows 的 wsl2 中部署后，本地其他机器访问时，下载速度很慢(最慢仅有 100KB/s，最快也只有 2.3MB/s)      


### 前端


```bash
$ cd front_end
$ npm install
$ npm run serve
```

注:    
目前只部署在本地某个机器，暂不使用 nginx 


### 访问

访问`服务器地址:前端启动时的端口`，然后在管理界面设置后端服务器地址
     
