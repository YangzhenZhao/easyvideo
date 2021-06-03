table! {
    video (id) {
        id -> Integer,
        name -> Text,
        bytes_size -> BigInt,
        video_path -> Text,
        cover_picture_path -> Text,
    }
}
table! {
    tag (id) {
        id -> Integer,
        tag_name -> Text,
    }
}
table! {
    video_tag (id) {
        id -> Integer,
        video_id -> Integer,
        tag_id -> Integer,
    }
}
