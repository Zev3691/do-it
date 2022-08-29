package errorx

type BaseErrMsg struct {
	ErrMsg   string `json:"err_msg,omitempty"`
	ErrMsgEn string `json:"err_msg_en,omitempty"`
}

type ErrCode int

var (
	baseErrors = map[ErrCode]BaseErrMsg{}
)

const (
	ErrSuccess                 ErrCode = 0
	ErrUnknownError            ErrCode = 1
	ErrUnstableNetwork         ErrCode = 2
	ErrPermissionDeny          ErrCode = 3
	ErrServiceUnderMaintaining ErrCode = 4
	ErrTooMuchRequest          ErrCode = 5
	ErrServiceNotFound         ErrCode = 6

	ErrInvalidParameter ErrCode = 400009

	ErrFrequencyControl ErrCode = 400010
	ErrBlackList        ErrCode = 400011
	ErrMethodNotSupport ErrCode = 400012

	ErrJsonFormat       ErrCode = 400013
	ErrOperationWorking ErrCode = 400014
	ErrJsonParse        ErrCode = 400015
	ErrStackOverflow    ErrCode = 400016
	ErrAccessFailed     ErrCode = 400017

	ErrDBOptFailed ErrCode = 400051
)

func init() {
	// err code definition rule:
	// 	0				System Reserved: Success
	//  400001 - 400099			System Reserved: Basic Error

	//  50000 -   	Business Logic

	baseErrors[ErrSuccess] = BaseErrMsg{ErrMsg: "成功", ErrMsgEn: "success"}

	baseErrors[ErrUnknownError] = BaseErrMsg{ErrMsg: "未知错误 %s", ErrMsgEn: "unknown error"}
	baseErrors[ErrUnstableNetwork] = BaseErrMsg{ErrMsg: "网络波动, 请重新尝试", ErrMsgEn: "unstable network connection, please retry"}
	baseErrors[ErrPermissionDeny] = BaseErrMsg{ErrMsg: "权限不足", ErrMsgEn: "permission denied"}
	baseErrors[ErrServiceUnderMaintaining] = BaseErrMsg{ErrMsg: "服务维护中, 请稍后", ErrMsgEn: "service maintaining, please wait"}
	baseErrors[ErrTooMuchRequest] = BaseErrMsg{ErrMsg: "访问量过大, 请稍后重试", ErrMsgEn: "router under flow control"}
	baseErrors[ErrServiceNotFound] = BaseErrMsg{ErrMsg: "请求的服务不存在", ErrMsgEn: "service not found"}
	baseErrors[400007] = BaseErrMsg{ErrMsg: "缺少参数: %s", ErrMsgEn: "lack of parameter"}
	baseErrors[400008] = BaseErrMsg{ErrMsg: "缺少参数", ErrMsgEn: "lack of parameter"}
	baseErrors[ErrInvalidParameter] = BaseErrMsg{ErrMsg: "%s", ErrMsgEn: "invalid parameter"}
	baseErrors[ErrFrequencyControl] = BaseErrMsg{ErrMsg: "您的请求过于频繁, 请稍后再试", ErrMsgEn: "request frequency control"}
	baseErrors[ErrBlackList] = BaseErrMsg{ErrMsg: "访问受限: %s", ErrMsgEn: "request frequency control [black list]"}
	baseErrors[ErrMethodNotSupport] = BaseErrMsg{ErrMsg: "不支持的方法", ErrMsgEn: "method not support"}
	baseErrors[ErrJsonFormat] = BaseErrMsg{ErrMsg: "JSON参数错误 %s", ErrMsgEn: "json param format error"}
	baseErrors[ErrOperationWorking] = BaseErrMsg{ErrMsg: "操作正在执行中，请稍后重试", ErrMsgEn: "operation is working, waiting"}
	baseErrors[ErrJsonParse] = BaseErrMsg{ErrMsg: "JSON解析出错 %s", ErrMsgEn: "json parse error"}
	baseErrors[ErrStackOverflow] = BaseErrMsg{ErrMsg: "函数栈溢出", ErrMsgEn: "func stack overflow"}
	baseErrors[ErrAccessFailed] = BaseErrMsg{ErrMsg: "用户认证失败", ErrMsgEn: "user authentication failed"}

	baseErrors[400020] = BaseErrMsg{ErrMsg: "类型错误: %s", ErrMsgEn: "invalid type"}
	baseErrors[400030] = BaseErrMsg{ErrMsg: "'%s'长度不得大于%d", ErrMsgEn: "field too long"}

	baseErrors[400040] = BaseErrMsg{ErrMsg: "未定义的环境变量: %s", ErrMsgEn: "undefined environment variables"}
	baseErrors[400041] = BaseErrMsg{ErrMsg: "无法识别的环境变量: %s=%s", ErrMsgEn: "undefined environment variables"}

	baseErrors[400050] = BaseErrMsg{ErrMsg: "数据层类型错误: %s", ErrMsgEn: "failed when transform type on orm"}
	baseErrors[ErrDBOptFailed] = BaseErrMsg{ErrMsg: "数据库操作失败: %s", ErrMsgEn: "db op error"}
	baseErrors[400052] = BaseErrMsg{ErrMsg: "无法获取数据库连接: %s", ErrMsgEn: "Failed to get db-connection"}
}
