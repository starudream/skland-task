package main

import (
	"context"
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/cron/v2"
	"github.com/starudream/go-lib/ntfy/v2"

	"github.com/starudream/skland-task/config"
	"github.com/starudream/skland-task/job"
)

var cronCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "cron"
	c.Short = "Run as cron job"
	c.RunE = func(cmd *cobra.Command, args []string) error {
		if config.C().Cron.Startup {
			cronRun()
		}
		err := cron.AddJob(config.C().Cron.Spec, "skland-cron", cronRun)
		if err != nil {
			return fmt.Errorf("add cron job error: %w", err)
		}
		cron.Run()
		return nil
	}
})

func init() {
	rootCmd.AddCommand(cronCmd)
}

func cronRun() {
	c := config.C()
	for i := 0; i < len(c.Accounts); i++ {
		cronCheckinAccount(config.C().Accounts[i])
		cronPostAccount(config.C().Accounts[i])
		cronAttendAccount(config.C().Accounts[i])
	}
}

func cronCheckinAccount(account config.Account) (msg string) {
	data, err := job.Checkin(account)
	if err != nil {
		msg = fmt.Sprintf("森空岛版区签到失败: %v", err)
		slog.Error(msg)
	} else {
		msg = account.Phone + "\n" + job.FormatCheckin(data)
		slog.Info(msg)
	}
	err = ntfy.Notify(context.Background(), msg)
	if err != nil {
		slog.Error("cron skland notify error: %v", err)
	}
	return
}

func cronPostAccount(account config.Account) (msg string) {
	err := job.Post(account)
	if err != nil {
		msg = fmt.Sprintf("森空岛版区等级任务失败: %v", err)
		slog.Error(msg)
	} else {
		msg = account.Phone + "\n" + "森空岛版区等级任务成功"
		slog.Info(msg)
	}
	err = ntfy.Notify(context.Background(), msg)
	if err != nil {
		slog.Error("cron skland notify error: %v", err)
	}
	return
}

func cronAttendAccount(account config.Account) (msg string) {
	awards, err := job.Attend(account)
	if err != nil {
		msg = fmt.Sprintf("森空岛福利签到失败: %v", err)
		slog.Error(msg)
	} else {
		msg = account.Phone + " " + job.FormatAwards(awards)
		slog.Info(msg)
	}
	err = ntfy.Notify(context.Background(), msg)
	if err != nil {
		slog.Error("cron skland notify error: %v", err)
	}
	return
}
