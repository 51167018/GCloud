package models

import "gorm.io/gorm"

type Folder struct {
	/*
	 `file_folder_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文件夹ID',
	  `file_folder_name` varchar(255) DEFAULT NULL COMMENT '文件夹名称',
	  `parent_folder_id` int(11) DEFAULT '0' COMMENT '父文件夹ID',
	  `file_store_id` int(11) DEFAULT NULL COMMENT '所属文件仓库ID',
	  `time` datetime DEFAULT NULL COMMENT '创建时间',
	*/
	gorm.Model
	FolderName string //文件夹名
	ParentFolderId int //父文件夹ID
	Files []File//该文件夹下的文件
	UserId int //文件夹所属用户的Id
}
