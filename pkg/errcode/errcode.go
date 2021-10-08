package errcode

const (
	SUCCESS = 200
	ERROR   = 500

	//用户
	ERROR_USERNAME_EXIST = 1000
	ERROR_PASSWORD_ERROR = 1001

	//文件
	UPLOAD_FILE_SUCCESS      = 2000
	UPLOAD_FILE_ERROR_CLIENT = 2001
	UPLOAD_FILE_ERROR        = 2002
	UPLOAD_ERROR_FILE_EXIST  = 2003

	//数据库
	INSERT_DATABASE_ERROR = 3000
	SELECT_DATABASE_ERROR = 3001
	UPDATE_DATABASE_ERROR = 3002
	DELETE_DATABASE_ERROR = 3003

	//文件夹
	FOLDER_NOT_FOUND = 4001
)

var errCode = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "FAILURE",
	ERROR_USERNAME_EXIST: "用户名已经存在",
	ERROR_PASSWORD_ERROR: "密码错误或账户不存在",

	UPLOAD_FILE_ERROR_CLIENT: "从客户端接收文件失败",
	UPLOAD_FILE_SUCCESS:      "文件上传成功",
	UPLOAD_FILE_ERROR:        "文件不存在,可以上传",
	UPLOAD_ERROR_FILE_EXIST:  "文件上传失败,文件已存在",
	INSERT_DATABASE_ERROR:    "插入数据库失败",
	SELECT_DATABASE_ERROR:    "查询数据库失败",
	UPDATE_DATABASE_ERROR:    "更新数据库失败",
	DELETE_DATABASE_ERROR:    "删除数据库失败",

	FOLDER_NOT_FOUND: "文件夹未找到",
}

func GetErrMessage(code int) string {
	return errCode[code]
}
