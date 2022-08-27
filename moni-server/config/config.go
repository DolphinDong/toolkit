package config

type Server struct {
	TaskSchedule map[string]string `json:"task_schedule" validate:"required"`
	ServerHost   string            `json:"server_host" validate:"required"`
	ServerPort   int               `json:"server_port" validate:"required"`
	Log          Log               `json:"log" validate:"required"`
	Mysql        Mysql             `json:"mysql" validate:"required"`
	Redis        Redis             `json:"redis" validate:"required"`
	Token        string            `json:"token" validate:"required"`
}

// Mysql配置
type Mysql struct {
	MysqlHost     string `json:"mysql_host" validate:"required"`
	MysqlPort     int    `json:"mysql_port" validate:"required"`
	MysqlUser     string `json:"mysql_user" validate:"required"`
	MysqlPassword string `json:"mysql_password" validate:"required"`
	DBName        string `json:"db_name" validate:"required"`
	MaxIdleConns  int    `json:"max_idle_conns" validate:"required"`
	MaxOpenConns  int    `json:"max_open_conns" validate:"required"`
}

type Redis struct {
	RedisHost      string `json:"redis_host" validate:"required"`
	RedisPort      int    `json:"redis_port" validate:"required"`
	RedisPassword  string `json:"redis_password" validate:"required"`
	PoolMaxIdle    int    `json:"pool_max_idle" validate:"required"`
	PoolMaxActive  int    `json:"pool_max_active" validate:"required"`
	RedisDB        int    `json:"redis_db" validate:"required"`
	ConnectTimeOut int    `json:"connect_time_out" validate:"required"`
	IdleTimeout    int    `json:"idle_timeout" validate:"required"`
}
type Log struct {
	LogLevel string `json:"log_level" validate:"required"`
	LogPath  string `json:"log_path" validate:"required"`
}
