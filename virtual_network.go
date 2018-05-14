package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	bridgeName string
	cmd        string
	cfgfile    string
	err        error
)

type VirtualNetwork struct {
	Region []struct {
		Name   string `yaml:"name"`
		Domain string `yaml:"domain"`
		Mtu    int    `yaml:"mtu"`
		Fabric []struct {
			Name    string `yaml:"name"`
			Network []struct {
				Name string `yaml:"name"`
				Fake bool   `yaml:"fake"`
				Vid  int    `yaml:"vid"`
				Cidr string `yaml:"cidr"`
				Desc string `yaml:"desc"`
			} `yaml:"network"`
		} `yaml:"fabric"`
	} `yaml:"region"`
}

func (vn VirtualNetwork) createNetworks() {
	var iface, mask, netid, parentName string
	var cidr, octets []string

	for _, rgn := range vn.Region {
		for _, fab := range rgn.Fabric {
			// Root bridge
			parentName = rgn.Name + "-" + fab.Name

			cmd = "ovs-vsctl add-br " + parentName
			if err, _ = runCommand(cmd); err != nil {
				log.Printf("failed to add bridge: %v", err)
			}

			for _, net := range fab.Network {
				iface = ""

				if net.Fake == true {

					bridgeName = parentName + "-" + strconv.Itoa(net.Vid)

					cmd = "ovs-vsctl add-br " +
						bridgeName + " " + parentName + " " +
						strconv.Itoa(net.Vid)
					if err, _ = runCommand(cmd); err != nil {
						log.Printf("failed to add bridge: %v", err)
					} else {
						iface = iface + "auto " + bridgeName + "\n"
						iface = iface + "iface " + bridgeName + " inet static\n"

						cidr = strings.Split(net.Cidr, "/")
						octets = strings.Split(cidr[0], ".")
						mask = cidr[1]
						netid = octets[0] + "." + octets[1] + "." + octets[2]

						iface = iface + "    address " + netid + ".1/" + mask + "\n"
						iface = iface + "    mtu " + strconv.Itoa(rgn.Mtu) + "\n"

						cfgfile = "/etc/network/interfaces.d/" + bridgeName + ".cfg"
					}
				} else {
					bridgeName = parentName
					iface = iface + "auto " + bridgeName + "\n"
					iface = iface + "iface " + bridgeName + " inet static\n"

					cidr = strings.Split(net.Cidr, "/")
					octets = strings.Split(cidr[0], ".")
					mask = cidr[1]
					netid = octets[0] + "." + octets[1] + "." + octets[2]

					iface = iface + "    address " + netid + ".1/" + mask + "\n"
					iface = iface + "    mtu " + strconv.Itoa(rgn.Mtu) + "\n"

					cfgfile = "/etc/network/interfaces.d/" + bridgeName + ".cfg"
				}

				fmt.Println("===============>> network_config for " + net.Name)

				// generate network interface files
				fmt.Println("Generating " + cfgfile)
				err = ioutil.WriteFile(cfgfile, []byte(iface), 0644)
				if err != nil {
					panic(err)
				}

				cmd = "ifup " + bridgeName
				if err, _ = runCommand(cmd); err != nil {
					log.Printf("failed to make the bridge up: %v", err)
				}

			} // net
		} // fab
	} // rgn
}

func (vn VirtualNetwork) destroyNetworks() {
	for _, rgn := range vn.Region {
		for _, fab := range rgn.Fabric {
			// Root(Fake) bridge
			parentName := rgn.Name + "-" + fab.Name

			for _, net := range fab.Network {

				if net.Fake == true {
					bridgeName = parentName + "-" + strconv.Itoa(net.Vid)
					cfgfile = "/etc/network/interfaces.d/" + bridgeName + ".cfg"

					cmd = "ifdown " + bridgeName
					if err, _ = runCommand(cmd); err != nil {
						log.Printf("failed to make the bridge down: %v", err)
					}

					cmd = "ovs-vsctl del-br " + bridgeName
					if err, _ = runCommand(cmd); err != nil {
						log.Printf("failed to delete bridge: %v", err)
					}

					fmt.Println("Removing " + cfgfile)
					err = os.Remove(cfgfile)
					if err != nil {
						panic(err)
					}

				}

			} // net

			bridgeName = parentName
			cfgfile = "/etc/network/interfaces.d/" + bridgeName + ".cfg"

			cmd = "ifdown " + bridgeName
			if err, _ = runCommand(cmd); err != nil {
				log.Printf("failed to make the bridge down: %v", err)
			}

			cmd = "ovs-vsctl del-br " + bridgeName
			if err, _ = runCommand(cmd); err != nil {
				log.Printf("failed to delete bridge: %v", err)
			}

			fmt.Println("Removing " + cfgfile)
			err = os.Remove(cfgfile)
			if err != nil {
				panic(err)
			}
		} // fab
	} // rgn
}

// Create VirtualNetwork object
func createVirtualNetworkObject(yamlFile string) VirtualNetwork {
	yamlData, readerr := ioutil.ReadFile(yamlFile)

	if readerr != nil {
		log.Printf("Failed to read from YAML file #%v ", readerr)
	}

	vn := VirtualNetwork{}
	vnerr := yaml.Unmarshal([]byte(yamlData), &vn)

	if vnerr != nil {
		log.Printf("error: %v", vnerr)
	}
	return vn
}
