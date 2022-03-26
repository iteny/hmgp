/*
 Navicat MySQL Data Transfer

 Source Server         : 我得
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : hmgp

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 27/03/2022 03:09:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for hm_account
-- ----------------------------
DROP TABLE IF EXISTS `hm_account`;
CREATE TABLE `hm_account`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_account
-- ----------------------------
INSERT INTO `hm_account` VALUES (1, 'admin', '90b9aa7e25f80cf4f64e990b78a9fc5ebd6cecad');

-- ----------------------------
-- Table structure for hm_menu
-- ----------------------------
DROP TABLE IF EXISTS `hm_menu`;
CREATE TABLE `hm_menu`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` int NOT NULL,
  `status` tinyint(1) NOT NULL,
  `pid` int NOT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `type` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_menu
-- ----------------------------
INSERT INTO `hm_menu` VALUES (1, '设置', '#', 1, 1, 0, 'fa-tachometer-alt', 1);
INSERT INTO `hm_menu` VALUES (2, '菜单管理', '/admin/menu', 1, 1, 1, 'nav-icon', 2);
INSERT INTO `hm_menu` VALUES (16, '啥地方地方', 'sadfasd', 2, 1, 1, 'nav-icon', 2);

SET FOREIGN_KEY_CHECKS = 1;
