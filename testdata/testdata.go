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

// Package "testdata" provides test data for tests.
package testdata

import (
	"github.com/aashishkoirala/mssql-launcher/config"
)

// Connections returns dummy connection data.
func Connections() []config.Connection {
	c := [3]config.Connection{
		{Name: "Item 1", Server: "Server1", Database: "Database1"},
		{Name: "Item 2", Server: "Server2", Database: "Database2"},
		{Name: "Item 3", Server: "Server3", Database: "Database3"},
	}
	return c[:]
}
