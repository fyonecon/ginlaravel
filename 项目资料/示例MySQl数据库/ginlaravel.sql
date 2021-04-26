/*
 Navicat MySQL Data Transfer

 Source Server         : 192.168.131.7
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : 192.168.131.7:3306
 Source Schema         : ginlaravel

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 15/04/2021 17:55:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gl_app_token
-- ----------------------------
DROP TABLE IF EXISTS `gl_app_token`;
CREATE TABLE `gl_app_token` (
  `app_token_id` int(11) NOT NULL AUTO_INCREMENT,
  `app_token` varchar(2000) COLLATE utf8mb4_unicode_ci NOT NULL,
  `create_time` char(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`app_token_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of gl_app_token
-- ----------------------------
BEGIN;
INSERT INTO `gl_app_token` VALUES (1, 'abc', '20210101010101');
COMMIT;

-- ----------------------------
-- Table structure for gl_user
-- ----------------------------
DROP TABLE IF EXISTS `gl_user`;
CREATE TABLE `gl_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户列表',
  `user_class_id` int(11) NOT NULL,
  `nickname` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  `state` int(5) NOT NULL DEFAULT '1' COMMENT '1正常，2删除',
  `create_time` char(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '20210101000000',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of gl_user
-- ----------------------------
BEGIN;
INSERT INTO `gl_user` VALUES (1, 1, '张三', 1, '20210101120101');
INSERT INTO `gl_user` VALUES (2, 1, '李4⃣️', 1, '20210201120102');
INSERT INTO `gl_user` VALUES (3, 4, '中央汇金', 1, '20210301120103');
INSERT INTO `gl_user` VALUES (4, 2, '张坤', 1, '20210401120104');
INSERT INTO `gl_user` VALUES (5, 3, '红杉资本', 1, '20210405120105');
INSERT INTO `gl_user` VALUES (6, 1, '王5⃣️', 2, '20210101000000');
COMMIT;

-- ----------------------------
-- Table structure for gl_user_class
-- ----------------------------
DROP TABLE IF EXISTS `gl_user_class`;
CREATE TABLE `gl_user_class` (
  `user_class_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_class_name` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`user_class_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of gl_user_class
-- ----------------------------
BEGIN;
INSERT INTO `gl_user_class` VALUES (1, '散户');
INSERT INTO `gl_user_class` VALUES (2, '基金经理');
INSERT INTO `gl_user_class` VALUES (3, '股票机构');
INSERT INTO `gl_user_class` VALUES (4, '国家基金');
COMMIT;

-- ----------------------------
-- Table structure for gl_user_token
-- ----------------------------
DROP TABLE IF EXISTS `gl_user_token`;
CREATE TABLE `gl_user_token` (
  `user_token_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户登录的token',
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `user_token` varchar(2000) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`user_token_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of gl_user_token
-- ----------------------------
BEGIN;
INSERT INTO `gl_user_token` VALUES (1, 1, 'cfvbnedrftgh');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
