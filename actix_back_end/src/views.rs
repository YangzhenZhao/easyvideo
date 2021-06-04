use crate::models;
use crate::schema;
use crate::schema::tag::dsl::*;
use crate::schema::video::dsl::*;
use crate::schema::video_tag::dsl::*;
use actix_files::NamedFile;
use actix_web::{get, web, post, Error, HttpResponse, Result};
use diesel::prelude::*;
use diesel::r2d2::{self, ConnectionManager};
use serde::{Deserialize, Serialize};
use actix_multipart::Multipart;
use futures::{StreamExt, TryStreamExt};
use std::io::Write;
use diesel::insert_into;

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


#[post("/upload/{video_name}")]
pub async fn upload(pool: web::Data<DbPool>, mut payload: Multipart, video_name: web::Path<String>) -> Result<HttpResponse, Error> {
    while let Ok(Some(mut field)) = payload.try_next().await {
        let content_type = field.content_disposition().unwrap();
        let filename = content_type.get_filename().unwrap();
        let filepath = format!("/mnt/f/easyvideo/storage/{}", video_name);

        let filepath_tmp = filepath.clone();
        web::block(|| std::fs::create_dir(filepath_tmp))
            .await
            .ok();
        

        let filepath = format!("{}/{}", filepath, sanitize_filename::sanitize(&filename));

        let mut f = web::block(|| std::fs::File::create(filepath))
            .await
            .unwrap();

        while let Some(chunk) = field.next().await {
            let data = chunk.unwrap();
            f = web::block(move || f.write_all(&data).map(|_| f)).await?;
        }
    }
    Ok(HttpResponse::Ok().into())
}



#[derive(Debug, Clone, Serialize, Deserialize)]
struct SaveFileParams {
    pub name: String,
    pub bytes_size: i64,
    pub video_file_name: String,
    pub cover_picture_file_name: String,
    pub tags: String,
}

#[post("/save_video")]
pub async fn svae_video(pool: web::Data<DbPool>, info: web::Query<SaveFileParams>) -> Result<HttpResponse, Error> {
    let conn = pool.get().expect("couldn't get db connection from pool");

    println!("SaveFileParams = {:?}", info);
    let new_video = models::NewVideo {
        name: &info.name,
        bytes_size: info.bytes_size,
        video_path: &format!("/mnt/f/easyvideo/storage/{}/{}", info.name, info.video_file_name),
        cover_picture_path: &format!("/mnt/f/easyvideo/storage/{}/{}", info.name, info.cover_picture_file_name),
    };
    insert_into(video).values(&new_video).execute(&conn);
    let tag_names: Vec<String> = serde_json::from_str(&info.tags).unwrap();
    println!("tag_names = {:?}", tag_names);
    let mut tag_ids: Vec<i32> = vec![];
    for tag_anme in tag_names.into_iter() {

    }
    let mut res_list: Vec<String> = vec![];
    // let tags = tag.load::<models::Tag>(&conn).unwrap();
    // for tag_item in tags.into_iter() {
    //     res_list.push(tag_item.tag_name);
    // }

    Ok(HttpResponse::Ok().json(res_list))
}
