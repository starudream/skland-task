package main

import (
	"github.com/starudream/go-lib/cobra/v2"
)

var rootCmd = cobra.NewRootCommand(func(c *cobra.Command) {
	c.Use = "skland-task"

	cobra.AddConfigFlag(c)
})
