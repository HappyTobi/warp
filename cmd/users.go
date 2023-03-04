package cmd

import (
	"github.com/HappyTobi/warp/pkg/cmd/users"
	"github.com/spf13/cobra"
)

func UserCmd() *cobra.Command {
	users := &cobra.Command{
		Use:   "users",
		Short: "Users command",
	}
	users.AddCommand(ListCmd())

	return users
}

func ListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all users command",
		RunE:  users.List,
	}
}
