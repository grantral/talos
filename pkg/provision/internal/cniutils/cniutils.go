// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package cniutils provides helper functions to parse CNI results.
package cniutils

import (
	"fmt"

	types100 "github.com/containernetworking/cni/pkg/types/100"
)

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// FilterBySandbox returns scans the provided list of interfaces and returns two lists: the first
// are a list of interfaces with the provided sandboxID, the second are the other interfaces not
// in that sandboxID.
func FilterBySandbox(sandbox string, ifaces ...*types100.Interface) (in, out []*types100.Interface) {
	for _, iface := range ifaces {
		if iface.Sandbox == sandbox {
			in = append(in, iface)
		} else {
			out = append(out, iface)
		}
	}

	return
}

// IfacesWithName scans the provided list of ifaces and returns the ones with the provided name.
func IfacesWithName(name string, ifaces ...*types100.Interface) []*types100.Interface {
	var foundIfaces []*types100.Interface

	for _, iface := range ifaces {
		if iface.Name == name {
			foundIfaces = append(foundIfaces, iface)
		}
	}

	return foundIfaces
}

// VMTapPair takes a CNI result and returns the vm iface and the tap iface corresponding
// to the provided vmID. See the vmconf package docs for details on the expected
// vm and tap iface configurations.
func VMTapPair(result *types100.Result, vmID string) (vmIface, tapIface *types100.Interface, err error) {
	vmIfaces, otherIfaces := FilterBySandbox(vmID, result.Interfaces...)
	if len(vmIfaces) > 1 {
		return nil, nil, fmt.Errorf(
			"expected to find at most 1 interface in sandbox %q, but instead found %d",
			vmID, len(vmIfaces))
	} else if len(vmIfaces) == 0 {
		return nil, nil, fmt.Errorf("no pseudo-device for %s", vmID)
	}

	vmIface = vmIfaces[0]

	// As specified in the package docstring, the vm interface is given the same name as the
	// corresponding tap device outside the VM. The tap device, however, will be in a sandbox
	// corresponding to a network namespace path.
	tapName := vmIface.Name

	tapIfaces := IfacesWithName(tapName, otherIfaces...)
	if len(tapIfaces) > 1 {
		return nil, nil, fmt.Errorf(
			"expected to find at most 1 interface with name %q, but instead found %d",
			tapName, len(tapIfaces))
	} else if len(tapIfaces) == 0 {
		return nil, nil, fmt.Errorf("device not found: %s", tapName)
	}

	tapIface = tapIfaces[0]

	return vmIface, tapIface, nil
}
