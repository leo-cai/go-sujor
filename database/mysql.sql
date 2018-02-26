# dbhost='localhost';
# dbport=3306;
# dbuser='root';
# dbpassword='';

DROP DATABASE sujor;

CREATE DATABASE IF NOT EXISTS sujor DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

use sujor;

CREATE TABLE `t_user` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(32) NOT NULL,
  `password` VARCHAR(64) NOT NULL,
  `signed_at` TIMESTAMP,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_role` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `role_name` VARCHAR(32) NOT NULL,
  `description` VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_permission` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `permission_name` VARCHAR(32) NOT NULL,
  `description` VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_role` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `user_id` INT(11) NOT NULL,
  `role_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_role_t_role_1` (`role_id`),
  KEY `fk_user_role_t_user_1` (`user_id`),
  CONSTRAINT `fk_user_role_t_role_1` FOREIGN KEY (`role_id`) REFERENCES `t_role` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_user_role_t_user_1` FOREIGN KEY (`user_id`) REFERENCES `t_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `role_permission` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `role_id` INT(11) NOT NULL,
  `permission_id` INT(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_role_permission_t_permission_1` (`permission_id`),
  KEY `fk_role_permission_t_role_1` (`role_id`),
  CONSTRAINT `fk_role_permission_t_permission_1` FOREIGN KEY (`permission_id`) REFERENCES `t_permission` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_role_permission_t_role_1` FOREIGN KEY (`role_id`) REFERENCES `t_role` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 注释部分为实际 golang 操作 sql 语句
-- INSERT INTO `t_user`(username) VALUES ('admin');
INSERT INTO `t_role`(role_name, description) VALUES ('管理员', NULL), ('用户', NULL);
INSERT INTO `t_permission`(permission_name, description) VALUES ('用户管理', NULL), ('标题管理', NULL), ('留言管理', NULL), ('文章管理', NULL);
-- INSERT INTO `user_role`(user_id, role_id) VALUES (1, 1);
INSERT INTO `role_permission`(role_id, permission_id) VALUES (1, 1), (1, 2), (1, 3), (1, 4), (2, 4);

-- 参考 https://www.cnblogs.com/jway1101/p/5789378.html

-- 查询某用户拥有的权限
SELECT p.`permission_name`
FROM
t_permission p,role_permission rp,t_role r
WHERE
r.id=rp.role_id AND rp.permission_id=p.id AND r.id
IN
(SELECT r.id
FROM
   t_user u,t_role r,user_role ur
WHERE
  u.username ='admin2' AND u.id=ur.user_id AND ur.role_id=r.id);


CREATE TABLE `project` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(64) NULL,
  `author` VARCHAR (32) NULL,
  `content` LONGTEXT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO project (title, author, content) VALUES ("biaoti", "zuozhe", "wenben");