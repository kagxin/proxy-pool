create table proxy
(
    `id`         int          unsigned auto_increment primary key comment '自增长ID',
    `schema`     varchar(8)   default 0 not null comment 'http, https',
    `ip`         varchar(64)  default '' not null comment '代理IP',
    `port`       int          unsigned default 0 not null comment '代理端口',
    `from`       tinyint      default 0 not null comment '0-未定义，1-西刺, 2-全网',
    `is_deleted` tinyint      default 0                 not null comment '0-正常，1-删除',
    `ctime`      timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
    `mtime`      timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '最后修改时间'
)  engine = InnoDB
    auto_increment = 1
    default charset = utf8mb4
    comment '代理信息';