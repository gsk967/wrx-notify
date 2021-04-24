package main

import (
	"errors"
	"fmt"
	"github.com/saikumar0x1967/wrx-notify/utils"
	"github.com/spf13/cobra"
	"os"
	"time"
)

const (
	FlagExchangesList    = "exchanges-list"
	FlagTimeInterval     = "watch"
	FlagTelegramNotify   = "tg-notify"
	FlagTelegramBotToken = "tg-token"
	FlagTelegramChatId   = "tg-chat-id"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "wrx-notify",
		Short: "wrx notify the price list of exchanges on wazirx ",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			tgNotify, err := cmd.Flags().GetBool(FlagTelegramNotify)
			if err != nil {
				return err
			}
			if tgNotify {
				tgBotToken, err := cmd.Flags().GetString(FlagTelegramBotToken)
				if err != nil {
					return err
				}
				if len(tgBotToken) == 0 || tgBotToken == "" {
					return errors.New("telegram bot token is required, --tg-token  110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw")
				}
				tgChatId, err := cmd.Flags().GetInt64(FlagTelegramChatId)
				if err != nil {
					return err
				}
				if tgChatId == 0 {
					return errors.New("telegram chat id is required, --tg-chat-id  110201543")
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var tgAlert *utils.TelegramAlert
			filePath, err := cmd.Flags().GetString(FlagExchangesList)
			if err != nil {
				return err
			}
			flagTimeInterval, err := cmd.Flags().GetUint64(FlagTimeInterval)
			if err != nil {
				return err
			}
			tgNotify, err := cmd.Flags().GetBool(FlagTelegramNotify)
			if err != nil {
				return err
			}
			if tgNotify {
				tgBotToken, err := cmd.Flags().GetString(FlagTelegramBotToken)
				if err != nil {
					return err
				}
				tgChatId, err := cmd.Flags().GetInt64(FlagTelegramChatId)
				if err != nil {
					return err
				}
				tgAlert = utils.NewTelegramAlerter(tgBotToken, tgChatId)
			}

			utils.StartWatchingExchanges(filePath, time.Duration(flagTimeInterval), tgAlert)
			return nil
		},
	}

	// flags
	rootCmd.Flags().Uint64(FlagTimeInterval, 10, "Time interval in seconds --watch 10")
	rootCmd.Flags().String(FlagExchangesList, "", "Json file for exchanges list : --exchanges-list exchanges.json")
	rootCmd.Flags().Bool(FlagTelegramNotify, false, "Telegram bot notifications : true/false")
	rootCmd.Flags().String(FlagTelegramBotToken, "", "Telegram bot authorization token : --tg-token 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw")
	rootCmd.Flags().Int64(FlagTelegramChatId, 0, "Telegram chat id  : --tg-chat-id 110201543")

	err := rootCmd.MarkFlagRequired(FlagExchangesList)
	if err != nil {
		panic(err)
	}

	err = rootCmd.MarkFlagRequired(FlagTimeInterval)
	if err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
