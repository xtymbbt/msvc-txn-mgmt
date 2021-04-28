drop database  if exists `tx_tcc_profile`;

CREATE DATABASE `tx_tcc_profile` charset utf8;

use `tx_tcc_profile`;

CREATE TABLE `profile` (
    `id` int(11) NOT NULL COMMENT 'primary key',
    `username` varchar(255) DEFAULT NULL COMMENT '现名',
    `age` int(8) DEFAULT NULL COMMENT '年龄',
    `gender` int(8) DEFAULT NULL COMMENT '性别',
    `politic_type` varchar(255) DEFAULT NULL COMMENT '群众、共青团员、中共党员、民主人士等',
    `birth_date` datetime DEFAULT NULL COMMENT '出生年月日',
    `photo` varchar(255) DEFAULT NULL COMMENT '照片',
    `old_name` varchar(255) DEFAULT NULL COMMENT '曾用名',
    `identiti_id` varchar(255) DEFAULT NULL COMMENT '身份证号',
    `career` varchar(255) DEFAULT NULL COMMENT '职业',
    `origin_hometown` varchar(255) DEFAULT NULL COMMENT '原籍',
    `birth_place` varchar(255) DEFAULT NULL COMMENT '出生地',
    `marital_status` varchar(255) DEFAULT NULL COMMENT '婚姻状况',
    `home_address` varchar(255) DEFAULT NULL COMMENT '家庭住址',
    `current_problem` varchar(255) DEFAULT NULL COMMENT '现实问题',
    `education_history` varchar(255) DEFAULT NULL COMMENT '教育史',
    `health_state` varchar(255) DEFAULT NULL COMMENT '健康状况',
    `marital_status_explicit` varchar(255) DEFAULT NULL COMMENT '详细婚姻状况',
    `employment_history` varchar(255) DEFAULT NULL COMMENT '就业史',
    `contract_institusion_history` varchar(255) DEFAULT NULL COMMENT '以前与社会机构的接触',
    `hobby` varchar(255) DEFAULT NULL COMMENT '兴趣爱好',
    `other_situation` varchar(255) DEFAULT NULL COMMENT '其他需要说明的情况',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `seata_tcc_payment`.`profile` (`id`, `user_id`, `total`, `used`, `residue`) VALUES ('1', '1', '100000', '0', '100000');

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
