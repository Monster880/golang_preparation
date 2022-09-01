
create table comment
(
    activity_id varchar(32),
    comment_id varchar(32),
    user_id varchar(32),
    text varchar(128),
    create_time date,
    update_time date
),

create table approve(
    activity_id varchar(32),
    approve_id varchar(32),
    user_id varchar(32),
    create_time date,
    update_time date
)

create table user(
     user_id varchar(32),
     user_name varchar(32),
     user_image varchar(32),
     create_time date,
     update_time date,
)

create table receive(
    activity_id varchar(32),
    user_id varchar(32),
    type integer,
    type_id varchar(32),
    create_time date,
    update_time date,
)

create table activity(
    activity_id varchar(32),
    user_id varchar(32),
    text varchar(128),
    user_image varchar(32),
    comment_num integer,
    approve_num integer,
    create_time date,
    update_time date,
)
