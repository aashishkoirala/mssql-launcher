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

package config

import (
	_ "embed"
	"log"
	"strings"
	"testing"
)

//go:embed test_valid.yaml
var testValidConfigYaml []byte

//go:embed test_invalid.yaml
var testInvalidConfigYaml []byte

func TestGet_Works_With_Valid_Config(t *testing.T) {
	getConfigYaml = getConfigYaml_TestValid
	config, err := Get("", log.Default())
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	len := len(*config)
	if len != 2 {
		t.Fatalf("Expected 2 connections, got %d.", len)
	}
}

func TestGet_Fails_With_Invalid_Config(t *testing.T) {
	getConfigYaml = getConfigYaml_TestInvalid
	_, err := Get("", log.Default())
	if err == nil {
		t.Fatal("Expected error but did not get one.")
	}
	msg := err.Error()
	if msg != getExpectedValidationErrors() {
		t.Fatalf("Validation errors don't match expected ones, got the following: %s", msg)
	}
}

func getConfigYaml_TestValid(cfgFile string, logger *log.Logger) ([]byte, error) {
	return testValidConfigYaml, nil
}

func getConfigYaml_TestInvalid(cfgFile string, logger *log.Logger) ([]byte, error) {
	return testInvalidConfigYaml, nil
}

func getExpectedValidationErrors() string {
	errs := [5]string{
		"Found more than one items with the name Test 2.",
		"Test 3: Server not set.",
		"An item does not have name set.",
		"Test 4: Database not set.",
		"Test 5: Username set but password not set.",
	}
	return strings.Join(errs[:], "\n")
}
