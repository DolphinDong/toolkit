package config

type Config struct {
	TaskSchedule map[string]string `json:"task_schedule" validate:"required"`
}
