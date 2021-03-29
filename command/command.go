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

// Package "command" provides logic to get the command and arguments corresponding
// to a connection entry, as well as running that command.
package command

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/aashishkoirala/mssql-launcher/config"
)

type commandGetter interface {
	getCommand(connection *config.Connection, password string, isWindows bool, logger *log.Logger) (cmd string, args string)
}

var decryptPassword func(encrypted string, decryptor string, logger *log.Logger) (string, error)

func init() {
	decryptPassword = decryptPassword_Default
}

// Get returns the command and arguments that need to be run to connect to the given
// connection entry.
func Get(item *config.Connection, logger *log.Logger) (string, string, error) {
	tool, err := getTool(item)
	if err != nil {
		return "", "", err
	}
	logger.Printf("For connection %s, tool is %T", item.Name, tool)

	password := item.Password
	if item.PasswordDecryptor != "" {
		password, err = decryptPassword(password, item.PasswordDecryptor, logger)
		if err != nil {
			return "", "", fmt.Errorf("Error decrypting password: %w", err)
		}
	}

	os := runtime.GOOS
	cmd, args := tool.getCommand(item, password, os == "windows", logger)

	return cmd, args, nil
}

func getTool(item *config.Connection) (commandGetter, error) {
	var tool commandGetter
	if item.Tool == "" {
		item.Tool = "mssql-cli"
	}
	switch item.Tool {
	case "mssql-cli":
		tool = sqlcli{}
	case "sqlcmd":
		tool = sqlcmd{}
	default:
		return nil, fmt.Errorf("Invalid tool: %s", item.Tool)
	}

	return tool, nil
}

func decryptPassword_Default(encrypted string, decryptor string, logger *log.Logger) (string, error) {
	logger.Println("Decrypting password...")

	decryptor = strings.Replace(decryptor, "${_PWD_}", encrypted, -1)
	args := strings.Split(decryptor, " ")
	if len(args) < 2 {
		return "", errors.New("Decryptor command line has too few arguments.")
	}
	cmd := args[0]
	args = args[1:]
	c := exec.Command(cmd, args...)
	decrypted, err := c.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

// Run runs the given command and arguments and in interactive mode.
func Run(cmd string, args string, logger *log.Logger) error {
	logger.Printf("Running command %s (args hidden for security)...", cmd)

	c := exec.Command(cmd, strings.Split(args, " ")...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin

	if err := c.Run(); err != nil {
		return err
	}
	return nil
}
