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

package command

import (
	"fmt"
	"log"

	"github.com/aashishkoirala/mssql-launcher/config"
)

type sqlcli struct {
}

func (_ sqlcli) getCommand(connection *config.Connection, password string, isWindows bool, logger *log.Logger) (cmd string, args string) {

	logger.Printf("Connecting to %s: Tool = sqlcli, Server = %s, Database = %s, Integrated = %v, Windows = %v",
		connection.Name, connection.Server, connection.Database, connection.Integrated, isWindows)

	fmt.Println("Connecting to", connection.Name, "using mssql-cli...")

	cmd = "mssql-cli"
	args = ""
	if isWindows {
		args = "/c python -m mssqlcli.main "
		cmd = "cmd"
	}
	args = fmt.Sprintf("%s-S %s -d %s", args, connection.Server, connection.Database)
	if connection.Integrated {
		args = fmt.Sprintf("%s -E", args)
	} else {
		args = fmt.Sprintf("%s -U %s -P %s", args, connection.Username, password)
	}
	return
}
