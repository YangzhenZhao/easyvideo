USERNAME = 'root'
PASSWORD = 'woaini123'
HOST = '127.0.0.1'
PORT = 3306
DATABASE = 'easyvideo'
DB_URI = 'mysql://{}:{}@{}:{}/{}?charset=utf8mb4'.format(
    USERNAME, PASSWORD, HOST, PORT, DATABASE
)
SQLALCHEMY_DATABASE_URI = DB_URI
# 动态追踪修改设置，如未设置，会给出警告信息
SQLALCHEMY_TRACK_MODIFICATIONS = False
