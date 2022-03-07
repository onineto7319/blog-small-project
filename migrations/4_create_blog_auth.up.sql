CREATE TABLE IF NOT EXISTS `blog_auth` (
		`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
		`app_key` varchar(20) DEFAULT '' COMMENT 'Key',
		`app_secret` varchar(50) DEFAULT '' COMMENT 'Sccret',
		`created_on` int(10) unsigned DEFAULT '0' COMMENT '建立時間',
		`created_by` varchar(100) DEFAULT '' COMMENT '建立人',
		`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
		`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
		`deleted_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
		`is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除0為未刪除、1為已刪除',
		PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='認證管理';