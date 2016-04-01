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

package config

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/naoina/toml"
)

func FromFile(path string) (*Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to load config: %s", err)
	}
	defer configFile.Close()

	buf, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, fmt.Errorf("Unable to read config file: %s", err)
	}
	return FromData(buf)
}

func FromData(data []byte) (*Config, error) {
	var conf Config
	if err := toml.Unmarshal(data, &conf); err != nil {
		return nil, fmt.Errorf("Unable to parse the config: %s", err)
	}
	if err := sanityCheck(&conf); err != nil {
		return nil, fmt.Errorf("Your config file has issues:\n%s", err)
	}
	return &conf, nil
}

func sanityCheck(conf *Config) error {
	var problems []string
	for _, prj := range conf.Projects {
		fi, err := os.Stat(prj.Path)
		if err != nil {
			problems = append(problems, err.Error())
		} else if !fi.IsDir() {
			problems = append(problems, fmt.Sprintf("%s is not a directory", prj.Path))
		}
	}

	if len(problems) == 0 {
		return nil
	}

	return fmt.Errorf("%s", strings.Join(problems, "\n"))
}
