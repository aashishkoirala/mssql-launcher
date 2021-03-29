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
	"log"
	"testing"

	"github.com/aashishkoirala/mssql-launcher/config"
)

func TestSqlCmd_GetCommand_Works_For_Linux_Non_Integrated(t *testing.T) {
	c := &config.Connection{
		Name:     "test",
		Server:   "server1",
		Database: "database1",
		Username: "user1",
		Password: "password1",
	}
	s := sqlcmd{}
	cmd, args := s.getCommand(c, c.Password, false, log.Default())
	const (
		expectedCmd  string = "sqlcmd"
		expectedArgs string = "-S server1 -d database1 -U user1 -P password1"
	)
	if cmd != expectedCmd {
		t.Fatalf("Expected cmd = %s, got %s", expectedCmd, cmd)
	}
	if args != expectedArgs {
		t.Fatalf("Expected args = %s, got %s", expectedArgs, args)
	}
}

func TestSqlCmd_GetCommand_Works_For_Windows_Non_Integrated(t *testing.T) {
	c := &config.Connection{
		Name:     "test",
		Server:   "server2",
		Database: "database2",
		Username: "user2",
		Password: "password2",
	}
	s := sqlcmd{}
	cmd, args := s.getCommand(c, c.Password, true, log.Default())
	const (
		expectedCmd  string = "cmd"
		expectedArgs string = "/c sqlcmd -S server2 -d database2 -U user2 -P password2"
	)
	if cmd != expectedCmd {
		t.Fatalf("Expected cmd = %s, got %s", expectedCmd, cmd)
	}
	if args != expectedArgs {
		t.Fatalf("Expected args = %s, got %s", expectedArgs, args)
	}
}

func TestSqlCmd_GetCommand_Works_For_Windows_Integrated(t *testing.T) {
	c := &config.Connection{
		Name:       "test",
		Server:     "server3",
		Database:   "database3",
		Integrated: true,
	}
	s := sqlcmd{}
	cmd, args := s.getCommand(c, c.Password, true, log.Default())
	const (
		expectedCmd  string = "cmd"
		expectedArgs string = "/c sqlcmd -S server3 -d database3 -E"
	)
	if cmd != expectedCmd {
		t.Fatalf("Expected cmd = %s, got %s", expectedCmd, cmd)
	}
	if args != expectedArgs {
		t.Fatalf("Expected args = %s, got %s", expectedArgs, args)
	}
}
