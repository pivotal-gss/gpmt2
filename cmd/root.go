/*
Greenplum Magic Tool

Authored by Tyler Ramer, Ignacio Elizaga
Copyright 2018

Licensed under the Apache License, Version 2.0 (the "License")

*/

// cobra command line

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(logCollectorCmd)
	flagsLogCollector()
}

var LCFlags LogCollectorFlags

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

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "GPDB Version number",
	Long:  `Greenplum Magic Tool version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version ALPHA 1")
	},
}

var rootCmd = &cobra.Command{
	Use:   "gpmt",
	Short: "GPMT - diagnostic and data collection for Greemplum Database",
	Long:  "Greenplum Magic Tool is a collection of diagnostic and data collection tools to assist in troubleshooting issues with Greenplum Database. Documentation and development information is available at https://github.com/pivotal-gss/gpmt2",
	Run: func(cmd *cobra.Command, args []string) {
		// probably just help
	},
}

var logCollectorCmd = &cobra.Command{
	Use:   "gp_log_collector",
	Short: "easy log collection",
	Long:  "gp_log_collector is used to automate Greenplum database log collection. Run without options, gp_log_collector will gather today's master and standby logs",
	Run: func(cmd *cobra.Command, args []string) {
		// log collect
		fmt.Println("I'll be a log collector one day")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
