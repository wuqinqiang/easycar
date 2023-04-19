CREATE TABLE IF NOT EXISTS `global`(
    `id`             int         NOT NULL AUTO_INCREMENT,
    `g_id`           varchar(25) NOT NULL COMMENT 'global transaction id',
    `state`          varchar(30) NOT NULL DEFAULT 'begin' COMMENT 'global state  A state with two phases',
    `end_time`       int         NOT NULL DEFAULT '0' COMMENT 'end time it may be the commit end time or callback end time ',
    `next_cron_time` int         NOT NULL DEFAULT '0' COMMENT 'pending',
    `create_time`    int         NOT NULL DEFAULT '0',
    `update_time`    int         NOT NULL DEFAULT '0',
    `try_times`      int         NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_gid` (`g_id`),
    KEY              `idx_next_cron_time` (`next_cron_time`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb3;
