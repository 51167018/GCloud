package models

import "gorm.io/gorm"

type File struct {
	/*
	 `my_file_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '文件ID',
	  `my_file_name` varchar(255) DEFAULT NULL COMMENT '文件名',
	  `file_store_id` int(11) DEFAULT NULL COMMENT '文件仓库ID',
	  `my_file_path` varchar(255) DEFAULT '/' COMMENT '文件存储路径',
	  `download_time` int(11) DEFAULT '0' COMMENT '下载次数',
	  `upload_time` datetime DEFAULT NULL COMMENT '上传时间',
	  `parent_folder_id` int(11) DEFAULT NULL COMMENT '父文件夹ID',
	  `size` int(11) DEFAULT NULL COMMENT '文件大小',
	  `type` int(11) DEFAULT NULL COMMENT '文件类型',
	  `postfix` varchar(255) DEFAULT NULL COMMENT '文件后缀',
	*/
	gorm.Model
	FileName string //文件名称
	UserId int //文件所属用户的ID
	RealPath string //文件所在的真实路径
	FileSize int //文件大小
	FileHash string //文件哈希值
	FileType int //文件类型
	FolderId int //所属文件夹的Id
}
