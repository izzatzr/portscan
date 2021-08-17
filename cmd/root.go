package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	protocol string
	host     string
	port     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "portscan",
	Short: "A brief description of your application",
	Long:  `.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		protocols := []string{
			"tcp",
			"tcp4",
			"tcp6",
			"udp",
			"udp4",
			"udp6",
			"ip",
			"ip4",
			"ip6",
			"unix",
			"unixgram",
			"unixpacket",
			"tcp",
			"tcp4",
			"tcp6",
			"udp",
			"udp4",
			"udp6",
			"ip",
			"ip4",
			"ip6",
			"unix",
			"unixgram",
			"unixpacket"}

		for _, p := range protocols {
			if protocol != p {
				return fmt.Errorf("protocol %s not supported", protocol)
			}
		}

		conn, err := net.Dial(protocol, fmt.Sprintf("%s:%s", host, port))
		if err != nil {
			errors.Wrap(err, err.Error())
			return err
		}

		println("Port Opened")
		defer conn.Close()
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.portscan.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().StringVarP(&host, "host", "h", "", "Host")
	rootCmd.MarkFlagRequired("host")
	rootCmd.Flags().StringVarP(&port, "port", "p", "", "Port")
	rootCmd.MarkFlagRequired("port")

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".portscan")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
