CREATE TABLE `user`
(
    `id`         varchar(100) NOT NULL,
    `name`       varchar(100) NOT NULL,
    `last_name`  varchar(100) DEFAULT NULL,
    `password`   varchar(100) NOT NULL,
    `username`   varchar(100) NOT NULL,
    `city`       varchar(100) NOT NULL,
    `country`    varchar(100) NOT NULL,
    `email`      varchar(100) DEFAULT NULL,
    `created_at` date         NOT NULL,
    `user_type`  varchar(100) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_UN` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8