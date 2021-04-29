drop database  if exists `tx_tcc_user_info`;

CREATE DATABASE `tx_tcc_user_info` charset utf8;

use `tx_tcc_user_info`;


CREATE TABLE `user_info` (
 `id` int(11) NOT NULL COMMENT 'primary key',
 `username` varchar(255) DEFAULT NULL COMMENT '用户名',
 `phone_number` bigint(32) DEFAULT NULL COMMENT '手机号',
 `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
 `photo` varchar(255) DEFAULT NULL COMMENT '照片',
 `longitude` float(64,0) DEFAULT NULL COMMENT '经度',
  `latitude` float(64,0) DEFAULT NULL COMMENT '维度',
  `current_address` varchar(255) DEFAULT NULL COMMENT '当前地点',
  `hobby_tags` varchar(255) DEFAULT NULL COMMENT '兴趣标签',
  `consume_tags` varchar(255) DEFAULT NULL COMMENT '消费倾向',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `seata_tcc_storage`.`userInfo` (`id`, `product_id`, `total`, `used`, `residue`) VALUES ('1', '1', '100000', '0', '100000');

-- for AT mode you must to init this sql for you business database. the seata server not need it.
CREATE TABLE IF NOT EXISTS `undo_log`
(
    `branch_id`     BIGINT(20)   NOT NULL COMMENT 'branch transaction id',
    `xid`           VARCHAR(100) NOT NULL COMMENT 'global transaction id',
    `context`       VARCHAR(128) NOT NULL COMMENT 'undo_log context,such as serialization',
    `rollback_info` LONGBLOB     NOT NULL COMMENT 'rollback info',
    `log_status`    INT(11)      NOT NULL COMMENT '0:normal status,1:defense status',
    `log_created`   DATETIME(6)  NOT NULL COMMENT 'create datetime',
    `log_modified`  DATETIME(6)  NOT NULL COMMENT 'modify datetime',
    UNIQUE KEY `ux_undo_log` (`xid`, `branch_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8 COMMENT ='AT transaction mode undo table';

CREATE TABLE IF NOT EXISTS segment
(
    id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '自增主键',
    VERSION       BIGINT      DEFAULT 0  NOT NULL COMMENT '版本号',
    business_type VARCHAR(63) DEFAULT '' NOT NULL COMMENT '业务类型，唯一',
    max_id        BIGINT      DEFAULT 0  NOT NULL COMMENT '当前最大id',
    step          INT         DEFAULT 0  NULL COMMENT '步长',
    increment     INT         DEFAULT 1  NOT NULL COMMENT '每次id增量',
    remainder     INT         DEFAULT 0  NOT NULL COMMENT '余数',
    created_at    BIGINT UNSIGNED        NOT NULL COMMENT '创建时间',
    updated_at    BIGINT UNSIGNED        NOT NULL COMMENT '更新时间',
    CONSTRAINT uniq_business_type UNIQUE (business_type)
    ) CHARSET = utf8mb4
    ENGINE INNODB COMMENT '号段表';


INSERT INTO segment
(VERSION, business_type, max_id, step, increment, remainder, created_at, updated_at)
VALUES (1, 'user_info_id', 1000, 1000, 1, 0, NOW(), NOW());
