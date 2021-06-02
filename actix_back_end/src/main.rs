use actix_web::{middleware, App, HttpServer};
use diesel::prelude::*;
use diesel::r2d2::{self, ConnectionManager};

#[macro_use]
extern crate diesel;

mod models;
mod schema;
mod views;

use crate::views::videos;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    env_logger::init();
    let manager = ConnectionManager::<MysqlConnection>::new(
        "mysql://root:123456@127.0.0.1:3306/easyvideo?charset=utf8mb4",
    );
    let pool = r2d2::Pool::builder()
        .build(manager)
        .expect("Failed to create pool.");

    let bind = "0.0.0.0:8080";

    println!("Starting server at: {}", &bind);

    HttpServer::new(move || {
        App::new()
            .data(pool.clone())
            .wrap(middleware::Logger::default())
            .service(videos)
    })
    .bind(&bind)?
    .run()
    .await
}
