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

// Package "logger" wraps a file based logger and provides logger
// initialization logic.
package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// Get initializes a file logger for the given file path (or using the default)
// logging file path if not given and returns the resulting logger and a function
// that should be called to close the logger.
func Get(logFile string) (*log.Logger, func()) {
	logfilepath := logFile
	if logfilepath == "" {
		logfilepath = getLogFilePath()
	}
	if logfilepath == "" {
		return log.Default(), func() {}
	}
	logfile, err := os.OpenFile(logfilepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return log.Default(), func() {}
	}
	closer := func() {
		logfile.Close()
	}
	writer := io.Writer(logfile)
	newLogger := log.New(writer, "[mssql-launcher] ", log.Ldate|log.Ltime)
	return newLogger, closer
}

func getLogFilePath() string {
	exe, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Join(filepath.Dir(exe), "mssql-launcher-log.log")
}
