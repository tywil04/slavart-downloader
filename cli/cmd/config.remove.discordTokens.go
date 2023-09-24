package cmd

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tywil04/slavartdl/cli/internal/config"
)

var configRemoveDiscordTokensCmd = &cobra.Command{
	Use:          "discordTokens [flags] tokenIndex(s)",
	Short:        "Removes discord session token(s) using index shown by the list command",
	SilenceUsage: true,
	Args:         cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		flags := cmd.Flags()

		// optional
		configPathRel, err := flags.GetString("configPath")
		if err != nil {
			return fmt.Errorf("unknown error when getting '--configPath'")
		}

		configPath, err := filepath.Abs(configPathRel)
		if err != nil {
			return fmt.Errorf("failed to resolve relative 'configPath' into absolute path")
		}

		// load config
		if err := config.Load(configPathRel == "", configPath); err != nil {
			return err
		}

		sessionTokens := viper.GetStringSlice("discordsessiontokens")

		for index := range sessionTokens {
			for _, arg := range args {
				argNumber, err := strconv.Atoi(arg)
				if err == nil && argNumber == index {
					sessionTokens[index] = "<DELETED>"
				}
			}
		}

		resultingSessionTokens := []string{}
		for _, token := range sessionTokens {
			if token != "<DELETED>" {
				resultingSessionTokens = append(resultingSessionTokens, token)
			}
		}

		viper.Set("discordsessiontokens", resultingSessionTokens)

		return config.Offload()
	},
}

func init() {
	flags := configRemoveDiscordTokensCmd.Flags()

	flags.StringP("configPath", "C", "", "a directory that contains an override config.json file\nor a file which contains an override config\n[a custom config file must end in .json]")

	configRemoveCmd.AddCommand(configRemoveDiscordTokensCmd)
}
