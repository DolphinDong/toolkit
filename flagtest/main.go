package main

import (
	"demo/flagtest/config"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var (
	configPath string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config.yaml", "config path")

	rootCmd.Flags().StringVar(&config.Cfg.Username, "username", "default value", "Username for database")
	rootCmd.Flags().StringVar(&config.Cfg.Password, "password", "default value", "Password for database")
	rootCmd.Flags().StringVar(&config.Cfg.DB, "db", "default value", "Database name")
	rootCmd.Flags().StringVar(&config.Cfg.Host, "host", "default value", "Database host")
	rootCmd.Flags().StringVar(&config.Cfg.Other.Args1, "other.args1", "default value", "Other Args1")
	rootCmd.Flags().StringVar(&config.Cfg.Other.Args2, "other.args2", "default value", "Other Args2")
	rootCmd.Flags().StringVar(&config.Cfg.Other.Args3, "other.args3", "default value", "Other Args3")
	rootCmd.Flags().StringVar(&config.Cfg.Other.Other2.Args1, "other.other2.args1", "default value", "Other2 Args1")
	rootCmd.Flags().StringVar(&config.Cfg.Other.Other2.Args2, "other.other2.args2", "default value", "Other2 Args2")
	rootCmd.Flags().StringVar(&config.Cfg.Other.Other2.Args3, "other.other2.args3", "default value", "Other2 Args3")

	// 绑定 flag 到 viper
	err := viper.BindPFlag("username", rootCmd.Flags().Lookup("username"))
	err = viper.BindPFlag("password", rootCmd.Flags().Lookup("password"))
	err = viper.BindPFlag("db", rootCmd.Flags().Lookup("db"))
	err = viper.BindPFlag("host", rootCmd.Flags().Lookup("host"))
	err = viper.BindPFlag("other.args1", rootCmd.Flags().Lookup("other.args1"))
	err = viper.BindPFlag("other.args2", rootCmd.Flags().Lookup("other.args2"))
	err = viper.BindPFlag("other.args3", rootCmd.Flags().Lookup("other.args3"))
	err = viper.BindPFlag("other.other2.args1", rootCmd.Flags().Lookup("other.other2.args1"))
	err = viper.BindPFlag("other.other2.args2", rootCmd.Flags().Lookup("other.other2.args2"))
	err = viper.BindPFlag("other.other2.args3", rootCmd.Flags().Lookup("other.other2.args3"))
	if err != nil {
		log.Fatalf("viper bind flag error: %+v", err)
	}

	// 绑定环境变量
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	err = viper.BindEnv("username", "APP_USERNAME")
	err = viper.BindEnv("password", "APP_PASSWORD")
	err = viper.BindEnv("db", "APP_DB")
	err = viper.BindEnv("host", "APP_HOST")
	err = viper.BindEnv("other.args1", "APP_OTHER_ARGS1")
	err = viper.BindEnv("other.args2", "APP_OTHER_ARGS2")
	err = viper.BindEnv("other.args3", "APP_OTHER_ARGS3")
	err = viper.BindEnv("other.other2.args1", "APP_OTHER_OTHER2_ARGS1")
	err = viper.BindEnv("other.other2.args2", "APP_OTHER_OTHER2_ARGS2")
	err = viper.BindEnv("other.other2.args3", "APP_OTHER_OTHER2_ARGS3")
	if err != nil {
		log.Fatalf("viper bind env error: %+v", err)
	}

}

func initConfig() {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
}

func setConfigValue() {
	config.Cfg.Username = viper.GetString("username")
	config.Cfg.Password = viper.GetString("password")
	config.Cfg.DB = viper.GetString("db")
	config.Cfg.Host = viper.GetString("host")
	config.Cfg.Other.Args1 = viper.GetString("other.args1")
	config.Cfg.Other.Args2 = viper.GetString("other.args2")
	config.Cfg.Other.Args3 = viper.GetString("other.args3")
	config.Cfg.Other.Other2.Args1 = viper.GetString("other.other2.args1")
	config.Cfg.Other.Other2.Args2 = viper.GetString("other.other2.args2")
	config.Cfg.Other.Other2.Args3 = viper.GetString("other.other2.args3")
}

var rootCmd = &cobra.Command{
	Use:   "flagtest",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		setConfigValue()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Username: %s\n", config.Cfg.Username)
		fmt.Printf("Password: %s\n", config.Cfg.Password)
		fmt.Printf("Database: %s\n", config.Cfg.DB)
		fmt.Printf("Host: %s\n", config.Cfg.Host)
		fmt.Printf("Other.Args1: %s\n", config.Cfg.Other.Args1)
		fmt.Printf("Other.Args2: %s\n", config.Cfg.Other.Args2)
		fmt.Printf("Other.Args3: %s\n", config.Cfg.Other.Args3)
		fmt.Printf("Other.Other2.Args1: %s\n", config.Cfg.Other.Other2.Args1)
		fmt.Printf("Other.Other2.Args2: %s\n", config.Cfg.Other.Other2.Args2)
		fmt.Printf("Other.Other2.Args3: %s\n", config.Cfg.Other.Other2.Args3)
		fmt.Printf("%#v", config.Cfg)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
