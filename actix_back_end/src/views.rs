use crate::models;
use crate::schema;
use crate::schema::tag::dsl::*;
use crate::schema::video::dsl::*;
use crate::schema::video_tag::dsl::*;
use actix_files::NamedFile;
use actix_web::{get, web, Error, HttpResponse, Result};
use diesel::prelude::*;
use diesel::r2d2::{self, ConnectionManager};
use serde::{Deserialize, Serialize};

type DbPool = r2d2::Pool<ConnectionManager<MysqlConnection>>;

#[derive(Debug, Clone, Serialize, Deserialize)]
struct VideoViewItem {
    pub name: String,
    pub bytesSize: i64,
    pub tags: Vec<String>,
}

fn get_tag_name(conn: &MysqlConnection, tag_id_t: i32) -> String {
    tag.filter(schema::tag::columns::id.eq(tag_id_t))
        .first::<models::Tag>(conn)
        .unwrap()
        .tag_name
}

#[get("/videos")]
pub async fn videos(pool: web::Data<DbPool>) -> Result<HttpResponse, Error> {
    let conn = pool.get().expect("couldn't get db connection from pool");

    let videos = video.load::<models::Video>(&conn).unwrap();
    let mut res_list: Vec<VideoViewItem> = vec![];
    for video_item in videos.into_iter() {
        let video_tags = video_tag
            .filter(video_id.eq(video_item.id))
            .load::<models::VideoTag>(&conn)
            .unwrap();
        let mut tag_list = vec![];
        for video_tag_item in video_tags.iter() {
            tag_list.push(get_tag_name(&conn, video_tag_item.tag_id));
        }
        res_list.push(VideoViewItem {
            name: video_item.name,
            bytesSize: video_item.bytes_size,
            tags: tag_list,
        });
    }

    Ok(HttpResponse::Ok().json(res_list))
}

#[get("/cover_picture/{video_name}")]
pub async fn cover_picture(
    pool: web::Data<DbPool>,
    video_name: web::Path<String>,
) -> Result<NamedFile> {
    let conn = pool.get().expect("couldn't get db connection from pool");

    let target_video = video
        .filter(name.eq(format!("{}", video_name)))
        .first::<models::Video>(&conn)
        .unwrap();

    Ok(NamedFile::open(target_video.cover_picture_path)?)
}

#[get("/download/{video_name}")]
pub async fn download(pool: web::Data<DbPool>, video_name: web::Path<String>) -> Result<NamedFile> {
    let conn = pool.get().expect("couldn't get db connection from pool");
    let target_video = video
        .filter(name.eq(format!("{}", video_name)))
        .first::<models::Video>(&conn)
        .unwrap();
    let named_file = NamedFile::open(target_video.video_path).unwrap();
    let mime_type = "application/octet-stream".parse::<mime::Mime>().unwrap();
    let named_file = named_file.set_content_type(mime_type);
    Ok(named_file)
}

#[get("/tags")]
pub async fn tags(pool: web::Data<DbPool>) -> Result<HttpResponse, Error> {
    let conn = pool.get().expect("couldn't get db connection from pool");

    let mut res_list: Vec<String> = vec![];
    let tags = tag.load::<models::Tag>(&conn).unwrap();
    for tag_item in tags.into_iter() {
        res_list.push(tag_item.tag_name);
    }

    Ok(HttpResponse::Ok().json(res_list))
}
