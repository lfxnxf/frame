# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: rm-2zel0996j0yjzc52w.mysql.rds.aliyuncs.com (MySQL 5.6.16-log)
# Database: shopping_store
# Generation Time: 2019-11-26 11:58:30 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table shopping_goods
# ------------------------------------------------------------

DROP TABLE IF EXISTS `shopping_goods`;

CREATE TABLE `shopping_goods` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `type` int(11) unsigned NOT NULL COMMENT '商品类型id',
  `name` varchar(256) NOT NULL COMMENT '商品名称',
  `goods_pic` varchar(512) NOT NULL COMMENT '商品图片',
  `goods_small_pic` varchar(512) NOT NULL COMMENT '商品缩略图',
  `goods_detail_pic` varchar(512) NOT NULL COMMENT '商品详情页图片，多个图片中间用分号隔开',
  `goods_detail_dyna_pic` varchar(512) NOT NULL COMMENT '商品详情页动效图',
  `small_pic_res_id` int(11) NOT NULL COMMENT '缩略图资源id',
  `desc_style` int(11) NOT NULL COMMENT '商品说明样式',
  `desc_text` text NOT NULL COMMENT '商品说明文案',
  `label` varchar(128) NOT NULL DEFAULT '' COMMENT '商品标签， 新品限时折扣、稀有',
  `cost_detail` varchar(4096) NOT NULL DEFAULT '' COMMENT '[{“id”:1, "价格":12,"使用天数":30,"折扣后价格":1234}]',
  `effect_id` int(11) NOT NULL COMMENT '第三方的id，效果关联，如特权id 或者 翻倍卡id',
  `effect_scene` int(11) NOT NULL DEFAULT '0' COMMENT '0 所有 1 直播间展示',
  `can_equip` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1.可装备 0不可装备',
  `extra` varchar(512) NOT NULL DEFAULT '' COMMENT '其他',
  `operator` varchar(128) NOT NULL,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `name` (`name`(255)),
  KEY `type` (`type`),
  KEY `update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商城商品信息';



# Dump of table shopping_goods_type
# ------------------------------------------------------------

DROP TABLE IF EXISTS `shopping_goods_type`;

CREATE TABLE `shopping_goods_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '类别名称，如座驾',
  `desc` varchar(512) NOT NULL DEFAULT '' COMMENT '描述',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品类型';



# Dump of table shopping_shelves_goods
# ------------------------------------------------------------

DROP TABLE IF EXISTS `shopping_shelves_goods`;

CREATE TABLE `shopping_shelves_goods` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `type` int(11) unsigned NOT NULL COMMENT '商品列表分类，如虚拟商城商品列表',
  `goods_id` int(11) unsigned NOT NULL COMMENT '商品id',
  `order` int(11) unsigned NOT NULL COMMENT '商品排序',
  `online_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上架时间',
  `offline_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下架时间',
  `operator` varchar(128) DEFAULT NULL,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `goods_id` (`goods_id`,`type`),
  KEY `goods_id_type` (`goods_id`,`type`),
  KEY `type` (`type`),
  KEY `update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='货架商品列表';



# Dump of table shopping_shelves_type
# ------------------------------------------------------------

DROP TABLE IF EXISTS `shopping_shelves_type`;

CREATE TABLE `shopping_shelves_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '类别名称，如虚拟商城商品列表',
  `desc` varchar(512) NOT NULL DEFAULT '' COMMENT '描述',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `operator` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品货架分类';



# Dump of table user_goods_order
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_goods_order`;

CREATE TABLE `user_goods_order` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` varchar(128) NOT NULL COMMENT '订单id',
  `goods_id` int(11) NOT NULL COMMENT '商品id',
  `num_diff` int(11) NOT NULL COMMENT '数量差值',
  `time_diff` int(11) NOT NULL COMMENT '时间',
  `type` varchar(128) NOT NULL COMMENT '类型',
  `params` varchar(4096) NOT NULL COMMENT '回放参数',
  `uid` int(11) NOT NULL COMMENT '用户uid',
  `status` int(11) NOT NULL COMMENT '订单状态',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_order_id` (`order_id`),
  KEY `idx_uid` (`uid`),
  KEY `type` (`type`),
  KEY `update_time` (`update_time`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品购买订单记录';




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
