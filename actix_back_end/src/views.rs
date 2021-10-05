use crate::models;
use crate::schema;
use crate::schema::tag::dsl::*;
use crate::schema::video::dsl::*;
use crate::schema::video_tag::dsl::*;
use actix_files::NamedFile;
use actix_multipart::Multipart;
use actix_web::{get, post, web, Error, HttpResponse, Result};
use diesel::insert_into;
use diesel::prelude::*;
use diesel::r2d2::{self, ConnectionManager};
use futures::{StreamExt, TryStreamExt};
use serde::{Deserialize, Serialize};
use std::collections::HashSet;
use std::io::Write;

type DbPool = r2d2::Pool<ConnectionManager<MysqlConnection>>;
const STORAGE_DIR: &str = "/mnt/d/video_storage/";

#[derive(Debug, Clone, Serialize, Deserialize)]
struct VideoViewItem {
    pub name: String,
    #[serde(rename(serialize = "bytesSize"))]
    pub bytes_size: i64,
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
            bytes_size: video_item.bytes_size,
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
pub async fn upload(
    _pool: web::Data<DbPool>,
    mut payload: Multipart,
    video_name: web::Path<String>,
) -> Result<HttpResponse, Error> {
    while let Ok(Some(mut field)) = payload.try_next().await {
        let content_type = field.content_disposition().unwrap();
        let filename = content_type.get_filename().unwrap();
        let filepath = format!("{}{}", STORAGE_DIR, video_name);

        let filepath_tmp = filepath.clone();
        web::block(|| std::fs::create_dir(filepath_tmp)).await.ok();

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
pub struct SaveFileParams {
    pub name: String,
    pub bytes_size: i64,
    pub video_file_name: String,
    pub cover_picture_file_name: String,
    pub tags: String,
}

#[post("/save_video")]
pub async fn svae_video(
    pool: web::Data<DbPool>,
    info: web::Query<SaveFileParams>,
) -> Result<HttpResponse, Error> {
    let conn = pool.get().expect("couldn't get db connection from pool");

    let new_video = models::NewVideo {
        name: &info.name,
        bytes_size: info.bytes_size,
        video_path: &format!("{}{}/{}", STORAGE_DIR, info.name, info.video_file_name),
        cover_picture_path: &format!(
            "{}{}/{}",
            STORAGE_DIR, info.name, info.cover_picture_file_name
        ),
    };
    insert_into(video).values(&new_video).execute(&conn).ok();
    let tag_names: Vec<String> = serde_json::from_str(&info.tags).unwrap();
    let mut tag_ids: Vec<i32> = vec![];
    for tag_name_tmp in tag_names.into_iter() {
        let tag_obj = tag
            .filter(tag_name.eq(tag_name_tmp.clone()))
            .first::<models::Tag>(&conn);
        match tag_obj {
            Ok(obj) => tag_ids.push(obj.id),
            Err(_) => {
                let new_tag = models::NewTag {
                    tag_name: &tag_name_tmp,
                };
                insert_into(tag).values(&new_tag).execute(&conn).ok();
                let tag_id_tmp = tag
                    .filter(tag_name.eq(tag_name_tmp))
                    .first::<models::Tag>(&conn)
                    .unwrap()
                    .id;
                tag_ids.push(tag_id_tmp);
            }
        }
    }
    let video_id_tmp = video
        .filter(name.eq(new_video.name))
        .first::<models::Video>(&conn)
        .unwrap()
        .id;
    for tag_id_tmp in tag_ids.into_iter() {
        let new_video_tag = models::NewVideoTag {
            video_id: video_id_tmp,
            tag_id: tag_id_tmp,
        };
        insert_into(video_tag)
            .values(&new_video_tag)
            .execute(&conn)
            .ok();
    }

    Ok(HttpResponse::Ok().into())
}

fn get_intersection(id_set: HashSet<i32>, video_tag_list: Vec<models::VideoTag>) -> HashSet<i32> {
    let mut res_set = HashSet::new();
    for video_tag_tmp in video_tag_list.into_iter() {
        if id_set.contains(&video_tag_tmp.video_id) {
            res_set.insert(video_tag_tmp.video_id);
        }
    }
    res_set
}

#[get("/tags_videos/{tags}")]
pub async fn tags_videos(
    pool: web::Data<DbPool>,
    tag_names: web::Path<String>,
) -> Result<HttpResponse, Error> {
    let conn = pool.get().expect("couldn't get db connection from pool");
    let tag_list: Vec<&str> = tag_names.split(',').collect();

    let mut video_id_set: HashSet<i32> = HashSet::new();
    for (i, tag_name_tmp) in tag_list.into_iter().enumerate() {
        let tag_id_tmp = tag
            .filter(tag_name.eq(tag_name_tmp))
            .first::<models::Tag>(&conn)
            .unwrap()
            .id;
        let video_tags_tmp = video_tag
            .filter(tag_id.eq(tag_id_tmp))
            .load::<models::VideoTag>(&conn)
            .unwrap();
        if i == 0 {
            for video_tag_tmp in video_tags_tmp.into_iter() {
                video_id_set.insert(video_tag_tmp.video_id);
            }
        } else {
            video_id_set = get_intersection(video_id_set, video_tags_tmp);
        }
    }

    let mut res_list: Vec<VideoViewItem> = vec![];
    for video_id_tmp in video_id_set.into_iter() {
        let video_tmp = video
            .filter(schema::video::columns::id.eq(video_id_tmp))
            .first::<models::Video>(&conn)
            .unwrap();
        let video_tag_list = video_tag
            .filter(video_id.eq(video_tmp.id))
            .load::<models::VideoTag>(&conn)
            .unwrap();
        let mut tag_list: Vec<String> = vec![];
        for video_tag_tmp in video_tag_list.into_iter() {
            let tag_name_tmp = tag
                .filter(schema::tag::columns::id.eq(video_tag_tmp.tag_id))
                .first::<models::Tag>(&conn)
                .unwrap()
                .tag_name;
            tag_list.push(tag_name_tmp);
        }
        res_list.push(VideoViewItem {
            name: video_tmp.name,
            bytes_size: video_tmp.bytes_size,
            tags: tag_list,
        });
    }

    Ok(HttpResponse::Ok().json(res_list))
}
