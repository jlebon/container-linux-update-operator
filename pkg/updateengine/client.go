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
	"os"
	"time"
	"strconv"

	"github.com/golang/glog"
	"github.com/godbus/dbus"
)

const (
	dbusName          = "org.projectatomic.rpmostree1"
	dbusSysroot       = "/org/projectatomic/rpmostree1/Sysroot"
	dbusIfaceOS       = "org.projectatomic.rpmostree1.OS"
	dbusIfaceSysroot  = "org.projectatomic.rpmostree1.Sysroot"
	statePollInterval = 1 * time.Hour
)

type Client struct {
	conn  *dbus.Conn
	osPath dbus.ObjectPath
}

func New() (*Client, error) {
	c := new(Client)
	var err error

	c.conn, err = dbus.SystemBusPrivate()
	if err != nil {
		return nil, err
	}

	methods := []dbus.Auth{dbus.AuthExternal(strconv.Itoa(os.Getuid()))}
	err = c.conn.Auth(methods)
	if err != nil {
		c.conn.Close()
		return nil, err
	}

	err = c.conn.Hello()
	if err != nil {
		c.conn.Close()
		return nil, err
	}

	/* find booted OS */
	sysroot := c.conn.Object(dbusName, dbus.ObjectPath(dbusSysroot))
	prop, err := sysroot.GetProperty(dbusMember(dbusIfaceSysroot, "Booted"))
	if err != nil {
		return nil, err
	}

	c.osPath = prop.Value().(dbus.ObjectPath)
	glog.Infof("Booted OS path is %v", c.osPath)
	return c, nil
}

// godbus works on interface.member notation
func dbusMember(iface, member string) string {
	return fmt.Sprintf("%s.%s", iface, member)
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// ReceiveStatuses receives signal messages from dbus and sends them as Statuses
// on the rcvr channel, until the stop channel is closed. An attempt is made to
// get the initial status and send it on the rcvr channel before receiving
// starts.
func (c *Client) ReceiveStatuses(rcvr chan Status, stop <-chan struct{}) {
	// if there is an error getting the current status, just log it and
	// move onto the main loop.
	st, err := c.GetStatus()
	if err != nil {
		glog.Warningf("error getting rpm-ostree status: %v", err)
	} else {
		rcvr <- st
	}

	// XXX: We could probably increase the interval even more if we use
	// filesystem notifications instead if we want to narrow the window between
	// deployment staging and reboot.
	t := time.Tick(statePollInterval)
	for {
		select {
		case <-stop:
			return
		case <-t:
			st, err := c.GetStatus()
			if err != nil {
				glog.Warningf("error getting rpm-ostree status: %v", err)
			} else {
				rcvr <- st
			}
		}
	}
}

// GetStatus gets the current status from rpm-ostree
func (c *Client) GetStatus() (Status, error) {
	glog.Info("Getting status from rpm-ostree")

	bootedOS := c.conn.Object(dbusName, c.osPath)
	prop, err := bootedOS.GetProperty(dbusMember(dbusIfaceOS, "CachedUpdate"))
	if err != nil {
		return Status{}, err
	}

	dict := prop.Value().(map[string]dbus.Variant)
	return NewStatus(dict), nil
}
