// This file is part of the minion task manager
// Copyright (C) 2016  Hans Meyer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"github.com/hamster21/minion/config"
)

const (
	gitRevision = ""
	version     = "0.1.0-dev"
)

func main() {
	fmt.Println(licenseHint())

	conf, err := config.FromFile("example-config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Complete config: %v\n", conf)
	fmt.Println("Projects:")
	for _, prj := range conf.Projects {
		fmt.Println(prj.Name, "@", prj.Path)
	}
}

func licenseHint() string {
	return fmt.Sprintf(
		`Minion task manager; Version %s
Copyright (C) 2016 Hans Meyer
This program comes with ABSOLUTELY NO WARRANTY and is free software.
You are welcome to redistribute it under certain conditions.
See https://github.com/HaMster21/minion/blob/%s/LICENSE for details
`,
		version, gitRevision)
}
