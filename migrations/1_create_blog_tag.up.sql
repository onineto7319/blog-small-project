CREATE TABLE IF NOT EXISTS `blog_tag` (
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(100) DEFAULT '' COMMENT '標籤名稱', 
	`state` tinyint(3) unsigned DEFAULT '1' COMMENT '狀態 0 為禁用、1為啟用',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '建立時間',
    `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
    `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除0為未刪除、1為已刪除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='標籤管理';