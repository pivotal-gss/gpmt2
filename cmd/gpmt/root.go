/*
Greenplum Magic Tool

Authored by Tyler Ramer, Ignacio Elizaga
Copyright 2018

Licensed under the Apache License, Version 2.0 (the "License")

*/
package main

import (
	"fmt"
	"time"

	"github.com/pivotal-gss/gpmt2/pkg/db"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Tool information constants

const (
	gpmtVersion = "Version (pre)ALPHA"
	githubRepo  = "https://github.com/pivotal-gss/gpmt2"
)

type logOptions struct {
	Verbose bool
	LogDir  string
	LogFile string
}

// Local Package Variables
var (
	// gp_log_collector flags
	lcOpts LogCollectorOptions

	// DB connection details
	connString db.ConnString //FIXME/TODO: Do we need a separate wrapper for DB?

	// logging flags
	logOpts = logOptions{LogFile: fmt.Sprintf("/gpmt_log_%s", time.Now().Format("2006-01-02"))}
)

// Sub Command: Version
// When this command is used the version of the gpmt is displayed on the screen
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "GPDB Version number",
	Long:  `Greenplum Magic Tool version`,
	Run: func(cmd *cobra.Command, args []string) {
		// print the version number on the screen when asked.
		fmt.Printf("%s: %s \n", cmd.Long, gpmtVersion)
	},
}

// The root CLI.
var rootCmd = &cobra.Command{
	Use:   "gpmt [options]",
	Short: "Diagnostic and data collection for Greenplum Database",
	Long: "\nGreenplum Magic Tool is a collection of diagnostic and data collection tools to " +
		"assist in troubleshooting issues with Greenplum Database. \n" +
		"Documentation and development information is available at: " + githubRepo,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Before running any command
		// Setup the logger log level
		if logOpts.Verbose {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}

		// Set logger logfile  hooks
		logName := logOpts.LogDir + logOpts.LogFile
		formatter := &log.TextFormatter{
			TimestampFormat : "2006-01-02 15:04:05",
			FullTimestamp: true,
		}
		log.SetFormatter(formatter)
		log.AddHook(lfshook.NewHook(logName, &log.TextFormatter{}))

	},
	Run: func(cmd *cobra.Command, args []string) {
		// if no argument specified throw the help menu on the screen
		cmd.Help()
	},
}

// Initialize the cobra command CLI.
func init() {

	// All global flag
	rootCmd.PersistentFlags().BoolVarP(&logOpts.Verbose, "verbose", "v", false, "Enable verbose or debug logging")
	rootCmd.PersistentFlags().StringVar(&logOpts.LogDir, "log-directory", "/tmp", "Directory where the logfile should be created") // TODO - logfile default may change

	// Database connection parameters.
	rootCmd.PersistentFlags().StringVar(&connString.Hostname, "hostname", "localhost", "Hostname where the database is hosted")
	rootCmd.PersistentFlags().IntVar(&connString.Port, "port", 5432, "Port number of the master database")
	rootCmd.PersistentFlags().StringVar(&connString.Database, "database", "template1", "Database name to connect")
	rootCmd.PersistentFlags().StringVar(&connString.Username, "username", "gpadmin", "Username that is used to connect to database")
	rootCmd.PersistentFlags().StringVar(&connString.Password, "password", "", "Password for the user")

	// Attach the sub command to the root command.
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(logCollectorCmd)
	flagsLogCollector()

}

// Execute the cobra CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}
