create database `simple_tiktok`;

use `simple_tiktok`;

-- -------------------------
-- Table structure for user
-- -------------------------
drop table if exists `user`;
create table `user` (
    `id` bigint not null primary key auto_increment,
    `user_id` bigint not null comment '用户ID，雪花算法生成',
    `user_name` varchar(32) not null comment '用户名称，可自定义',
    `password` varchar(64) not null comment '密码，登录使用',
    `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp default null
);

-- -------------------------
-- Table structure for video
-- -------------------------
drop table if exists `video`;
create table `video` (
     `id` bigint not null primary key auto_increment,
     `video_id` bigint not null comment '视频ID，雪花算法生成',
     `video_title` varchar(128) not null comment '视频标题',
     `video_url` varchar(128) not null comment '视频链接',
     `cover_url` varchar(128) default null comment '视频封面链接',
     `user_id` bigint not null comment '视频作者ID，哪个用户发布的视频',
     `like_count` bigint not null default 0 comment '点赞量，默认为 0',
     `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
     `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
     `deleted_at` timestamp default null comment '删除时间',
     index `idx_user_id`(`user_id` ASC) comment '对发布顺序的用户id进行升序排列',
     index `idx_video_title`(`video_title`) comment '以视频的标题来进行排序'
);

-- -------------------------
-- Table structure for like
-- -------------------------
drop table if exists `like`;
create table `like` (
    `id` bigint not null primary key auto_increment,
    `like_id` bigint not null comment '点赞ID，雪花算法生成',
    `user_id` bigint not null comment '用户ID，是谁点了赞',
    `video_id` bigint not null comment '点了哪个视频的赞',
    `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
    `deleted_at` timestamp default null comment '删除时间'
);

-- ----------------------------
-- Table structure for comment
-- ----------------------------
drop table if exists `comment`;
create table `comment` (
        `id` bigint not null primary key auto_increment,
        `comment_id` bigint not null comment '评论ID，雪花算法生成',
        `content` varchar(128) not null comment '评论内容',
        `user_id` bigint not null comment '评论者ID，是谁评论的',
        `video_id` bigint not null comment '评论所属的视频ID',
        `like_count` bigint not null default 0 comment '评论点赞量，默认为 0',
        `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
        `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
        `deleted_at` timestamp default null comment '删除时间'
);

-- --------------------------
-- Table structure for follow
-- --------------------------
drop table if exists `follow`;
create table `follow` (
      `id` bigint not null primary key auto_increment,
      `follow_id` bigint not null comment '关注ID，雪花算法生成',
      `user_id` bigint not null comment '被关注人的ID',
      `follower_id` bigint not null comment '关注者的ID',
      `created_at` timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
      `updated_at` timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
      `deleted_at` timestamp default null comment '删除时间'
);