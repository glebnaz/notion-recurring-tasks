package service

import (
	"encoding/json"
	"github.com/glebnaz/notion-recurring-tasks/internal/config"
	"github.com/glebnaz/notion-recurring-tasks/internal/notion"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
)

type RecurringTaskController struct {
	Config config.RecurringTaskConfig

	token string

	cron *cron.Cron

	notion notion.Controller
}

func NewConfigFromFile(path string) (config.RecurringTaskConfig, error) {
	var config config.RecurringTaskConfig
	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil

}

func NewRecurringTaskController(config config.RecurringTaskConfig, token string) *RecurringTaskController {
	cronCtrl := cron.New()
	cliNotion := notion.NewController(token)
	return &RecurringTaskController{
		Config: config,
		token:  token,
		cron:   cronCtrl,
		notion: cliNotion,
	}
}

func (r *RecurringTaskController) Start(ctx context.Context) {
	r.cron.Start()
}

func (r *RecurringTaskController) RegisterConfig(ctx context.Context) error {
	for _, task := range r.Config.Tasks {
		err := r.registerTask(ctx, task)
		if err != nil {
			return err
		}
	}
	log.Infof("Registered %d tasks", len(r.Config.Tasks))
	return nil
}

func (r *RecurringTaskController) registerTask(ctx context.Context, task config.RecurringTask) error {
	err := r.cron.AddFunc(task.Schedule, r.runTask(ctx, task))
	if err != nil {
		log.Errorf("Error adding task %s: %s", task.Name, err)
		return err
	}
	log.Infof("Registered task %s", task.Name)
	return nil
}

func (r *RecurringTaskController) runTask(ctx context.Context, task config.RecurringTask) func() {
	return func() {
		log.Infof("Run task %s", task.Name)
		err := r.notion.AddNewPageToDataBase(ctx, task)
		if err != nil {
			log.Errorf("Error run task %s: %s", task.Name, err)
			return
		}
		log.Infof("Task %s done", task.Name)
	}
}
