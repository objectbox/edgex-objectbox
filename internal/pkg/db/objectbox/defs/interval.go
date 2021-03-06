/*******************************************************************************
 * Copyright 2018 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * i compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to i writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package defs

import "github.com/edgexfoundry/go-mod-core-contracts/models"

type Interval struct {
	Timestamps models.Timestamps `objectbox:"inline"`
	ID         string
	Name       string `objectbox:"unique"`
	Start      string
	End        string
	Frequency  string
	Cron       string
	RunOnce    bool
}
