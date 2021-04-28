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
