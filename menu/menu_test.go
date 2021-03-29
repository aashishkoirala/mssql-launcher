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

package menu

import (
	"fmt"
	"testing"

	"github.com/aashishkoirala/mssql-launcher/testdata"
)

func TestReadChoice_Valid_Works(t *testing.T) {
	readChoiceFromUser = func() (int, error) {
		return 1, nil
	}
	connections := testdata.Connections()
	item, err := ReadChoice(&connections)
	if err != nil {
		t.Fatal(err)
	}
	if item == nil {
		t.Fatal("Item is nil.")
	}
	expected := connections[0]
	if *item != expected {
		t.Fatal("Did not receive expected item.")
	}
}

func TestReadChoice_Invalid_Gives_Error(t *testing.T) {
	readChoiceFromUser = func() (int, error) {
		return 7, nil
	}
	connections := testdata.Connections()
	item, err := ReadChoice(&connections)
	if err == nil {
		t.Fatal("Expected error, didn't get one.")
	}
	if item != nil {
		t.Fatal("Expected nil item, got non-nil.")
	}
	msg := err.Error()
	if msg != "Invalid choice: 7." {
		t.Fatalf("Expected error message did not match actual: %s", msg)
	}
}

func TestReadChoice_Error_Is_Passed_Through(t *testing.T) {
	readChoiceFromUser = func() (int, error) {
		return 0, fmt.Errorf("Test Error")
	}
	connections := testdata.Connections()
	item, err := ReadChoice(&connections)
	if err == nil {
		t.Fatal("Expected error, didn't get one.")
	}
	if item != nil {
		t.Fatal("Expected nil item, got non-nil.")
	}
	msg := err.Error()
	if msg != "Test Error" {
		t.Fatalf("Expected error message did not match actual: %s", msg)
	}
}

func TestReadChoice_Zero_Selects_Nothing(t *testing.T) {
	readChoiceFromUser = func() (int, error) {
		return 0, nil
	}
	connections := testdata.Connections()
	item, err := ReadChoice(&connections)
	if err != nil {
		t.Fatal(err)
	}
	if item != nil {
		t.Fatal("Expected nil item, got non-nil.")
	}
}
