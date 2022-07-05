package main

import (
	"github.com/spf13/cobra"
	"log"
)

const command = "deterge"
const description = "Tool for sanitising data files of sensitive information through substitution with fake information."

func main() {
	if err := defineCommands().Execute(); err != nil {
		log.Fatal(err)
	}
}

func defineCommands() *cobra.Command {
	cobra.OnInitialize(initConfig)

	ui := &cobra.Command{
		Use:   command,
		Short: description,
		Run: func(cmd *cobra.Command, args []string) {
			// Show help when no command is specified
			_ = cmd.Help()
		},
	}

	var thing string

	sqlCommand := &cobra.Command{
		Use: "sql",
		Run: func(cmd *cobra.Command, args []string) {
			// -in io.Reader
			// -in io.Writer
			// the "preset" can be used to customise the options
			// detergeSQL(in io.Reader, out io.Writer, options)
		},
	}

	sqlCommand.PersistentFlags().StringVar(&thing, "thing", "", "")
	ui.AddCommand(sqlCommand)

	return ui
}

func initConfig() {
	//// Don't forget to read config either from cfgFile or from home directory!
	//if cfgFile != "" {
	//	// Use config file from the flag.
	//	viper.SetConfigFile(cfgFile)
	//} else {
	//	// Find home directory.
	//	home, err := homedir.Dir()
	//	if err != nil {
	//		fmt.Println(err)
	//		os.Exit(1)
	//	}
	//
	//	// Search config in home directory with name ".cobra" (without extension).
	//	viper.AddConfigPath(home)
	//	viper.SetConfigName(".cobra")
	//}
	//
	//if err := viper.ReadInConfig(); err != nil {
	//	fmt.Println("Can't read config:", err)
	//	os.Exit(1)
	//}
}
