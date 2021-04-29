-- MySQL dump 10.13  Distrib 5.6.17, for linux-glibc2.5 (x86_64)
--
-- Host: 172.19.248.171    Database: db_stat
-- ------------------------------------------------------
-- Server version	5.6.17

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `db_stat`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `db_stat` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `db_stat`;

--
-- Table structure for table `account`
--

DROP TABLE IF EXISTS `account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `account` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `accountid` int(10) unsigned NOT NULL COMMENT '账号id',
  `ostype` int(10) unsigned NOT NULL COMMENT '操作系统类型 1/ios 2/and 3/win',
  `channel` varchar(20) NOT NULL COMMENT '热更新资源渠道',
  `date_fk` int(10) unsigned NOT NULL COMMENT '账号创建日期',
  `daytime` int(10) unsigned NOT NULL COMMENT '账号创建时间',
  `ip` char(15) NOT NULL COMMENT '账号创建ip',
  `lang` char(2) NOT NULL COMMENT '账号创建语言',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `accountid_idx` (`accountid`) USING BTREE,
  KEY `date_accountid_idx` (`date_fk`,`accountid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `chat_dirty_history`
--

DROP TABLE IF EXISTS `chat_dirty_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chat_dirty_history` (
  `_rid` int(10) unsigned NOT NULL,
  `time` datetime DEFAULT NULL,
  `zoneid` int(10) unsigned DEFAULT NULL,
  `mapid` int(10) unsigned DEFAULT NULL,
  `type` int(10) unsigned DEFAULT NULL,
  `fromroleid` bigint(20) unsigned DEFAULT NULL,
  `fromrolename` varchar(256) DEFAULT NULL,
  `tozoneid` int(10) unsigned DEFAULT NULL,
  `toroleid` bigint(20) unsigned DEFAULT NULL,
  `torolename` varchar(256) DEFAULT NULL,
  `allianceid` bigint(20) unsigned DEFAULT NULL,
  `alliancename` varchar(256) DEFAULT NULL,
  `content` varchar(256) DEFAULT NULL,
  `dirtyword` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`_rid`),
  KEY `time` (`time`),
  KEY `zoneid_fromroleid_time` (`zoneid`,`fromroleid`,`time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `chat_dirty_word`
--

DROP TABLE IF EXISTS `chat_dirty_word`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chat_dirty_word` (
  `id` int(10) unsigned NOT NULL,
  `words` longtext CHARACTER SET utf8mb4,
  `sync_rid` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `client_dispose`
--

DROP TABLE IF EXISTS `client_dispose`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `client_dispose` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `client_version` varchar(255) NOT NULL,
  `stackmd5` char(32) NOT NULL,
  `stack` text NOT NULL,
  `status` int(10) NOT NULL,
  `note` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `date`
--

DROP TABLE IF EXISTS `date`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `date` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `yyyymmdd` char(10) NOT NULL COMMENT '日期，例如2020-08-21',
  `desc` text NOT NULL COMMENT '当日描述',
  `year` int(10) unsigned NOT NULL COMMENT '年',
  `month` tinyint(10) unsigned NOT NULL COMMENT '月',
  `day` tinyint(10) NOT NULL COMMENT '日',
  `week` tinyint(10) unsigned NOT NULL COMMENT '星期(1-7)',
  `ymd` int(10) unsigned NOT NULL COMMENT '日期，例如20200821',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `ymd_idx` (`ymd`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `log`
--

DROP TABLE IF EXISTS `log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log` (
  `time` datetime NOT NULL COMMENT '记录时间',
  `username` varchar(50) CHARACTER SET latin1 NOT NULL COMMENT '用户名',
  `action` varchar(100) CHARACTER SET latin1 NOT NULL COMMENT '行为',
  `desc` varchar(255) CHARACTER SET latin1 NOT NULL COMMENT '行为描述'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `login`
--

DROP TABLE IF EXISTS `login`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `login` (
  `rid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `zoneid_fk` int(10) unsigned NOT NULL COMMENT '分区id',
  `accountid_fk` int(10) unsigned NOT NULL COMMENT '账号id',
  `date_fk` int(10) unsigned NOT NULL COMMENT '登录日期',
  `daytime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '登录时间',
  PRIMARY KEY (`rid`) USING BTREE,
  KEY `date_zone_account_time` (`date_fk`,`zoneid_fk`,`accountid_fk`,`daytime`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `logout`
--

DROP TABLE IF EXISTS `logout`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `logout` (
  `rid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `zoneid_fk` int(10) unsigned NOT NULL COMMENT '分区id',
  `accountid_fk` int(10) NOT NULL COMMENT '账号id',
  `date_fk` int(10) NOT NULL COMMENT '日期',
  `daytime` int(10) NOT NULL COMMENT '时间',
  `online_time` int(10) unsigned NOT NULL COMMENT '在线时长秒',
  PRIMARY KEY (`rid`) USING BTREE,
  KEY `date_zone_account_time` (`date_fk`,`zoneid_fk`,`accountid_fk`,`daytime`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `recharge`
--

DROP TABLE IF EXISTS `recharge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recharge` (
  `rid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `zoneid_fk` int(10) unsigned NOT NULL COMMENT '分区id',
  `accountid_fk` int(10) unsigned NOT NULL COMMENT '账号id',
  `date_fk` int(11) unsigned NOT NULL COMMENT '日期',
  `daytime` int(11) unsigned NOT NULL COMMENT '时间',
  `product_id` int(11) unsigned NOT NULL COMMENT '商品id',
  `money` int(10) unsigned NOT NULL COMMENT '商品价格',
  `first` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否是首次充值',
  PRIMARY KEY (`rid`) USING BTREE,
  KEY `date_zoneid_accountid_time` (`date_fk`,`zoneid_fk`,`accountid_fk`,`daytime`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `role`
--

DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `zoneid_fk` int(10) unsigned NOT NULL COMMENT '分区id',
  `accountid_fk` int(10) unsigned NOT NULL COMMENT '账号id',
  `first` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否是全服唯一角色',
  `reg_date_fk` int(10) NOT NULL COMMENT '创角日期',
  `daytime` int(10) unsigned NOT NULL COMMENT '创角时间',
  `login_1` bigint(10) unsigned NOT NULL DEFAULT '0' COMMENT '90日留存',
  `login_2` bigint(10) unsigned NOT NULL DEFAULT '0' COMMENT '90日留存',
  `last_login_date_fk` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次登陆日期',
  `rge_total_3` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '3日累计充值，单位分',
  `rge_total_7` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '7日累计充值，单位分',
  `rge_total_14` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '14日累计充值，单位分',
  `rge_total_30` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '30日累计充值，单位分',
  `rge_total` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '累计充值，单位分',
  `rge_day1` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day1',
  `rge_day2` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day2',
  `rge_day3` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day3',
  `rge_day4` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day4',
  `rge_day5` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day5',
  `rge_day6` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day6',
  `rge_day7` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day7',
  `rge_day8` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day8',
  `rge_day9` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day9',
  `rge_day10` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day10',
  `rge_day11` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day11',
  `rge_day12` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day12',
  `rge_day13` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day13',
  `rge_day14` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day14',
  `rge_day15` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day15',
  `rge_day16` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day16',
  `rge_day17` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day17',
  `rge_day18` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day18',
  `rge_day19` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day19',
  `rge_day20` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day20',
  `rge_day21` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day21',
  `rge_day22` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day22',
  `rge_day23` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day23',
  `rge_day24` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day24',
  `rge_day25` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day25',
  `rge_day26` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day26',
  `rge_day27` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day27',
  `rge_day28` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day28',
  `rge_day29` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day29',
  `rge_day30` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day30',
  `rge_day31` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day31',
  `rge_day32` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day32',
  `rge_day33` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day33',
  `rge_day34` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day34',
  `rge_day35` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day35',
  `rge_day36` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day36',
  `rge_day37` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day37',
  `rge_day38` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day38',
  `rge_day39` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day39',
  `rge_day40` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day40',
  `rge_day41` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day41',
  `rge_day42` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day42',
  `rge_day43` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day43',
  `rge_day44` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day44',
  `rge_day45` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day45',
  `rge_day46` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day46',
  `rge_day47` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day47',
  `rge_day48` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day48',
  `rge_day49` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day49',
  `rge_day50` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day50',
  `rge_day51` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day51',
  `rge_day52` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day52',
  `rge_day53` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day53',
  `rge_day54` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day54',
  `rge_day55` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day55',
  `rge_day56` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day56',
  `rge_day57` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day57',
  `rge_day58` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day58',
  `rge_day59` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day59',
  `rge_day60` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day60',
  `rge_day61` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day61',
  `rge_day62` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day62',
  `rge_day63` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day63',
  `rge_day64` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day64',
  `rge_day65` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day65',
  `rge_day66` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day66',
  `rge_day67` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day67',
  `rge_day68` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day68',
  `rge_day69` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day69',
  `rge_day70` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day70',
  `rge_day71` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day71',
  `rge_day72` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day72',
  `rge_day73` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day73',
  `rge_day74` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day74',
  `rge_day75` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day75',
  `rge_day76` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day76',
  `rge_day77` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day77',
  `rge_day78` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day78',
  `rge_day79` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day79',
  `rge_day80` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day80',
  `rge_day81` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day81',
  `rge_day82` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day82',
  `rge_day83` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day83',
  `rge_day84` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day84',
  `rge_day85` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day85',
  `rge_day86` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day86',
  `rge_day87` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day87',
  `rge_day88` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day88',
  `rge_day89` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day89',
  `rge_day90` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '充值day90',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `zoneid_accountid_idx` (`zoneid_fk`,`accountid_fk`) USING BTREE,
  KEY `reg_zoneid_accountid_idx` (`reg_date_fk`,`zoneid_fk`,`accountid_fk`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sync_rid`
--

DROP TABLE IF EXISTS `sync_rid`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sync_rid` (
  `table` varchar(50) NOT NULL COMMENT '表名',
  `zoneid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '分区id',
  `rid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上一次同步进度',
  PRIMARY KEY (`table`,`zoneid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_action`
--

DROP TABLE IF EXISTS `user_action`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_action` (
  `action_name` varchar(80) NOT NULL DEFAULT '',
  `action` varchar(80) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `welfare_roles`
--

DROP TABLE IF EXISTS `welfare_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `welfare_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `zoneid` int(10) unsigned NOT NULL COMMENT '分区id',
  `roleid` int(10) unsigned NOT NULL COMMENT '玩家id',
  `mapid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '地图id',
  `time` datetime NOT NULL COMMENT '执行时间',
  `cmd` varchar(255) NOT NULL COMMENT '执行命令',
  `status` int(5) unsigned NOT NULL COMMENT '0 待执行 1已执行',
  `taskid_pk` int(10) unsigned NOT NULL COMMENT 'welfare_task 外键',
  `exec_time` datetime DEFAULT NULL COMMENT '执行时间',
  `exec_result` varchar(255) DEFAULT NULL COMMENT '执行结果',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `task_zoneid_roleid_idx` (`time`,`zoneid`,`roleid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `welfare_task`
--

DROP TABLE IF EXISTS `welfare_task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `welfare_task` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '描述',
  `roles` text NOT NULL COMMENT '玩家列表(分区id,玩家id)，以分号分隔',
  `cmds` text NOT NULL COMMENT 'gm命令，以分号分隔',
  `cmd_time` varchar(200) NOT NULL COMMENT '每天执行时间范围 开始时间-结束时间',
  `status` int(10) unsigned NOT NULL COMMENT ' 0停用 1启用',
  `begin_time` date NOT NULL COMMENT '开始日期',
  `end_time` date NOT NULL COMMENT '结束日期',
  `cur_time` date DEFAULT NULL COMMENT '当前已完成发放日期',
  `step` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '默认发放时间间隔',
  `slg` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '是否是sgl福利',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `zone`
--

DROP TABLE IF EXISTS `zone`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zone` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `zoneid` int(11) NOT NULL COMMENT '分区id',
  `zonename` varchar(50) NOT NULL COMMENT '分区名字',
  `openday_fk` int(10) unsigned NOT NULL COMMENT '开服日期',
  `logdbhost` varchar(15) NOT NULL DEFAULT '' COMMENT '日志数据库地址',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-04-29 17:21:29
