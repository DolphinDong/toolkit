package common

const (
	LoggerKey             = "logger"
	RequestIDKey          = "requestID"
	Salt                  = "4652e015-e0d7-4aa5-ae7a-f99195796b25"
	SessionKey            = "userID"      // session 中用户信息的key
	ContextUserIdKey      = "currentUser" // context 中存放用户的key
	LoadData              = 1
	AddData               = 2
	DeleteData            = 3
	UpdateData            = 4
	EnableData            = 5
	LayuiTabelDataOk      = 0
	LayuiTabelDataError   = 1
	InvalidQueryParameter = "Invalid query parameter"
	CommonRole            = "common"
	QueryStr              = "%%%v%%"
)
