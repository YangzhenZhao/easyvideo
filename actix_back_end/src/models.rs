use serde::{Deserialize, Serialize};

use crate::schema::video;

#[derive(Debug, Clone, Serialize, Deserialize, Queryable, Insertable)]
#[table_name = "video"]
pub struct Video {
    pub id: i32,
    pub name: String,
    pub bytes_size: i64,
    pub video_path: String,
    pub cover_picture_path: String,
}
