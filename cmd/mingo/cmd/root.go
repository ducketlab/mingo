package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var vers bool

var RootCmd = &cobra.Command{
	Use:   "mingo",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no flags find")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(
		&vers, "version", "v", false, "the mingo version")
}