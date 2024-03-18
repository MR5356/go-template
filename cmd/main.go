package cmd

import (
	"github.com/MR5356/go-template/config"
	"github.com/MR5356/go-template/pkg/server"
	"github.com/spf13/cobra"
)

var (
	port, gracePeriod int
	debug             bool
	dbDriver, dbDSN   string
)

func NewApplication() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.New(
				config.WithPort(port),
				config.WithDebug(debug),
				config.WithGracePeriod(gracePeriod),
				config.WithDatabase(dbDriver, dbDSN),
			)

			srv, err := server.New(cfg)
			if err != nil {
				return err
			}
			return srv.Run()
		},
	}

	// 屏蔽 cobra 自带的帮助信息
	cmd.SilenceUsage = true
	// 屏蔽 cobra 自带的错误信息
	cmd.SilenceErrors = true

	// 注册命令参数
	cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")
	cmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "server port")
	cmd.PersistentFlags().IntVar(&gracePeriod, "grace-period", 30, "server grace period")
	cmd.PersistentFlags().StringVar(&dbDriver, "db-driver", "sqlite", "database driver")
	cmd.PersistentFlags().StringVar(&dbDSN, "db-dsn", "db.sqlite", "database dsn")

	return cmd
}
