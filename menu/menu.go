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

// Package "menu" provides logic to show menu of choices corresponding to
// connection entries to the user and getting the user's input.
package menu

import (
	"fmt"

	"github.com/aashishkoirala/mssql-launcher/config"
)

var readChoiceFromUser func() (int, error)

func init() {
	readChoiceFromUser = defaultReadChoiceFromUser
}

// ReadChoice takes in a list of Connection instances and renders a menu
// on screen for the user to choose. It returns the chosen connection, or nil
// if nothing was chosen.
func ReadChoice(connectionList *[]config.Connection) (*config.Connection, error) {
	for i, c := range *connectionList {
		fmt.Printf("%d - %s", i+1, c.Name)
		fmt.Println()
	}
	fmt.Println("0 - Quit")
	choice, err := readChoiceFromUser()
	if err != nil {
		return nil, err
	}
	if choice == 0 {
		return nil, nil
	}
	if choice < 1 || choice > len(*connectionList) {
		return nil, fmt.Errorf("Invalid choice: %d.", choice)
	}
	item := (*connectionList)[choice-1]
	return &item, nil
}

func defaultReadChoiceFromUser() (int, error) {
	choice := 0
	num, err := fmt.Scanln(&choice)
	if num != 1 {
		return 0, nil
	}
	if err != nil {
		msg := err.Error()
		if msg == "unexpected newline" || msg == "expected integer" {
			choice = 0
		} else {
			return 0, err
		}
	}
	return choice, nil
}
