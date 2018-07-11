-- CREATE DATABASE IF NOT EXISTS `goredisorm`;
-- GRANT ALL PRIVILEGES ON *.* TO `goredisorm_user`@'%' IDENTIFIED BY 'goredisorm_pass';
-- GRANT ALL PRIVILEGES ON *.* TO 'goredisorm_user'@'localhost' IDENTIFIED BY 'goredisorm_pass';
-- GRANT ALL PRIVILEGES ON `goredisorm`.* TO `redisorm_user`@'%';
-- FLUSH PRIVILEGES;

USE goredisorm;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT,
  `name`        VARCHAR(32)      NOT NULL  DEFAULT '',
  `mailbox`     VARCHAR(128)     NOT NULL  DEFAULT '',
  `sex`         TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,
  `age`         INT UNSIGNED     NOT NULL  DEFAULT 0,
  `description` VARCHAR(256)     NOT NULL  DEFAULT '',
  `password`    VARCHAR(32)      NOT NULL  DEFAULT '',
  `head_url`    VARCHAR(255)     NOT NULL  DEFAULT '',
  `longitude`   float                                  DEFAULT NULL
  COMMENT 'longitude',
  `latitude`    float                                  DEFAULT NULL
  COMMENT 'latitude',
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,
  `created_at`  BIGINT(20)      NOT NULL DEFAULT '0',
  `updated_at`  BIGINT(20)      NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `blogs`;
CREATE TABLE `blogs` (
  `id`          INT UNSIGNED     NOT NULL  PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT UNSIGNED     NOT NULL DEFAULT 0,
  `title`       VARCHAR(32)      NOT NULL  DEFAULT '',
  `content`     TEXT             NOT NULL ,
  `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,
  `readed`  INT UNSIGNED     NOT NULL DEFAULT 0,
  `created_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `updated_at`   TIMESTAMP       NOT NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE VIEW user_base_info AS SELECT `id`,`name`,`mailbox`,`sex` FROM users;
