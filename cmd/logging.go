/*
Greenplum Magic Tool

Authored by Tyler Ramer, Ignacio Elizaga
Copyright 2018

Licensed under the Apache License, Version 2.0 (the "License")

*/
package cmd

import (
	"github.com/op/go-logging"
	"os"
	"fmt"
	"time"
)

// Logger name and logging format
var (
	log = logging.MustGetLogger(toolName)
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfile}(%{shortfunc}) ▶ %{level:.4s} ▶ %{color:reset} %{message}`,
	)
)

// Setup the logger.
func SetupLogger()  {

	// New On Screen logger
	backendScreen := logging.NewLogBackend(os.Stderr, "", 0)
	screenFormatted := logging.NewBackendFormatter(backendScreen, format)
	screenLeveled := logging.AddModuleLevel(screenFormatted)

	// Setup logfile logger
	logFilename := fmt.Sprintf(logDestination + "/gpmt_log_%s.log", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(logFilename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	backendFile := logging.NewLogBackend(file, "", 0)
	fileFormatter := logging.NewBackendFormatter(backendFile, format)
	fileLeveled := logging.AddModuleLevel(fileFormatter)

	// Setup the logging level based on users request
	if verbose {
		screenLeveled.SetLevel(logging.DEBUG, "")
		fileLeveled.SetLevel(logging.DEBUG, "")
	} else {
		screenLeveled.SetLevel(logging.INFO, "")
		fileLeveled.SetLevel(logging.INFO, "")
	}

	// If the user wants all the logs on the logfile save them on the location ( if provided )
	// or use the default destination that is /tmp
	if logfile {
		logging.SetBackend(screenLeveled, fileLeveled)
	} else { // Start the screen only logger
		logging.SetBackend(screenLeveled)
	}

	// If the file creation failed, then place a warning message on the screen.
	// NOTE: the reason why place the err catcher here is simply because the logger has not
	// initialized yet, until we set the backend.
	if err != nil {
		log.Warningf("Cannot log messages onto logfile: \"%s\", ERROR: %v", logFilename, err)
		log.Infof("Continuing the script with on screen logging...")
	} else if logfile {
		log.Debugf("All %s tool log messages are logged into: \"%s\"", toolName, logFilename)
	}
}
