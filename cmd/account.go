package main

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"

	"github.com/starudream/skland-task/api/hypergryph"
	"github.com/starudream/skland-task/config"
	"github.com/starudream/skland-task/job"
	"github.com/starudream/skland-task/util"
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
			phone := util.Scan("please enter phone (use ctrl+c to exit): ")
			if phone == "" {
				return nil
			}

			err := hypergryph.SendPhoneCode(phone)
			if err != nil {
				return fmt.Errorf("send phone code error: %w", err)
			}

			code := util.Scan("please enter the verification code you received (use ctrl+c to exit): ")
			if code == "" {
				return nil
			}

			res1, err := hypergryph.LoginByPhoneCode(phone, code)
			if err != nil {
				return fmt.Errorf("login by phone code error: %w", err)
			}

			account, err := job.Login(res1.Token)
			if err != nil {
				return err
			}

			account.Phone = phone
			account.Hypergryph.Token = res1.Token

			config.AddAccount(account)
			return config.Save()
		}
	})

	accountListCmd = cobra.NewCommand(func(c *cobra.Command) {
		c.Use = "list"
		c.Short = "List accounts"
		c.Run = func(cmd *cobra.Command, args []string) {
			accounts := config.C().Accounts
			lines := make([]string, len(accounts)+1)
			lines[0] = "phone\thg-token\thg-code\tsk-token\tsk-cred"
			for i, a := range accounts {
				fmt.Printf("%#v\n", a)
				lines[i+1] = fmt.Sprintf("%s\t%s\t%s\t%s\t%s",
					a.Phone,
					a.Hypergryph.Token, a.Hypergryph.Code,
					a.Skland.Token, a.Skland.Cred,
				)
			}
			fmt.Println(util.TabWriter(lines...))
		}
	})
)

func init() {
	accountCmd.AddCommand(accountLoginCmd)
	accountCmd.AddCommand(accountListCmd)

	rootCmd.AddCommand(accountCmd)
}
