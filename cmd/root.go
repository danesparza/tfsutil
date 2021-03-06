package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/logutils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile               string
	ProblemWithConfigFile bool
	tfsurl                string
	personalaccesstoken   string
	collection            string
	project               string
	loglevel              string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tfsutil",
	Short: "A set of utilities for TFS",
	Long: `A set of command line utilities to make your
life a little easier when working with Team Foundation Server.  

NOTE: tfsutil uses the TFS API and it requires credentials.  
To set the personal access token (PAT) credentials used with 
each command, pass them in using flags or create a config file.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/tfsutil.yml)")

	rootCmd.PersistentFlags().StringVarP(&project, "project", "p", "", "TFS project")
	rootCmd.PersistentFlags().StringVarP(&tfsurl, "url", "u", "", "TFS root url")
	rootCmd.PersistentFlags().StringVarP(&personalaccesstoken, "pat", "t", "", "Personal access token (available in TFS)")
	rootCmd.PersistentFlags().StringVarP(&collection, "collection", "c", "DefaultCollection", "TFS collection")
	rootCmd.PersistentFlags().StringVarP(&loglevel, "loglevel", "l", "WARN", "Log level: DEBUG/INFO/WARN/ERROR")

	//	Bind config flags for optional config file override:
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("pat", rootCmd.PersistentFlags().Lookup("pat"))
	viper.BindPFlag("collection", rootCmd.PersistentFlags().Lookup("collection"))
	viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("project"))
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetConfigName("tfsutil") // name of config file (without extension)
		viper.AddConfigPath(home)      // adding home directory as first search path
		viper.AddConfigPath(".")       // also look in the working directory
	}

	viper.AutomaticEnv() // read in environment variables that match

	//	Set our defaults
	viper.SetDefault("tfsurl", "")
	viper.SetDefault("pat", "")
	viper.SetDefault("collection", "")
	viper.SetDefault("project", "")
	viper.SetDefault("loglevel", "WARN")

	// If a config file is found, read it in
	// otherwise, make note that there was a problem
	if err := viper.ReadInConfig(); err != nil {
		ProblemWithConfigFile = true
	}

	//	Set the log level from config (if we have it)
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(viper.GetString("loglevel")),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	//	If we have a config file, report it:
	if viper.ConfigFileUsed() != "" {
		log.Println("[DEBUG] Using config file:", viper.ConfigFileUsed())
	}

	//	If we have  tfs url or a PAT set, indicate it:
	if viper.GetString("tfsurl") != "" {
		log.Printf("[DEBUG] Using TFS url: \n%s\n", viper.GetString("tfsurl"))
	}

	if viper.GetString("pat") != "" {
		pat := viper.GetString("pat")
		log.Printf("[DEBUG] Using PAT that starts with: '%s'", pat[:4])
	}

	//	If we have  tfs collection or project set, indicate it:
	if viper.GetString("collection") != "" {
		log.Printf("[DEBUG] Using TFS collection: %s\n", viper.GetString("collection"))
	}

	if viper.GetString("project") != "" {
		log.Printf("[DEBUG] Using TFS project: %s\n", viper.GetString("project"))
	}
}
