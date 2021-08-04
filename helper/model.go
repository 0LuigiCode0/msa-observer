package helper

//Config модель конфига
type Config struct {
	Host     string                    `json:"host"`
	Port     int32                     `json:"port"`
	DBS      map[string]*DbConfig      `json:"dbs"`
	Handlers map[string]*HandlerConfig `json:"handlers"`
	Admins   []*Admin                  `json:"admins"`
}

type DbConfig struct {
	Type     DBType `json:"type"`
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type HandlerConfig struct {
	Type     HandlerType `json:"type"`
	Host     string      `json:"host"`
	Port     int32       `json:"port"`
	Key      string      `json:"key"`
	User     string      `json:"user"`
	Password string      `json:"password"`
	IsTSL    bool        `json:"is_tsl"`
}

type Admin struct {
	Login string `json:"login"`
	Pwd   string `json:"pwd"`
}

//ErrLocal модель локализации ошибок
type ErrLocal struct {
	TitleEn string `bson:"title_en" json:"title_en"`
	TitleRu string `bson:"title_ru" json:"title_ru"`
}

//Модель для отправки по ws
type WsType string

const ()

type SendModel struct {
	Type WsType      `json:"type"`
	Data interface{} `json:"data"`
}

type DBType string

const (
	Postgres DBType = "postgres"
	Mongodb  DBType = "mongodb"
)

type HandlerType string

const (
	TCP  HandlerType = "tcp"
	MQTT HandlerType = "mqtt"
	WS   HandlerType = "ws"
	GRPC HandlerType = "grpc"
)

//Основные константы
const (
	UploadDir  = "./source/uploads/"
	ConfigDir  = "./source/configs/"
	ConfigFile = "configServer.json"
	Secret     = "qfdQjmVLiW"
)

//Названеия коллекции

type Collection string

const (
	CollUsers    Collection = "users"
	CollServices Collection = "services"
	CollPeers    Collection = "peers"
	CollGroups   Collection = "groups"
)

//Ключи контекста

type CtxKey int

const (
	CtxKeyValue CtxKey = iota
)

//Роли пользователей

type Role byte

const (
	RoleSuperAdmin Role = iota + 1
	RoleAdmin
	RoleUser
)

//Статусы

//Значения сортировки

//Типы полей
