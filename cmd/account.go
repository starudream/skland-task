package main

import (
	"fmt"

	"github.com/starudream/go-lib/cobra/v2"

	"github.com/starudream/skland-task/api/hypergryph"
	"github.com/starudream/skland-task/api/skland"
	"github.com/starudream/skland-task/config"
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
			phone := util.Scan("please enter phone: ")
			if phone == "" {
				return fmt.Errorf("phone is empty")
			}

			err := hypergryph.SendPhoneCode(phone)
			if err != nil {
				return fmt.Errorf("send phone code error: %w", err)
			}

			code := util.Scan("please enter the verification code you received: ")
			if code == "" {
				return fmt.Errorf("verification code is empty")
			}

			res1, err := hypergryph.LoginByPhoneCode(phone, code)
			if err != nil {
				return fmt.Errorf("login by phone code error: %w", err)
			}

			res2, err := hypergryph.GrantApp(res1.Token, hypergryph.AppCodeSKLAND)
			if err != nil {
				return fmt.Errorf("grant app error: %w", err)
			}

			res3, err := skland.AuthLoginByCode(res2.Code)
			if err != nil {
				return fmt.Errorf("auth login by code error: %w", err)
			}

			config.AddAccount(config.Account{
				Phone: phone,
				Hypergryph: config.AccountHypergryph{
					Token: res1.Token,
					Code:  res2.Code,
				},
				Skland: config.AccountSkland{
					Token: res3.Token,
					Cred:  res3.Cred,
				},
			})
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
