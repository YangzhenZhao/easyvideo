use serde::{Deserialize, Serialize};

use crate::schema::{tag, video, video_tag};

#[derive(Debug, Clone, Serialize, Deserialize, Queryable, Insertable)]
#[table_name = "video"]
pub struct Video {
    pub id: i32,
    pub name: String,
    pub bytes_size: i64,
    pub video_path: String,
    pub cover_picture_path: String,
}

#[derive(Debug, Clone, Serialize, Deserialize, Queryable, Insertable)]
#[table_name = "tag"]
pub struct Tag {
    pub id: i32,
    pub tag_name: String,
}

#[derive(Debug, Clone, Serialize, Deserialize, Queryable, Insertable)]
#[table_name = "video_tag"]
pub struct VideoTag {
    pub id: i32,
    pub video_id: i32,
    pub tag_id: i32,
}

#[derive(Insertable)]
#[table_name = "video"]
pub struct NewVideo<'a> {
    pub name: &'a str,
    pub bytes_size: i64,
    pub video_path: &'a str,
    pub cover_picture_path: &'a str,
}
