package tools

import (
	"fmt"

	"github.com/spf13/cobra"
)

func LoadGlobalParams(cmd *cobra.Command, req func(charger, username, password string)) error {
	charger, _ := cmd.Root().Flags().GetString("charger")

	user, _ := cmd.Root().Flags().GetString("username")
	pass, _ := cmd.Root().Flags().GetString("password")

	if len(charger) == 0 {
		return fmt.Errorf("could not find warp charger address")
	}

	req(charger, user, pass)

	return nil
}
