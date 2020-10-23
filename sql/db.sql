CREATE DATABASE IF NOT EXISTS db_stat;
use db_stat;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `accountid` int(10) UNSIGNED NOT NULL COMMENT '账号id',
  `ostype` int(10) UNSIGNED NOT NULL COMMENT '操作系统类型 1/ios 2/and 3/win',
  `channel` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '热更新资源渠道',
  `date_fk` int(10) UNSIGNED NOT NULL COMMENT '账号创建日期',
  `daytime` int(10) UNSIGNED NOT NULL COMMENT '账号创建时间',
  `ip` char(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账号创建ip',
  `lang` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账号创建语言',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `accountid_idx`(`accountid`) USING BTREE,
  INDEX `date_accountid_idx`(`date_fk`, `accountid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for date
-- ----------------------------
DROP TABLE IF EXISTS `date`;
CREATE TABLE `date`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `yyyymmdd` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '日期，例如2020-08-21',
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '当日描述',
  `year` int(10) UNSIGNED NOT NULL COMMENT '年',
  `month` tinyint(10) UNSIGNED NOT NULL COMMENT '月',
  `day` tinyint(10) NOT NULL COMMENT '日',
  `week` tinyint(10) UNSIGNED NOT NULL COMMENT '星期(1-7)',
  `ymd` int(10) UNSIGNED NOT NULL COMMENT '日期，例如20200821',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `ymd_idx`(`ymd`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for login
-- ----------------------------
DROP TABLE IF EXISTS `login`;
CREATE TABLE `login`  (
  `rid` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `zoneid_fk` int(10) UNSIGNED NOT NULL COMMENT '分区id',
  `accountid_fk` int(10) UNSIGNED NOT NULL COMMENT '账号id',
  `date_fk` int(10) UNSIGNED NOT NULL COMMENT '登录日期',
  `daytime` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录时间',
  PRIMARY KEY (`rid`) USING BTREE,
  INDEX `date_zone_account_time`(`date_fk`, `zoneid_fk`, `accountid_fk`, `daytime`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for logout
-- ----------------------------
DROP TABLE IF EXISTS `logout`;
CREATE TABLE `logout`  (
  `rid` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `zoneid_fk` int(10) UNSIGNED NOT NULL COMMENT '分区id',
  `accountid_fk` int(10) NOT NULL COMMENT '账号id',
  `date_fk` int(10) NOT NULL COMMENT '日期',
  `daytime` int(10) NOT NULL COMMENT '时间',
  `online_time` int(10) UNSIGNED NOT NULL COMMENT '在线时长秒',
  PRIMARY KEY (`rid`) USING BTREE,
  INDEX `date_zone_account_time`(`date_fk`, `zoneid_fk`, `accountid_fk`, `daytime`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for recharge
-- ----------------------------
DROP TABLE IF EXISTS `recharge`;
CREATE TABLE `recharge`  (
  `rid` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `zoneid_fk` int(10) UNSIGNED NOT NULL COMMENT '分区id',
  `accountid_fk` int(10) UNSIGNED NOT NULL COMMENT '账号id',
  `date_fk` int(11) UNSIGNED NOT NULL COMMENT '日期',
  `daytime` int(11) UNSIGNED NOT NULL COMMENT '时间',
  `product_id` int(11) UNSIGNED NOT NULL COMMENT '商品id',
  `money` int(10) UNSIGNED NOT NULL COMMENT '商品价格',
  `first` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否是首次充值',
  PRIMARY KEY (`rid`) USING BTREE,
  INDEX `date_zoneid_accountid_time`(`date_fk`, `zoneid_fk`, `accountid_fk`, `daytime`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `zoneid_fk` int(10) UNSIGNED NOT NULL COMMENT '分区id',
  `accountid_fk` int(10) UNSIGNED NOT NULL COMMENT '账号id',
  `first` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否是全服唯一角色',
  `reg_date_fk` int(10) NOT NULL COMMENT '创角日期',
  `daytime` int(10) UNSIGNED NOT NULL COMMENT '创角时间',
  `login_1` bigint(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '90日留存',
  `login_2` bigint(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '90日留存',
  `last_login_date_fk` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后一次登陆日期',
  `rge_total_3` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '3日累计充值，单位分',
  `rge_total_7` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '7日累计充值，单位分',
  `rge_total_14` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '14日累计充值，单位分',
  `rge_total_30` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '30日累计充值，单位分',
  `rge_total` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '累计充值，单位分',
  `rge_day1` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day1',
  `rge_day2` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day2',
  `rge_day3` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day3',
  `rge_day4` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day4',
  `rge_day5` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day5',
  `rge_day6` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day6',
  `rge_day7` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day7',
  `rge_day8` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day8',
  `rge_day9` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day9',
  `rge_day10` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day10',
  `rge_day11` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day11',
  `rge_day12` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day12',
  `rge_day13` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day13',
  `rge_day14` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day14',
  `rge_day15` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day15',
  `rge_day16` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day16',
  `rge_day17` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day17',
  `rge_day18` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day18',
  `rge_day19` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day19',
  `rge_day20` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day20',
  `rge_day21` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day21',
  `rge_day22` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day22',
  `rge_day23` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day23',
  `rge_day24` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day24',
  `rge_day25` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day25',
  `rge_day26` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day26',
  `rge_day27` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day27',
  `rge_day28` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day28',
  `rge_day29` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day29',
  `rge_day30` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day30',
  `rge_day31` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day31',
  `rge_day32` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day32',
  `rge_day33` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day33',
  `rge_day34` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day34',
  `rge_day35` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day35',
  `rge_day36` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day36',
  `rge_day37` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day37',
  `rge_day38` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day38',
  `rge_day39` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day39',
  `rge_day40` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day40',
  `rge_day41` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day41',
  `rge_day42` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day42',
  `rge_day43` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day43',
  `rge_day44` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day44',
  `rge_day45` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day45',
  `rge_day46` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day46',
  `rge_day47` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day47',
  `rge_day48` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day48',
  `rge_day49` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day49',
  `rge_day50` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day50',
  `rge_day51` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day51',
  `rge_day52` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day52',
  `rge_day53` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day53',
  `rge_day54` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day54',
  `rge_day55` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day55',
  `rge_day56` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day56',
  `rge_day57` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day57',
  `rge_day58` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day58',
  `rge_day59` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day59',
  `rge_day60` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day60',
  `rge_day61` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day61',
  `rge_day62` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day62',
  `rge_day63` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day63',
  `rge_day64` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day64',
  `rge_day65` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day65',
  `rge_day66` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day66',
  `rge_day67` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day67',
  `rge_day68` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day68',
  `rge_day69` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day69',
  `rge_day70` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day70',
  `rge_day71` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day71',
  `rge_day72` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day72',
  `rge_day73` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day73',
  `rge_day74` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day74',
  `rge_day75` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day75',
  `rge_day76` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day76',
  `rge_day77` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day77',
  `rge_day78` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day78',
  `rge_day79` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day79',
  `rge_day80` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day80',
  `rge_day81` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day81',
  `rge_day82` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day82',
  `rge_day83` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day83',
  `rge_day84` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day84',
  `rge_day85` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day85',
  `rge_day86` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day86',
  `rge_day87` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day87',
  `rge_day88` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day88',
  `rge_day89` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day89',
  `rge_day90` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值day90',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `zoneid_accountid_idx`(`zoneid_fk`, `accountid_fk`) USING BTREE,
  INDEX `reg_zoneid_accountid_idx`(`reg_date_fk`, `zoneid_fk`, `accountid_fk`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for sync_rid
-- ----------------------------
DROP TABLE IF EXISTS `sync_rid`;
CREATE TABLE `sync_rid`  (
  `table` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '表名',
  `zoneid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '分区id',
  `rid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上一次同步进度',
  PRIMARY KEY (`table`, `zoneid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for welfare_roles
-- ----------------------------
DROP TABLE IF EXISTS `welfare_roles`;
CREATE TABLE `welfare_roles`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `zoneid` int(10) UNSIGNED NOT NULL COMMENT '分区id',
  `roleid` int(10) UNSIGNED NOT NULL COMMENT '玩家id',
  `time` datetime(0) NOT NULL COMMENT '执行时间',
  `cmd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '执行命令',
  `status` int(5) UNSIGNED NOT NULL COMMENT '0 待执行 1已执行',
  `taskid_pk` int(10) UNSIGNED NOT NULL COMMENT 'welfare_task 外键',
  `exec_time` datetime(0) NULL DEFAULT NULL COMMENT '执行时间',
  `exec_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '执行结果',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `task_zoneid_roleid_idx`(`time`, `zoneid`, `roleid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for welfare_task
-- ----------------------------
DROP TABLE IF EXISTS `welfare_task`;
CREATE TABLE `welfare_task`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `roles` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '玩家列表(分区id,玩家id)，以分号分隔',
  `cmds` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'gm命令，以分号分隔',
  `cmd_time` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '每天执行时间范围 开始时间-结束时间',
  `status` int(10) UNSIGNED NOT NULL COMMENT ' 0停用 1启用',
  `begin_time` date NOT NULL COMMENT '开始日期',
  `end_time` date NOT NULL COMMENT '结束日期',
  `cur_time` date NULL DEFAULT NULL COMMENT '当前已完成发放日期',
  `step` int(10) UNSIGNED NOT NULL DEFAULT 1 COMMENT '默认发放时间间隔',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for zone
-- ----------------------------
DROP TABLE IF EXISTS `zone`;
CREATE TABLE `zone`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `zoneid` int(11) NOT NULL COMMENT '分区id',
  `zonename` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分区名字',
  `openday_fk` int(10) UNSIGNED NOT NULL COMMENT '开服日期',
  `logdbhost` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '日志数据库地址',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
