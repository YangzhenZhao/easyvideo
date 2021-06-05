use actix_cors::Cors;
use actix_web::{middleware, App, HttpServer};
use diesel::prelude::*;
use diesel::r2d2::{self, ConnectionManager};

#[macro_use]
extern crate diesel;

mod models;
mod schema;
mod views;

use crate::views::{cover_picture, download, svae_video, tags, tags_videos, upload, videos};

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init();
    let manager = ConnectionManager::<MysqlConnection>::new(
        "mysql://root:123456@127.0.0.1:3306/easyvideo?charset=utf8mb4",
    );
    let pool = r2d2::Pool::builder()
        .build(manager)
        .expect("Failed to create pool.");

    let bind = "0.0.0.0:8000";

    println!("Starting server at: {}", &bind);

    HttpServer::new(move || {
        App::new()
            .data(pool.clone())
            .wrap(
                Cors::default()
                    .allow_any_origin()
                    .allow_any_method()
                    .allow_any_header(),
            )
            .wrap(middleware::Logger::default())
            .service(videos)
            .service(download)
            .service(cover_picture)
            .service(tags)
            .service(upload)
            .service(svae_video)
            .service(tags_videos)
    })
    .bind(&bind)?
    .run()
    .await
}
