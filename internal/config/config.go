package config

type RecurringTaskConfig struct {
	Tasks []RecurringTask `json:"tasks"`
}

type RecurringTask struct {
	Name     string `json:"name"`
	Schedule string `json:"schedule"` // cron format
	ParentID string `json:"parent_id"`

	Title  string `json:"title"`
	Status Status `json:"status"`
}

type Status struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}
