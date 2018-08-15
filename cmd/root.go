/*
Greenplum Magic Tool

Authored by Tyler Ramer, Ignacio Elizaga
Copyright 2018

Licensed under the Apache License, Version 2.0 (the "License")

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)


// constants
const (
	toolName = "gpmt"
	version = "Version ALPHA 1"
	githubRepo = "https://github.com/pivotal-gss/gpmt2"
)

// global variables
var (
	LCFlags LogCollectorFlags
	verbose bool
	logfile bool
	logDestination string
	hostname string
	port int
	username string
	password string
	database string
)

// The root CLI.
var rootCmd = &cobra.Command{
	Use:   toolName,
	Short: "Diagnostic and data collection for Greenplum Database",
	Long:  "\nGreenplum Magic Tool is a collection of diagnostic and data collection tools to " +
		   "assist in troubleshooting issues with Greenplum Database. \n" +
		   "Documentation and development information is available at: " + githubRepo,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Before running any command setup the logger
		SetupLogger()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// if no argument specified throw the help menu on the screen
		cmd.Help()
	},
}

// Sub Command: Version
// When this command is used the version of the gpmt is displayed on the screen
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "GPDB Version number",
	Long:  `Greenplum Magic Tool version`,
	Run: func(cmd *cobra.Command, args []string) {
		// print the version number on the screen when asked.
		fmt.Printf("%s: %s \n", cmd.Long, version)
	},
}

// Sub Command: Log Collector
// This command line arguments helps to obtain the logs from the greenplum database
var logCollectorCmd = &cobra.Command{
	Use:   "gp_log_collector",
	Short: "easy log collection",
	Long:  "\ngp_log_collector is used to automate Greenplum database log collection. \n" +
		   "Run without options, gp_log_collector will gather today's master and standby logs",
	Run: func(cmd *cobra.Command, args []string) {
		// log collect
		fmt.Println("I'll be a log collector one day")
	},
}

// -failed-segs only failed segs
// -free-space threshold
// -c contents
// -hostfile
// -h hostnames
// -start
// -end
// -a no propmt
// -dir
// -segdir
// -skip-master
// -standby (?)

// All the usage flags of the log collector
func flagsLogCollector() {
	logCollectorCmd.Flags().BoolVar(&LCFlags.failedOnly, "failed-segs", false, "Query gp_configuration_history for list of faulted content ids")
	logCollectorCmd.Flags().IntVar(&LCFlags.freeSpace, "free-space", 10, "default=10  Free space threshold which will abort log collection if reached")
	logCollectorCmd.Flags().StringArrayVar(&LCFlags.contentIds, "c", nil, "Space seperated list of content ids")
	logCollectorCmd.Flags().BoolVar(&LCFlags.noPrompt, "no-prompts", false, "Accept all prompts")
	logCollectorCmd.Flags().StringVarP(&LCFlags.hostfile, "hostfile", "f", "", "Read hostnames from a hostfile")
	logCollectorCmd.Flags().StringArrayVarP(&LCFlags.hostnames, "hostnames", "n", nil, "Space seperated list of hostnames")
	// FIXME: If date is empty string startDate and endDate it should default to current date
	logCollectorCmd.Flags().StringVar(&LCFlags.startDate, "start", "", "Start date for logs to collect (defaults to current date)")
	logCollectorCmd.Flags().StringVar(&LCFlags.endDate, "end", "", "End date for logs to collect (defaults to current date)")
	// FIXME: If workingDir is empty string it should default to cwd
	logCollectorCmd.Flags().StringVar(&LCFlags.workingDir, "dir", "", "Working directory (defaults to current directory)")
	// FIXME: If segmentDir is empty string it should default to /tmp
	logCollectorCmd.Flags().StringVar(&LCFlags.segmentDir, "segdir", "", "Segment temporary directory (defaults to /tmp)")
	logCollectorCmd.Flags().BoolVar(&LCFlags.osOnly, "os-only", false, "Only collect minimal infrastucture information")
	logCollectorCmd.Flags().BoolVar(&LCFlags.standby, "collect-standby", false, "Collect information from the standby master")
}


// Initialize the cobra command CLI.
func init() {

	// All global flag
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,"Enable verbose or debug logging")
	rootCmd.PersistentFlags().BoolVarP(&logfile, "log-file", "l", false, "Enable recording all the log messages to the logfile")
	rootCmd.PersistentFlags().StringVarP(&logDestination, "log-destination", "d", "/tmp", "Directory where the logfile should be created, only works with --log-file flag")

	// Database connection parameters.
	rootCmd.PersistentFlags().StringVar(&hostname, "hostname", "localhost","Hostname where the database is hosted")
	rootCmd.PersistentFlags().IntVar(&port, "port", 5432, "Port number of the master database")
	rootCmd.PersistentFlags().StringVar(&database, "database", "template1", "Database name to connect")
	rootCmd.PersistentFlags().StringVar(&username, "username", "gpadmin", "Username that is used to connect to database")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "password for the user")

	// Attach the sub command to the root command.
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(logCollectorCmd)
	flagsLogCollector()

}

// Execute the cobra CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
