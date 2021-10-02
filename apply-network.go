package main

import (
	"fmt"
	"os/exec"
)

func ApplyNetwork(config *NetworkConfig) error {
	if len(config.Addresses) == 0 {
		return nil
	}
	for _, address := range config.Addresses {
		err := ApplyAddress(address)
		if err != nil {
			return err
		}

	}
	return nil
}

func ApplyAddress(address *NetworkAddressConfig) error {
	if address.Interface == "" {
		return fmt.Errorf("field \"Interface\" is empty")
	}
	var cmd *exec.Cmd
	if address.Dhcp {
		cmd = exec.Command("netsh", "interface", "ipv4", "set", "address", address.Interface, "source=dhcp")
	} else {
		cmd = exec.Command(
			"netsh", "interface", "ipv4", "set", "address",
			address.Interface,
			"source=static",
			"address="+address.Address,
			"mask="+address.Netmask,
			"gateway="+address.Gateway,
			"gwmetric=1",
		)
	}
	return cmd.Run()
}
