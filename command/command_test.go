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
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/aashishkoirala/mssql-launcher/config"
)

var decryptPasswordCalled bool

func TestGetTool_Works_For_SqlCli(t *testing.T) {
	c := &config.Connection{Tool: "mssql-cli"}
	tool, err := getTool(c)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	_, ok := tool.(sqlcli)
	if !ok {
		t.Fatal("Did not get sqlcli instance.")
	}
}

func TestGetTool_Works_For_SqlCmd(t *testing.T) {
	c := &config.Connection{Tool: "sqlcmd"}
	tool, err := getTool(c)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	_, ok := tool.(sqlcmd)
	if !ok {
		t.Fatal("Did not get sqlcmd instance.")
	}
}

func TestGetTool_Empty_Treated_As_SqlCli(t *testing.T) {
	c := &config.Connection{}
	tool, err := getTool(c)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	_, ok := tool.(sqlcli)
	if !ok {
		t.Fatal("Did not get sqlcli instance.")
	}
}

func TestGetTool_Invalid_Fails(t *testing.T) {
	c := &config.Connection{Tool: "Invalid"}
	_, err := getTool(c)
	if err == nil {
		t.Fatal("Expected error, but did not get one.")
	}
	msg := err.Error()
	if msg != "Invalid tool: Invalid" {
		t.Fatalf("Did not get expected error message, got this: %s", msg)
	}
}

func TestGet_Decrypt_Password_Not_Called_For_Empty_Decryptor(t *testing.T) {
	var called bool
	decryptPassword = func(encrypted string, decryptor string, logger *log.Logger) (string, error) {
		called = true
		return decryptPassword_Test_Valid(encrypted, decryptor, logger)
	}
	c := &config.Connection{
		Server:   "s1",
		Database: "d1",
		Username: "u1",
		Password: "p1",
	}
	_, _, err := Get(c, log.Default())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if called {
		t.Fatal("Did not expect decryptPassword to be called.")
	}
}

func TestGet_Decrypt_Password_Called_For_Non_Empty_Decryptor(t *testing.T) {
	var called bool
	decryptPassword = func(encrypted string, decryptor string, logger *log.Logger) (string, error) {
		called = true
		return decryptPassword_Test_Valid(encrypted, decryptor, logger)
	}
	c := &config.Connection{
		Server:            "s1",
		Database:          "d1",
		Username:          "u1",
		Password:          "p1",
		PasswordDecryptor: "d1",
	}
	_, args, err := Get(c, log.Default())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if !called {
		t.Fatal("Expected decryptPassword to be called.")
	}
	if !strings.Contains(args, "Decrypted:p1") {
		t.Fatal("Decrypted password not passed back properly.")
	}
}

func TestGet_Decrypt_Password_Error_For_Non_Empty_Decryptor(t *testing.T) {
	expectedErr := errors.New("Decryption Error")
	decryptPassword = func(_ string, _ string, _ *log.Logger) (string, error) {
		return "", expectedErr
	}
	c := &config.Connection{
		Server:            "s1",
		Database:          "d1",
		Username:          "u1",
		Password:          "p1",
		PasswordDecryptor: "d1",
	}
	_, _, err := Get(c, log.Default())
	if err == nil {
		t.Fatalf("Expected error.")
	}
	if errors.Unwrap(err) != expectedErr {
		t.Fatalf("Expected %v, got %v.", expectedErr, err)
	}
}

func decryptPassword_Test_Valid(encrypted string, decryptor string, logger *log.Logger) (string, error) {
	return fmt.Sprintf("Decrypted:%s", encrypted), nil
}
