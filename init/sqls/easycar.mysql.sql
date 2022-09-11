CREATE TABLE `branch`
(
    `id`           int           NOT NULL AUTO_INCREMENT,
    `branch_id`    varchar(100)  NOT NULL COMMENT 'branch_id',
    `url`          varchar(1000) NOT NULL COMMENT 'This action request address',
    `req_data`     varchar(1999) NOT NULL COMMENT 'the action request body',
    `tran_type`    varchar(20)   NOT NULL COMMENT 'transaction type suh as tcc,saga....',
    `protocol`     varchar(30)   NOT NULL DEFAULT 'http' COMMENT 'the action Network protocol  like http,grpc...',
    `action`       varchar(20)   NOT NULL COMMENT 'action like try confirm cancel ...',
    `state`        varchar(30)   NOT NULL DEFAULT 'branchReady' COMMENT 'branch state ',
    `level`        int           NOT NULL DEFAULT '1' COMMENT 'The execution priority of sub-transactions in a distributed transaction, the smaller the number, the first to be executed',
    `last_err_msg` varchar(255)  NOT NULL COMMENT 'last err message for the action ',
    `req_header`   varchar(1999) NOT NULL COMMENT 'the action request header',
    `timeout`      int           NOT NULL DEFAULT '0' COMMENT 'custom request timeout',
    `g_id`         varchar(25)   NOT NULL COMMENT 'global transaction id',
    `create_time`  int           NOT NULL DEFAULT '0',
    `update_time`  int           NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


CREATE TABLE `global`
(
    `id`             int         NOT NULL AUTO_INCREMENT,
    `g_id`           varchar(25) NOT NULL COMMENT 'global transaction id',
    `state`          varchar(30) NOT NULL DEFAULT 'begin' COMMENT 'global state  A state with two phases',
    `end_time`       int         NOT NULL DEFAULT '0' COMMENT 'end time it may be the commit end time or callback end time ',
    `next_cron_time` int         NOT NULL DEFAULT '0' COMMENT 'pending',
    `create_time`    int         NOT NULL DEFAULT '0',
    `update_time`    int         NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
