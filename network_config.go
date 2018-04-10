package main

import (
  "fmt"
  "reflect"
)

type Subnets struct {
  Type string
}

type Params struct {
  BondMode string
}

type Config struct {
  Type              string
  Name              string
  MacAddress        string
  VlanLink          string
  VlanID            int
  Mtu               int
  BondInterfaces    []string
  Params            Params
  BridgeInterfaces  []string
  Address           []string
  Search            []string
  Destination       string
  Gateway           string
  Metric            int
  Subnets           []Subnets
}

type Network struct {
  Version     int
  ConfigName  string
  Config      []Config
}

type NetworkConfig struct {
  Network struct {
	  Version int `yaml:"version"`
	  ConfigName string `yaml:"config_name"`
	  Config  []struct {
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
  } `yaml:"network"`
}

func createNetworkConfigObject(vn VirtualNetwork) []Network {
  //ncs := []Network{}
  //var nc Network

  ncs := make([]map[string]interface)
  nc := make(map[string]interface)

  for i, rgn := range vn.Region {
    //fmt.Println("region index:", i)
    //fmt.Println("region data:", rgn)
    fmt.Println(reflect.TypeOf(i))

    for ii, fab := range rgn.Fabric {
      // There will be one config file per a region
      fmt.Println(reflect.TypeOf(ii))

      //nc = Network{Version: 1, ConfigName: rgn.Name + "-" + fab.Name}
      nc = map[string]string{
        "Version": 1,
        "ConfigName": rgn.Name + "-" + fab.Name,
      }

      //fmt.Println("fabric index:", ii)
      //fmt.Println("fabric data:", fab)

      for iii, net := range fab.Network {
        fmt.Println(reflect.TypeOf(iii), reflect.TypeOf(net))
        //fmt.Println("network index:", iii)
        //fmt.Println("network data:", net)
      }

      //structToYaml(nc)
      mapToYaml(nc)
      ncs = append(ncs, nc)
    }
    fmt.Println()
  }

  fmt.Printf("ncs:\n%#v", ncs)

  return ncs
}

func generateNetConf(ncs []NetworkConfig, hostId int) {
}
