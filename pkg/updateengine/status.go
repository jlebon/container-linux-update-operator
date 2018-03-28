// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package updateengine

import (
	"fmt"

	"github.com/godbus/dbus"
)

const (
	RpmOstreeUpdateNone             = "RPM_OSTREE_UPDATE_NONE"
	RpmOstreeUpdateAvailable        = "RPM_OSTREE_UPDATE_AVAILABLE"
	RpmOstreeUpdateStaged           = "RPM_OSTREE_UPDATE_STAGED"
)

type Status struct {
	CurrentStatus string
	LastCheckedTime uint64
	// Let's just proxy "version" and "checksum" for now. AFAICT, this is only
	// for the benefit of a sysadmin doing `oc describe node`.
	NewVersion string
	NewChecksum string
}

func NewStatus(cachedUpdate map[string]dbus.Variant) (s Status) {

	if len(cachedUpdate) == 0 {
		return Status{CurrentStatus: RpmOstreeUpdateNone}
	}

	checksum := cachedUpdate["checksum"].Value().(string)
	lastChecked := cachedUpdate["update-timestamp"].Value().(uint64)
	// convert to seconds to be easier to consume
	lastChecked /= 1000000

	versionProp, ok := cachedUpdate["version"]
	var version string
	if ok {
		version = versionProp.Value().(string)
	}

	staged := cachedUpdate["staged"].Value().(bool)
	if !staged {
		return Status{RpmOstreeUpdateAvailable, lastChecked, version, checksum}
	}
	return Status{RpmOstreeUpdateStaged, lastChecked, version, checksum}
}

func (s *Status) String() string {
	return fmt.Sprintf("CurrentStatus=%s LastChecked=%d NewVersion=%s NewChecksum=%s",
	                   s.CurrentStatus, s.LastCheckedTime, s.NewVersion, s.NewChecksum)
}
