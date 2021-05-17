from flask import Flask, send_file
from flask_sqlalchemy import SQLAlchemy
from flask_cors import CORS

app = Flask("flask_back_end")
app.config.from_pyfile('settings.py')
db = SQLAlchemy(app)
CORS(app, resources=r'/*')


from flask_back_end import views