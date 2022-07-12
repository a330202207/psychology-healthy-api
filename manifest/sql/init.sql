CREATE TABLE `sys_config`
(
    `id`         bigint                                                  NOT NULL COMMENT '配置ID',
    `name`       varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '参数名称',
    `key`        varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '参数键名',
    `value`      varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '参数键值',
    `status`     tinyint                                                 NOT NULL DEFAULT '1' COMMENT '状态：10-开启，20-关闭',
    `remark`     varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '备注',
    `created_at` datetime                                                NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime                                                NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='参数配置表';

CREATE TABLE `sys_member`
(
    `id`                   bigint                                                        NOT NULL AUTO_INCREMENT,
    `username`             varchar(20)                                                   NOT NULL DEFAULT '' COMMENT '帐号',
    `password`             char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci     NOT NULL DEFAULT '' COMMENT '密码',
    `salt`                 char(16)                                                      NOT NULL COMMENT '密码盐',
    `auth_key`             char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci     NOT NULL DEFAULT '' COMMENT '授权令牌',
    `password_reset_token` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码重置令牌',
    `type`                 tinyint                                                       NOT NULL DEFAULT '1' COMMENT '账户类型：1-普通管理员，10-超级管理员',
    `nick_name`            varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar`               varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
    `sex`                  tinyint                                                       NOT NULL DEFAULT '0' COMMENT '性别：0-未知，1-男，2-女',
    `email`                varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '' COMMENT '邮箱',
    `birthday`             datetime                                                      NOT NULL COMMENT '生日',
    `province_id`          int                                                           NOT NULL DEFAULT '0' COMMENT '省',
    `city_id`              int                                                           NOT NULL DEFAULT '0' COMMENT '城市',
    `area_id`              int                                                           NOT NULL DEFAULT '0' COMMENT '地区',
    `address`              varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认地址',
    `mobile`               varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '' COMMENT '手机号码',
    `visit_count`          int unsigned NOT NULL DEFAULT '0' COMMENT '访问次数',
    `last_at`              datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后登录时间',
    `last_ip`              varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL DEFAULT '' COMMENT '最后一次登录ip',
    `status`               tinyint                                                       NOT NULL DEFAULT '1' COMMENT '状态:10-正常，20-关闭，30-待验证',
    `created_at`           datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`           datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `salt` (`salt`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统后台管理员表';

CREATE TABLE `sys_member_role`
(
    `member_id` bigint NOT NULL COMMENT '用户ID',
    `role_id`   bigint unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
    PRIMARY KEY (`member_id`, `role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户和角色关联表';

CREATE TABLE `sys_menu`
(
    `id`         bigint                                                  NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
    `pid`        bigint unsigned NOT NULL DEFAULT '0' COMMENT '父菜单ID',
    `name`       varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '菜单名称',
    `code`       varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '菜单编码',
    `icon`       varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '菜单图标',
    `type`       char(1) CHARACTER SET utf8 COLLATE utf8_general_ci      NOT NULL DEFAULT '' COMMENT '菜单类型（M目录 C菜单 F按钮）',
    `perms`      varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
    `path`       varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
    `component`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
    `query`      varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '路由参数',
    `is_cache`   tinyint unsigned NOT NULL DEFAULT '20' COMMENT '是否缓存：10-是，20-否',
    `is_visible` tinyint unsigned NOT NULL DEFAULT '20' COMMENT '是否隐藏：10-是，20-否',
    `level`      int unsigned NOT NULL DEFAULT '1' COMMENT '级别',
    `tree`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '树',
    `sort`       int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
    `status`     tinyint unsigned NOT NULL DEFAULT '10' COMMENT '菜单状态：10-开启，20-关闭',
    `created_at` datetime                                                NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime                                                NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY          `pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='菜单权限表';


CREATE TABLE `sys_role`
(
    `id`                  bigint                                                  NOT NULL AUTO_INCREMENT COMMENT '角色ID',
    `name`                varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL DEFAULT '' COMMENT '角色名称',
    `key`                 varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '角色权限字符串',
    `menu_check_strictly` tinyint unsigned NOT NULL DEFAULT '10' COMMENT '菜单树选择项是否关联显示：10-显示，20-隐藏',
    `sort`                int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
    `status`              tinyint unsigned NOT NULL DEFAULT '10' COMMENT '角色状态：10-开启，20-关闭',
    `created_at`          datetime                                                NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`          datetime                                                NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色信息表';


CREATE TABLE `sys_role_menu`
(
    `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
    `menu_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
    PRIMARY KEY (`role_id`, `menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色和菜单关联表';

