package main

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/utils/sliceutil"

	"github.com/starudream/skland-task/config"
	"github.com/starudream/skland-task/job"
)

var attendCmd = cobra.NewCommand(func(c *cobra.Command) {
	c.Use = "attend <account phone>"
	c.Short = "Attend skland"
	c.Args = func(cmd *cobra.Command, args []string) error {
		phone, _ := sliceutil.GetValue(args, 0)
		if phone == "" {
			return fmt.Errorf("requires account phone")
		}
		_, exists := config.GetAccount(phone)
		if !exists {
			return fmt.Errorf("account %s not exists", phone)
		}
		return nil
	}
	c.RunE = func(cmd *cobra.Command, args []string) error {
		phone, _ := sliceutil.GetValue(args, 0)
		account, _ := config.GetAccount(phone)
		_, err := job.Attend(account)
		return err
	}
})

func init() {
	rootCmd.AddCommand(attendCmd)
}
