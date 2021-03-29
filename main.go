/*
	MSSQL Launcher by Aashish Koirala (c) 2021

	This file is part of mssql-launcher.

	mssql-launcher is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	mssql-launcher is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with mssql-launcher. If not, see <http://www.gnu.org/licenses/>.
*/

// mssql-launcher is an interactive tool that lets you define SQL Server connections
// in a file and pick them from a menu to connect using either the MSSQL-CLI tool or
// the SQLCMD tool.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aashishkoirala/mssql-launcher/command"
	"github.com/aashishkoirala/mssql-launcher/config"
	"github.com/aashishkoirala/mssql-launcher/logger"
	"github.com/aashishkoirala/mssql-launcher/menu"
)

func main() {
	os.Exit(run())
}

func run() int {
	var cfgFile, logFile string
	flag.StringVar(&cfgFile, "config", "", "Path to the config file to use.")
	flag.StringVar(&logFile, "log", "", "Path to the log file to use.")
	flag.Parse()

	logger, logcloser := logger.Get(logFile)
	defer logcloser()
	logger.Println("*** STARTING NEW SESSION ***")

	cfg, err := config.Get(cfgFile, logger)
	if c := handleError(&err, "reading configuration", logger); c != 0 {
		return c
	}

	item, err := menu.ReadChoice(cfg)
	if c := handleError(&err, "figuring out what you want", logger); c != 0 {
		return c
	}

	if item == nil {
		logger.Println("User exited due to choosing nothing.")
		fmt.Println("You chose nothing. Well, okay then.")
		return 0
	}

	cmd, args, err := command.Get(item, logger)
	if c := handleError(&err, "establishing the command to run", logger); c != 0 {
		return c
	}

	err = command.Run(cmd, args, logger)
	if c := handleError(&err, "running the command for what you picked", logger); c != 0 {
		return c
	}

	return 0
}

func handleError(err *error, doing string, logger *log.Logger) int {
	if *err == nil {
		return 0
	}

	logger.Printf("%v\n", *err)
	fmt.Printf("There was an error while %s. Please see the log file for details.", doing)
	return 1
}
