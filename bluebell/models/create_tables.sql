
-- 创建用户表
CREATE TABLE USER (
                      id BIGINT (20) NOT NULL auto_increment,
                      user_id BIGINT (20) NOT NULL,
                      username VARCHAR (54) NOT NULL,
                      `password` VARCHAR (64) NOT NULL,
                      email VARCHAR (64),
                      gender TINYINT (4) NOT NULL DEFAULT "0",
                      create_time TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
                      update_time TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      PRIMARY KEY (id),
                      UNIQUE KEY idx_username (username) USING BTREE,
                      UNIQUE KEY idx_user_id (user_id) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4