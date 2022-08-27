package config

type Config struct {
	TaskSchedule    map[string]string `json:"task_schedule" validate:"required"`
	MoniServerHost  string            `json:"moni_server_host" validate:"required,url"`
	MoniServerToken string            `json:"moni_server_token" validate:"required"`
}
