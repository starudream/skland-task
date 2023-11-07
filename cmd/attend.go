package main

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/slog"

	"github.com/starudream/skland-task/config"
	"github.com/starudream/skland-task/job"
)

var attend = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "attend <account phone>"
	c.Short = "Attend skland"
	c.Args = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || args[0] == "" {
			return fmt.Errorf("requires account phone")
		}
		_, exists := config.GetAccount(args[0])
		if !exists {
			return fmt.Errorf("account %s not exists", args[0])
		}
		return nil
	}
	c.RunE = func(cmd *cobra.Command, args []string) error {
		account, _ := config.GetAccount(args[0])
		awards, err := job.Attend(account)
		if err != nil {
			return err
		}
		slog.Info("attend awards:\n%s", job.FormatAwards(awards))
		return nil
	}
})

func init() {
	rootCmd.AddCommand(attend)
}
