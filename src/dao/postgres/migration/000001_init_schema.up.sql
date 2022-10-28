create type relationType As ENUM ('like','matched','dislike');
create type userType as enum ('user','manager');
-- 用户
create table "user"
(
    id   bigserial primary key,        -- 用户id
    name varchar(255) not null unique, --用户姓名
    user_type userType                      --用户类型
);
-- 用户关系表
create table "relationship"
(
    fid  bigserial not null references "user" (id) on delete cascade on update cascade,
    tid  bigserial not null references "user" (id) on delete cascade on update cascade,
    relation_type relationType
)
