package main

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"
	"github.com/starudream/go-lib/core/v2/utils/fmtutil"
	"github.com/starudream/go-lib/tablew/v2"

	"github.com/starudream/skland-task/api/hypergryph"
	"github.com/starudream/skland-task/config"
	"github.com/starudream/skland-task/job"
)

var (
	accountCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "account"
		c.Short = "Manage accounts"
	})

	accountLoginCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "login"
		c.Short = "Login account"
		c.RunE = func(cmd *cobra.Command, args []string) error {
			phone := fmtutil.Scan("please enter phone (use ctrl+c to exit): ")
			if phone == "" {
				return nil
			}

			err := hypergryph.SendPhoneCode(phone)
			if err != nil {
				return fmt.Errorf("send phone code error: %w", err)
			}

			code := fmtutil.Scan("please enter the verification code you received (use ctrl+c to exit): ")
			if code == "" {
				return nil
			}

			res, err := hypergryph.LoginByPhoneCode(phone, code)
			if err != nil {
				return fmt.Errorf("login by phone code error: %w", err)
			}

			account := config.Account{Phone: phone, Hypergryph: config.AccountHypergryph{Token: res.Token}}

			account, err = job.Login(account)
			if err != nil {
				return err
			}

			config.AddAccount(account)
			return config.Save()
		}
	})

	accountListCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "list"
		c.Short = "List accounts"
		c.Run = func(cmd *cobra.Command, args []string) {
			fmt.Println(tablew.Structs(config.C().Accounts))
		}
	})
)

func init() {
	accountCmd.AddCommand(accountLoginCmd)
	accountCmd.AddCommand(accountListCmd)

	rootCmd.AddCommand(accountCmd)
}
