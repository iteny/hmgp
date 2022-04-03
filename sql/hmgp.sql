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

 Date: 04/04/2022 01:08:40
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
-- Table structure for hm_gp_category
-- ----------------------------
DROP TABLE IF EXISTS `hm_gp_category`;
CREATE TABLE `hm_gp_category`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `sort` tinyint UNSIGNED NOT NULL,
  `status` tinyint UNSIGNED NOT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_gp_category
-- ----------------------------
INSERT INTO `hm_gp_category` VALUES (1, '撒旦法', '/asdfsa', 1, 1, '0xe620');
INSERT INTO `hm_gp_category` VALUES (2, '奥法', 'sad', 1, 1, '0xe620');
INSERT INTO `hm_gp_category` VALUES (3, '撒旦法', 'sadfasd', 1, 1, '0xe620');
INSERT INTO `hm_gp_category` VALUES (4, '萨法是', 'sfasdf', 1, 1, '0xe620');
INSERT INTO `hm_gp_category` VALUES (5, '撒地方的', 'sadfasd', 1, 1, '0xe620');
INSERT INTO `hm_gp_category` VALUES (6, '撒发大水', 'safdsadf', 1, 1, '0xe620');
INSERT INTO `hm_gp_category` VALUES (7, '阿斯蒂芬', 'asdfdsa', 1, 1, '0xe620');
INSERT INTO `hm_gp_category` VALUES (8, '萨法是', 'asfsadf', 1, 1, '0xe605');

-- ----------------------------
-- Table structure for hm_gp_category_content
-- ----------------------------
DROP TABLE IF EXISTS `hm_gp_category_content`;
CREATE TABLE `hm_gp_category_content`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `pid` int UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_gp_category_content
-- ----------------------------
INSERT INTO `hm_gp_category_content` VALUES (1, '奥术飞弹', '阿士大夫撒地方', 1);
INSERT INTO `hm_gp_category_content` VALUES (2, '奥法', '人工费', 1);
INSERT INTO `hm_gp_category_content` VALUES (3, '天然芙蓉', '地方和个人', 2);
INSERT INTO `hm_gp_category_content` VALUES (4, '地方广东省', '大哥豆腐乳', 2);

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
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hm_menu
-- ----------------------------
INSERT INTO `hm_menu` VALUES (1, '设置', '#', 1, 1, 0, 'fa-tachometer-alt', 1);
INSERT INTO `hm_menu` VALUES (2, '菜单管理', '/admin/menu', 1, 1, 1, 'nav-icon', 2);
INSERT INTO `hm_menu` VALUES (17, '河马股票', '#', 2, 1, 0, 'fa-building', 1);
INSERT INTO `hm_menu` VALUES (18, '分类', '/admin/gpCategory', 1, 1, 17, 'fa-arrows-alt', 2);

SET FOREIGN_KEY_CHECKS = 1;
