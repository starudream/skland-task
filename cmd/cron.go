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
		cfg := config.C().CronAttend
		if cfg.Startup {
			cronAttend()
		}
		err := cron.AddJob(cfg.Spec, "skland-attend", cronAttend)
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

func cronAttend() {
	accounts := config.C().Accounts
	for i := 0; i < len(accounts); i++ {
		cronAttendAccount(accounts[i])
	}
}

func cronAttendAccount(account config.Account) (msg string) {
	awards, err := job.Attend(account)
	if err != nil {
		msg = fmt.Sprintf("cron skland attend error: %v", err)
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
