from flask_back_end import app, db
from flask import jsonify
from flask_back_end.models import Tag, Video, VideoTag


def all_videos():
    videos = Video.query.all()
    res_list = []
    for video in videos:
        videotags = VideoTag.query.filter(VideoTag.video_id == video.id).all()
        tag_list = []
        for videotag in videotags:
            tag_name = Tag.query.filter(Tag.id == videotag.tag_id).first().tag_name
            tag_list.append(tag_name)
        res_list.append({
            "name": video.name,
            "bytesSize": video.bytes_size,
            "tags": tag_list,
        })
    return jsonify(res_list)
