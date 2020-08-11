package cli

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userID string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "DeputyTest\n",
	Short: "\nDeputy tech challenge by Justin Davidson (justindavidson23@gmail.com)",
	Long:  "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println("rootCmd.Execute error:", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&userID, "userID", "", "UserID of which to find sub-ordinates of")
}

func initConfig() {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	dir, _ := os.Getwd()
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	// working dir for docker image
	viper.AddConfigPath(dir)

	// basepath/config folder for running locally
	viper.AddConfigPath(fmt.Sprintf("%v/config", basepath))
	viper.Unmarshal(&cfg)

	fmt.Printf("cfg?: %+v\n", cfg)
}
