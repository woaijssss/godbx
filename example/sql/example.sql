create table user_info
(
    id        bigint(20) not null AUTO_INCREMENT primary key,
    `name`    varchar(200) not null comment 'user name',
    create_at datetime     not null,
    modify_at datetime null
) ENGINE=innodb CHARACTER SET utf8mb4 comment 'axxx';