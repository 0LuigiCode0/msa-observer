package helper

import (
	"context"
	"os"
	"sync"
)

var Ctx context.Context
var CloseCtx context.CancelFunc
var Wg sync.WaitGroup
var C chan os.Signal

type GroupsType string

const (
	Observer    GroupsType = "observer"
	Monitor     GroupsType = "monitor"
	WS          GroupsType = "ws"
	Bot         GroupsType = "bot"
	Support     GroupsType = "support"
	Valid       GroupsType = "valid"
	Auth        GroupsType = "auth"
	User        GroupsType = "user"
	Recrut      GroupsType = "recrut"
	Event       GroupsType = "event"
	Office      GroupsType = "office"
	Product     GroupsType = "product"
	News        GroupsType = "news"
	Knows       GroupsType = "knows"
	Filling     GroupsType = "filling"
	Group       GroupsType = "group"
	Quota       GroupsType = "quota"
	Cert        GroupsType = "cert"
	Proposal    GroupsType = "proposal"
	Comment     GroupsType = "comment"
	Notif       GroupsType = "notif"
	Transatcion GroupsType = "transatcion"
	Skillpoint  GroupsType = "skillpoint"
	Analitic    GroupsType = "analitic"
	Uploads     GroupsType = "uploads"
	Phone       GroupsType = "phone"
	Email       GroupsType = "email"
	Xtech       GroupsType = "xtech"
)

//Коды ошибок

type ErrCode byte

const (
	ErrorNotFound ErrCode = iota
	ErrorExist
	ErrorSave
	ErrorUpdate
	ErrorDelete
	ErorrAccessDeniedToken
	ErorrAccessDeniedParams
	ErrorInvalidParams
	ErrorParse
	ErrorOpen
	ErrorClose
	ErrorWrite
	ErrorRead
	ErrorGenerate
	ErrorSend
)
const (
	KeyErrorNotFound      = "not found"
	KeyErrorExist         = "already exists"
	KeyErrorSave          = "create is failed"
	KeyErrorUpdate        = "update is failed"
	KeyErrorDelete        = "delete is failed"
	KeyErorrAccessDenied  = "access denied"
	KeyErrorInvalidParams = "invalid params"
	KeyErrorParse         = "parse is failed"
	KeyErrorOpen          = "open is failed"
	KeyErrorClose         = "close is failed"
	KeyErrorRead          = "read is failed"
	KeyErrorWrite         = "write is faleid"
	KeyErrorGenerate      = "generate is falied"
	KeyErrorSend          = "send is faied"
)

//ResponseError модель ошибки
type ResponseError struct {
	Code ErrCode `json:"code"`
	Msg  string  `json:"msg"`
}

//ResponseModel модель ответа
type ResponseModel struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}
