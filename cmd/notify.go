package main

import (
	"context"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/ntfy/v2"
)

var (
	notifyCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "notify"
		c.Short = "Manage notify"
	})

	notifySendCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "send <message>"
		c.Short = "Send notify"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 || args[0] == "" {
				args = []string{"Hello World!"}
			}
			return ntfy.Notify(context.Background(), args[0])
		}
	})
)

func init() {
	notifyCmd.AddCommand(notifySendCmd)

	rootCmd.AddCommand(notifyCmd)
}
