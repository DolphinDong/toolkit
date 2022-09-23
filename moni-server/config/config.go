package config

type Server struct {
	TaskSchedule  map[string]string `json:"task_schedule" validate:"required"`
	ServerHost    string            `json:"server_host" validate:"required,ip4_addr"`
	ServerPort    int               `json:"server_port" validate:"required,gt=0"`
	Log           Log               `json:"log" validate:"required"`
	Mysql         Mysql             `json:"mysql" validate:"required"`
	Redis         Redis             `json:"redis" validate:"required"`
	Token         string            `json:"token" validate:"required"`
	RobotKey      string            `json:"robot_key" validate:"required"`
	WechatBaseUrl string            `json:"wechat_base_url" validate:"required"`
}

// Mysql配置
type Mysql struct {
	MysqlHost     string `json:"mysql_host" validate:"required,ip4_addr|hostname"`
	MysqlPort     int    `json:"mysql_port" validate:"required,gt=0"`
	MysqlUser     string `json:"mysql_user" validate:"required"`
	MysqlPassword string `json:"mysql_password" validate:"required"`
	DBName        string `json:"db_name" validate:"required"`
	MaxIdleConns  int    `json:"max_idle_conns" validate:"required,gt=0"`
	MaxOpenConns  int    `json:"max_open_conns" validate:"required,gt=0"`
}

type Redis struct {
	RedisHost      string `json:"redis_host" validate:"required,ip4_addr|hostname"`
	RedisPort      int    `json:"redis_port" validate:"required,gt=0"`
	RedisPassword  string `json:"redis_password" validate:"required"`
	PoolMaxIdle    int    `json:"pool_max_idle" validate:"required,gt=0"`
	PoolMaxActive  int    `json:"pool_max_active" validate:"required,gt=0"`
	RedisDB        int    `json:"redis_db" validate:"required,lte=15,gte=0"`
	ConnectTimeOut int    `json:"connect_time_out" validate:"required,gt=0"`
	IdleTimeout    int    `json:"idle_timeout" validate:"required,gt=0"`
}
type Log struct {
	LogLevel string `json:"log_level" validate:"required"`
	LogPath  string `json:"log_path" validate:"required"`
}
