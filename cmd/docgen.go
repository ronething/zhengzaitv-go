package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docgenCmd represents the docgen command
var docgenCmd = &cobra.Command{
	Use:   "docgen",
	Short: "文档生成",
	Long:  `文档生成`,
	Run: func(cmd *cobra.Command, args []string) {
		doc.GenMarkdownTree(ticketsCmd, "./")
	},
}

func init() {
	rootCmd.AddCommand(docgenCmd)
}
