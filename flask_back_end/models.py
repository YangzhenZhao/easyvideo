from flask_back_end import db


class Video(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(50), unique=True, nullable=False)
    size = db.Column(db.Float(2), nullable=False)
    video_path = db.Column(db.String(100), unique=True, nullable=False)
    cover_picture_path = db.Column(db.String(100))


class Tag(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    video_id = db.Column(db.Integer, nullable=False)
    tag_name = db.Column(db.String(50), nullable=False)
