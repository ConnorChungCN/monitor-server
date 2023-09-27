
CREATE DATABASE IF NOT EXISTS algorithm_server
    DEFAULT CHARACTER SET = 'utf8mb4';


USE algorithm_server;

CREATE TABLE IF NOT EXISTS `task` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '任务自增ID',
    `task_id` VARCHAR(64) NOT NULL COMMENT '任务ID',
    `worker_id` VARCHAR(64) NOT NULL COMMENT 'worker ID',
    `algorithm_name` VARCHAR(32) NOT NULL COMMENT '算法名字',
    `algorithm_version` VARCHAR(16) NOT NULL COMMENT '算法版本',
    `status` SMALLINT NOT NULL COMMENT '任务状态',
    `executor_state` SMALLINT NOT NULL COMMENT '执行结果状态',
    `desc` TEXT NOT NULL COMMENT '描述',
    `input` TEXT NOT NULL COMMENT '任务输入',
    `output` TEXT NOT NULL COMMENT '任务输出',
    `ctime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `utime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `creator` VARCHAR(32) DEFAULT 'SYSTEM' NOT NULL COMMENT '创建者',
    `editor` VARCHAR(32) DEFAULT 'SYSTEM' NOT NULL COMMENT '更新者',
    PRIMARY KEY(`id`),
    UNIQUE KEY ux_task_id(`task_id`)
);


CREATE TABLE IF NOT EXISTS `algorithm` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '算法ID',
    `algorithm_id` VARCHAR(64) NOT NULL COMMENT '算法ID',
    `algorithm_type` SMALLINT NOT NULL COMMENT '算法类型',
    `part` SMALLINT NOT NULL COMMENT '算法目标部位',
    `algorithm_name` VARCHAR(32) NOT NULL COMMENT '算法名字',
    `ctime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `utime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `creator` VARCHAR(32) DEFAULT 'SYSTEM' NOT NULL COMMENT '创建者',
    `editor` VARCHAR(32) DEFAULT 'SYSTEM' NOT NULL COMMENT '更新者',
    PRIMARY KEY(`id`),
    UNIQUE KEY ux_algorithm_id(`algorithm_id`),
    UNIQUE KEY ux_algorithm_name(`algorithm_name`)
);


CREATE TABLE IF NOT EXISTS `algorithm_version` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '算法ID',
    `algorithm_id` VARCHAR(64) NOT NULL COMMENT '算法ID',
    `version_id` VARCHAR(64) NOT NULL COMMENT '算法版本ID',
    `algorithm_version` VARCHAR(32) NOT NULL COMMENT '算法类型',
    `storage_path` VARCHAR(128) NOT NULL COMMENT '算法名字',
    `algorithm_schema` TEXT NOT NULL COMMENT '算法schama',
    `upload` BOOLEAN NOT NULL COMMENT '镜像文件是否已经上传',
    `ctime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `utime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `creator` VARCHAR(32) DEFAULT 'SYSTEM' NOT NULL COMMENT '创建者',
    `editor` VARCHAR(32) DEFAULT 'SYSTEM' NOT NULL COMMENT '更新者',
    PRIMARY KEY(`id`),
    UNIQUE KEY ux_version_id(`version_id`),
    UNIQUE KEY ux_algoid_version(`algorithm_id`, `algorithm_version`)
);

CREATE TABLE IF NOT EXISTS `pipeline` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '流水线自增ID',
    `uuid` VARCHAR(64) NOT NULL COMMENT '流水线UUID',
    `name` VARCHAR(32) NOT NULL COMMENT '流水线名字',
    `template` TEXT NOT NULL COMMENT '流水线模板',
    `ctime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `utime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `creator` VARCHAR(32) DEFAULT 'SYSTEM' NOT NULL COMMENT '创建者',
    `editor` VARCHAR(32) DEFAULT 'SYSTEM' NOT NULL COMMENT '更新者',
    PRIMARY KEY(`id`),
    UNIQUE KEY ux_version_id(`uuid`)
);


CREATE TABLE IF NOT EXISTS `pipeline_execution_log` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '流水线执行记录自增ID',
    `uuid` VARCHAR(64) NOT NULL COMMENT '流水线执行记录UUID',
    `pipeline_uuid` VARCHAR(64)  COMMENT '流水线uuid',
    `template` TEXT NOT NULL COMMENT '流水线模板',
    `status` SMALLINT NOT NULL COMMENT '任务状态',
    `input` TEXT COMMENT '流水线输入',
    `output` TEXT COMMENT '流水线输出',
    `err_msg` TEXT COMMENT '错误信息',
    `ctime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `utime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `creator` VARCHAR(32) DEFAULT 'SYSTEM' NOT NULL COMMENT '创建者',
    `editor` VARCHAR(32) DEFAULT 'SYSTEM' NOT NULL COMMENT '更新者',
    PRIMARY KEY(`id`),
    UNIQUE KEY ux_version_id(`uuid`)
);


GRANT ALL PRIVILEGES ON algorithm_server.* TO 'hanglok'@'%';