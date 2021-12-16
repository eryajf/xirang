SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for groups
-- ----------------------------
DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
  `group_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户组ID',
  `group_name` varchar(255) DEFAULT NULL COMMENT '用户组名',
  `state` int(11) DEFAULT '0' COMMENT '状态：1 启用 0 禁用',
  `desc` varchar(128) DEFAULT NULL COMMENT '备注',
  `create_by` varchar(128) DEFAULT NULL COMMENT '谁创建',
  `update_by` varchar(128) DEFAULT NULL COMMENT '谁更新',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`group_id`),
  KEY `idx_groups_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of groups
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_group_relations
-- ----------------------------
DROP TABLE IF EXISTS `user_group_relations`;
CREATE TABLE `user_group_relations` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_group_relations_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user_group_relations
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `job_number` varchar(255) DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` char(11) DEFAULT NULL,
  `gender` varchar(255) DEFAULT NULL,
  `state` int(11) DEFAULT '0',
  `avatar_url` varchar(255) DEFAULT 'static/upload/avatar/default.png',
  `role_id` int(11) DEFAULT NULL,
  `dept_id` int(11) DEFAULT NULL,
  `username` varchar(64) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
