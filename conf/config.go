package conf

import "sync"

type DbConfig struct {
	DriverName string
	Dsn string
	MaxLifetime int
	MaxOpen int
	MaxId int
}

type Config struct {
	Language string
	Token string
	Super string
	RedisPre string
	Host string
	OpenJwt bool
	Routes []string
}
var (
	Cfg  Config
	mutex   sync.Mutex
	declare sync.Once
)

//parseTime=true 更改输出类型 DATE 和 DATETIME 值为 time.Time;
//loc=true使用本地时间;
var Db = map[string]DbConfig {
	"db1":{
		DriverName: "mysql",
		Dsn:        "root:123456@tcp(127.0.0.1)/test?charset=utf8mb4&parseTime=true&loc=Local",
        // MaxLifetime: 3,
		// MaxId:       10,
		// MaxOpen:     200,
	},
}

func  Set(cfg Config) {
	mutex.Lock()
	Cfg.RedisPre=setDefault(cfg.RedisPre,"","go.sso.redis")
	Cfg.Language=setDefault(cfg.Language,"","cn")
	Cfg.Token=setDefault(cfg.Token,"","token")
	Cfg.Super=setDefault(cfg.Super,"","admin")//超级账户
	Cfg.Host=setDefault(cfg.Host,"","http://localhost:8282")//域名
	Cfg.Routes=cfg.Routes
	Cfg.OpenJwt=cfg.OpenJwt
	mutex.Unlock()
}
func setDefault( value,def ,defValue string) string {
	if value==def {
		return defValue
	}
	return value
}

