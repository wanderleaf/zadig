/*
Copyright 2021 The KodeRover Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"strings"

	"github.com/spf13/viper"

	// init the config first
	_ "github.com/koderover/zadig/pkg/config"
	"github.com/koderover/zadig/pkg/setting"
)

func AslanAPI() string {
	return viper.GetString(setting.ENVAslanAPI)
}

func AslanxAPI() string {
	return viper.GetString(setting.ENVAslanxAPI)
}

func RootToken() string {
	return viper.GetString(setting.ENVRootToken)
}

func CollieAPI() string {
	return viper.GetString(setting.ENVCollieAPIAddress)
}

func CollieToken() string {
	return RootToken()
}

func NsqLookupAddrs() []string {
	return strings.Split(viper.GetString(setting.ENVNsqLookupAddrs), ",")
}
