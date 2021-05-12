from flask import Flask, send_file
from flask_sqlalchemy import SQLAlchemy

app = Flask("flask_back_end")
app.config.from_pyfile('settings.py')
db = SQLAlchemy(app)


from flask_back_end import views