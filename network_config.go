package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	//"reflect"
	"github.com/davecgh/go-spew/spew"
)

//	type Subnets struct {
//		Type string
//	}
//
//	type Params struct {
//		BondMode string
//	}
//
//	type Config struct {
//		Type             string
//		Name             string
//		MacAddress       string
//		VlanLink         string
//		VlanID           int
//		Mtu              int
//		BondInterfaces   []string
//		Params           Params
//		BridgeInterfaces []string
//		Address          []string
//		Search           []string
//		Destination      string
//		Gateway          string
//		Metric           int
//		Subnets          []Subnets
//	}
//
//	type Network struct {
//		Version    int
//		ConfigName string
//		Config     []Config
//	}

type Network struct {
	Version    int    `yaml:"version"`
	ConfigName string `yaml:"config_name"`
	Config     []struct {
		Type           string   `yaml:"type"`
		Name           string   `yaml:"name,omitempty"`
		MacAddress     string   `yaml:"mac_address,omitempty"`
		VlanLink       string   `yaml:"vlan_link,omitempty"`
		VlanID         int      `yaml:"vlan_id,omitempty"`
		Mtu            int      `yaml:"mtu,omitempty"`
		BondInterfaces []string `yaml:"bond_interfaces,omitempty"`
		Params         struct {
			BondMode string `yaml:"bond-mode"`
		} `yaml:"params,omitempty"`
		BridgeInterfaces []string `yaml:"bridge_interfaces,omitempty"`
		Address          []string `yaml:"address,omitempty"`
		Search           []string `yaml:"search,omitempty"`
		Destination      string   `yaml:"destination,omitempty"`
		Gateway          string   `yaml:"gateway,omitempty"`
		Metric           int      `yaml:"metric,omitempty"`
		Subnets          []struct {
			Type string `yaml:"type"`
		} `yaml:"subnets,omitempty"`
		ID string `yaml:"id,omitempty"`
	} `yaml:"config"`
}

func createNetworkConfigWithMap(vn VirtualNetwork, hostId int) {
	// have problem initializing nested map
	var networks []map[string]interface{}
	network := map[string]interface{}{}
	network["network"] = map[string]interface{}{}

	if vn.Region != nil {
		for _, rgn := range vn.Region {

			if rgn.Fabric != nil {
				for _, fab := range rgn.Fabric {
					// There will be one config file per a region
					for _, net := range fab.Network {
						if net.Fake {
						} else {
						}
					} // net
				} // fab
			}
			networks = append(networks, network)
		} // rgn

	}
	spew.Dump(networks)
	fmt.Println()
	for _, net := range networks {
		mapToYaml(net)
	}
}

func createNetworkConfigWithStruct(vn VirtualNetwork, hostId int) {
	// have problem initializing nested struct
	if vn.Region != nil {
		for _, rgn := range vn.Region {
			if rgn.Fabric != nil {
				for _, fab := range rgn.Fabric {
					// There will be one config file per a region
					for _, net := range fab.Network {
						if net.Fake {
						} else {
						}
					} // net
				} // fab
			}
		} // rgn

	}
}

func createNetworkConfigWithString(vn VirtualNetwork, hostId int) {

	// How to initialize nested struct
	// https://stackoverflow.com/questions/24809235/initialize-a-nested-struct-in-golang
	// https://stackoverflow.com/questions/26866879/initialize-nested-struct-definition-in-golang/26867130
	// https://medium.com/@xcoulon/nested-structs-in-golang-2c750403a007
	// https://gist.github.com/hvoecking/10772475

	fmt.Println()

	var ns string

	if vn.Region != nil {
		for i, rgn := range vn.Region {
			ns = "network:\n"
			ns = ns + "  " + "version: 1\n"
			ns = ns + "  " + "config_name: " + rgn.Name + "\n"

			if rgn.Fabric != nil {
				ns = ns + "  " + "config:\n"
				for ii, fab := range rgn.Fabric {
					// There will be one config file per a region
					for _, net := range fab.Network {
						if net.Fake {
							ns = ns + "    " + "- type: vlan\n"
							ns = ns +
								"    " +
								"  name: ech" +
								strconv.Itoa(ii) +
								"." +
								strconv.Itoa(net.Vid) +
								"\n"
							ns = ns +
								"    " +
								"  vlan_link: ech" +
								strconv.Itoa(ii) +
								"\n"
							ns = ns +
								"    " +
								"  vlan_id: " +
								strconv.Itoa(net.Vid) +
								"\n"
						} else {
							ns = ns + "    " + "- type: physical\n"
							ns = ns + "    " + "  name: ech" + strconv.Itoa(ii) + "\n"
						}
						ns = ns + "    " + "  mtu: " + strconv.Itoa(rgn.Mtu) + "\n"
						ns = ns + "    " + "  subnets:\n"
						ns = ns + "        " + "- type: static\n"
						cidr := strings.Split(net.Cidr, "/")
						octets := strings.Split(cidr[0], ".")
						netid := octets[0] + "." + octets[1] + "." + octets[2]
						ns = ns +
							"        " +
							"  address: " +
							netid +
							"." +
							strconv.Itoa(hostId) +
							"/" +
							cidr[1] +
							"\n"
						ns = ns + "        " + "  gateway: " + netid + ".1\n"

					} // net
				} // fab
			}
			ns = ns + "    " + "- type: nameserver\n"
			ns = ns + "    " + "  address:\n"
			ns = ns + "        " + "- 8.8.8.8\n"
			fmt.Println("===============>> network_config " + strconv.Itoa(i+1))
			fmt.Println(ns)
			ioutil.WriteFile(rgn.Name+".yml", []byte(ns), 0644)
		} // rgn
	}
}
