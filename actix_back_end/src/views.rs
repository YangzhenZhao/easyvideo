use crate::models;
use actix_web::{get, web, Error, HttpResponse};
use diesel::prelude::*;
use diesel::r2d2::{self, ConnectionManager};

type DbPool = r2d2::Pool<ConnectionManager<MysqlConnection>>;

fn all_videos(conn: &MysqlConnection) -> Result<Vec<models::Video>, diesel::result::Error> {
    use crate::schema::video::dsl::*;

    video.load(conn)
}

#[get("/videos")]
pub async fn videos(pool: web::Data<DbPool>) -> Result<HttpResponse, Error> {
    let conn = pool.get().expect("couldn't get db connection from pool");

    let videos = web::block(move || all_videos(&conn)).await.map_err(|e| {
        eprintln!("{}", e);
        HttpResponse::InternalServerError().finish()
    })?;

    Ok(HttpResponse::Ok().json(videos))
}
