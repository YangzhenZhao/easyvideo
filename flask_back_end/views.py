from flask_back_end import app, db
from flask import send_file, jsonify, request
from flask_back_end.models import Tag, Video
import os


@app.route('/download/<video_name>')
def download(video_name: str):
    video: Video = Video.query.filter(Video.name == video_name).first()
    res = send_file(video.video_path)
    res.mimetype = "application/octet-stream"
    return res


@app.route('/cover_picture/<video_name>')
def cover_picture(video_name: str):
    video: Video = Video.query.filter(Video.name == video_name).first()
    return send_file(video.cover_picture_path)


@app.route('/videos')
def videos():
    videos = Video.query.all()
    res_list = []
    for video in videos:
        tags = Tag.query.filter(Tag.video_id == video.id).all()
        tag_list = [tag.tag_name for tag in tags]
        res_list.append({
            "name": video.name,
            "bytesSize": video.bytes_size,
            "tags": tag_list,
        })
    return jsonify(res_list)


@app.route('/upload/<video_name>', methods=['POST'])
def upload(video_name: str):
    f = request.files['file']
    save_dir = os.path.join(app.config["STORAGE_DIR"], video_name)
    if not os.path.exists(save_dir):
        os.makedirs(save_dir)
    path = os.path.join(app.config["STORAGE_DIR"], save_dir, f.filename)
    f.save(path)
    return ''


@app.route('/save_video', methods=['POST'])
def svae_video():
    args = request.args
    video = Video()
    video.name = args["name"]
    video.bytes_size = int(args["bytes_size"])
    storage_dir = app.config["STORAGE_DIR"]
    video.video_path = os.path.join(storage_dir, video.name, args["video_file_name"])
    video.cover_picture_path = os.path.join(storage_dir, video.name, args["cover_picture_file_name"])
    db.session.add(video)
    db.session.commit()
    return ''
