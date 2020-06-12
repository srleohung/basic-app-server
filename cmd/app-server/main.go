package main

import (
	"basic-app-server/config"
	"basic-app-server/core"
	apperr "basic-app-server/errors"
	"basic-app-server/logger"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/coreos/go-systemd/daemon"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const SOFTWARE_NAME = "UNKNOWN_SOFTWARE_NAME"

var (
	BUILD_TAGS         = "UNKNOWN_BUILD_TAGS"
	VERSION_MAJOR      = "UNKNOWN_VERSION_MAJOR"
	VERSION_MINOR      = "UNKNOWN_VERSION_MINOR"
	VERSION_DERIVATIVE = "UNKNOWN_VERSION_DERIVATIVE"
	GIT_BRANCH         = "UNKNOWN_GIT_BRANCH"
	BUILD_DATE         = time.Now().Local().Format(time.RFC3339)
)

var cfgFile string
var mainLogger *logrus.Entry = logger.GetLogger("main")

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(configCmd)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml")
}

var RootCmd = &cobra.Command{
	Use:   SOFTWARE_NAME,
	Short: SOFTWARE_NAME + " is a basic application server.",
	Long:  "A basic application server built by Leo Hung. Complete documentation is available at https://github.com/srleohung/basic-app-server",
	Run: func(cmd *cobra.Command, args []string) {
		appServer()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + SOFTWARE_NAME,
	Long:  "All software has versions. This is " + SOFTWARE_NAME + `'s`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Software name      :", SOFTWARE_NAME)
		fmt.Println("Build tags         :", BUILD_TAGS)
		fmt.Println("Version major      :", VERSION_MAJOR)
		fmt.Println("Version minor      :", VERSION_MINOR)
		fmt.Println("Version Derivative :", VERSION_DERIVATIVE)
		fmt.Println("Git branch         :", GIT_BRANCH)
		fmt.Println("Build date         :", BUILD_DATE)
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the configuration of " + SOFTWARE_NAME,
	Run: func(cmd *cobra.Command, args []string) {
		configuration := loadConfig()
		if bytes, err := yaml.Marshal(*configuration); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(bytes))
		}
	},
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(apperr.EXITCODE_PROGRAM_COMMAND_ERROR)
	}
}

func appServer() {
	version := VERSION_MAJOR + "." + VERSION_MINOR + "_" + VERSION_DERIVATIVE
	mainLogger.Infoln("--------------------------------------------------")
	mainLogger.Infoln("Starting application server version:", version)

	appServer := appServerSetup(version)
	daemon.SdNotify(false, "READY=1")
	mainLogger.Infoln("--------------------------------------------------")
	mainLogger.Infoln("Finished application server configuration and setup.")

	appServer.Start()
	mainLogger.Infoln("--------------------------------------------------")
	mainLogger.Infoln("Start serving application server.")

	go watchdog()
	<-quit
}

var quit = make(chan bool)

func watchdog() {
	interval, err := daemon.SdWatchdogEnabled(false)
	if err != nil || interval == 0 {
		mainLogger.Warn("Watchdog is not enabled.")
		return
	}
	for {
		daemon.SdNotify(false, "WATCHDOG=1")
		time.Sleep(interval / 3)
	}
}

func loadConfig() *config.Config {
	var configuration config.Config
	useDefaultConfig := true
	if len(cfgFile) > 0 {
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			useDefaultConfig = true
		} else {
			mainLogger.Infoln("Load the application server configuration from", cfgFile)
			if tmpConfig := config.LoadConfigFromYaml(cfgFile); tmpConfig == nil {
				useDefaultConfig = true
			} else {
				configuration = *tmpConfig
				useDefaultConfig = false
			}
		}
	} else {
		workingDir, _ := os.Getwd()
		defaultConfigFile := path.Join(workingDir, config.DAFAULT_CONFIG_FILENAME)
		if _, err := os.Stat(defaultConfigFile); os.IsNotExist(err) {
			useDefaultConfig = true
		} else {
			if tmpConfig := config.LoadConfigFromYaml(defaultConfigFile); tmpConfig == nil {
				useDefaultConfig = true
			} else {
				configuration = *tmpConfig
				useDefaultConfig = false
			}
		}
	}

	if useDefaultConfig {
		mainLogger.Infoln("Load the default application server configuration.")
		configuration = config.DEFAULT_CONFIG
		if workingDir, err := os.Getwd(); err != nil {
			mainLogger.Errorln("The current working directory cannot be recognized.", err)
			os.Exit(apperr.EXITCODE_UNEXPECTED_ERROR)
		} else {
			if err := configuration.SaveConfigToYamlFile(path.Join(workingDir, config.DAFAULT_CONFIG_FILENAME)); err != nil {
				mainLogger.Warnln("Unable to save application server configuration file to " + path.Join(workingDir, config.DAFAULT_CONFIG_FILENAME))
			}
		}
	}
	return &configuration
}

func appServerSetup(version string) *core.AppServer {
	if configuration := loadConfig(); configuration != nil {
		return core.GetAppServer(configuration, version)
	} else {
		mainLogger.Errorln("The application server configuration is not recognized.")
		os.Exit(apperr.EXITCODE_UNEXPECTED_ERROR)
	}
	return nil
}
