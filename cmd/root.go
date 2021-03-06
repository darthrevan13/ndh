package cmd

import (
	"fmt"
	"os"

	"github.com/darthrevan13/ndh/pkg/pkgManager"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ndh",
	Short: "Get hierarchical dependencies for node packages",
	Long:  "Given a node package name show hierarchical/tree dependencies of the package",
	RunE: func(_ *cobra.Command, args []string) error {
		pkgName := args[0]
		pkgVer := args[1]
		dep, err := pkgManager.GetAllDependencies(pkgName, pkgVer)
		if err != nil {
			return err
		}
		tree, err := dep.ToJson()
		if err != nil {
			return err
		}
		fmt.Println(tree)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	//TODO: Implement loading .env config
}
