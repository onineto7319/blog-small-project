CREATE TABLE IF NOT EXISTS `blog_article_tag` (
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`article_id` int(10) NOT NULL COMMENT '文章ID',
	`tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '標籤ID', 
	`created_on` int(10) unsigned DEFAULT '0' COMMENT '建立時間',
    `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
    `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
    `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
    `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章標籤連結';