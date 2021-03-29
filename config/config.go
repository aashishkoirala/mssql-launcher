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

// Package "config" defines the structure of connections and provides logic
// to load and validate the configuration from YAML.
package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Connection represents each entry in the YAML file that represents a
// connection to a SQL Server database.
type Connection struct {
	Name              string
	Server            string
	Database          string
	Integrated        bool
	Username          string
	Password          string
	PasswordDecryptor string
	Tool              string
}

var (
	getConfigYaml func(cfgFile string, logger *log.Logger) ([]byte, error)
)

func init() {
	getConfigYaml = getConfigYaml_Default
}

// Get reads and returns an array of Connection instances given
// the path to the YAML config file - or the default path if not given.
func Get(cfgFile string, logger *log.Logger) (*[]Connection, error) {

	config, err := getConfigYaml(cfgFile, logger)
	if err != nil {
		return nil, err
	}

	logger.Println("Reading config YAML...")
	connectionList := make([]Connection, 0)
	err = yaml.Unmarshal([]byte(config), &connectionList)
	if err != nil {
		return nil, err
	}

	if len(connectionList) == 0 {
		return nil, fmt.Errorf("No connections defined.")
	}

	logger.Println("Validating configuration entries...")
	err = validateConfigs(&connectionList)
	if err != nil {
		return nil, err
	}
	return &connectionList, nil
}

func getConfigYaml_Default(cfgFile string, logger *log.Logger) ([]byte, error) {
	err := getConfigFilePath(&cfgFile)
	if err != nil {
		return nil, err
	}
	logger.Printf("Using config file %s", cfgFile)

	return ioutil.ReadFile(cfgFile)
}

func getConfigFilePath(cfgFile *string) error {
	if *cfgFile != "" {
		return nil
	}
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	*cfgFile = filepath.Join(filepath.Dir(exe), "mssql-launcher-config.yaml")
	return nil
}

func validateConfigs(connectionList *[]Connection) error {
	errormsgs := make([]string, 0)
	names := make(map[string]bool, 0)
	for _, c := range *connectionList {
		err := validateConfig(&c, &names)
		if err != nil {
			errormsgs = append(errormsgs, err.Error())
		}
	}
	if len(errormsgs) == 0 {
		return nil
	}
	return fmt.Errorf("%s", strings.Join(errormsgs, "\n"))
}

func validateConfig(c *Connection, names *map[string]bool) error {
	if isUnset(&c.Name) {
		return fmt.Errorf("An item does not have name set.")
	}
	_, exists := (*names)[c.Name]
	if exists {
		return fmt.Errorf("Found more than one items with the name %s.", c.Name)
	}
	(*names)[c.Name] = true
	if isUnset(&c.Server) {
		return fmt.Errorf("%s: Server not set.", c.Name)
	}
	if isUnset(&c.Database) {
		return fmt.Errorf("%s: Database not set.", c.Name)
	}
	if !isUnset(&c.Username) && isUnset(&c.Password) {
		return fmt.Errorf("%s: Username set but password not set.", c.Name)
	}
	if !isUnset(&c.Tool) && c.Tool != "mssql-cli" && c.Tool != "sqlcmd" {
		return fmt.Errorf("%s: Invalid tool, must be empty or mssq-cli or sqlcmd.", c.Name)
	}
	return nil
}

func isUnset(s *string) bool {
	if s == nil {
		return true
	}
	return strings.TrimSpace(*s) == ""
}
