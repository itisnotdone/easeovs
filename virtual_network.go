package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strconv"
	//"reflect"
)

var (
	bridgeName string
	cmd        string
)

type VirtualNetwork struct {
	Region []struct {
    Name    string `yaml:"name"`
    Domain  string `yaml:"domain"`
    Mtu     int `yaml:"mtu"`
    Fabric []struct {
      Name    string `yaml:"name"`
      Network []struct {
        Name   string `yaml:"name"`
        Fake   bool   `yaml:"fake"`
        Vid    int    `yaml:"vid"`
        Cidr   string `yaml:"cidr"`
        Desc   string `yaml:"desc"`
      } `yaml:"network"`
    } `yaml:"fabric"`
	} `yaml:"region"`
}

func (vn VirtualNetwork) createNetworks() {
	for _, rgn := range vn.Region {
    for _, fab := range rgn.Fabric {
      // Root(Fake) bridge
      parentName := rgn.Name + "-" + fab.Name

      cmd = "sudo ovs-vsctl add-br " + parentName
      if err, _ := runCommand(cmd); err != nil {
        log.Printf("failed to add bridge: %v", err)
      }

      for _, net := range fab.Network {
        bridgeName = parentName + "-" + strconv.Itoa(net.Vid)

        if net.Fake == true {
          cmd = "sudo ovs-vsctl add-br " +
          bridgeName + " " + parentName + " " +
          strconv.Itoa(net.Vid)

          if err, _ := runCommand(cmd); err != nil {
            log.Printf("failed to add bridge: %v", err)
          }
        }
      }
    }
	}
}

func (vn VirtualNetwork) destroyNetworks() {
	for _, rgn := range vn.Region {
    for _, fab := range rgn.Fabric {
      // Root(Fake) bridge
      parentName := rgn.Name + "-" + fab.Name

      for _, net := range fab.Network {
        bridgeName = parentName + "-" + strconv.Itoa(net.Vid)

        if net.Fake == true {
          cmd = "sudo ovs-vsctl del-br " + bridgeName

          if err, _ := runCommand(cmd); err != nil {
            log.Printf("failed to delete bridge: %v", err)
          }
        }
      }

      cmd = "sudo ovs-vsctl del-br " + parentName
      if err, _ := runCommand(cmd); err != nil {
        log.Printf("failed to add bridge: %v", err)
      }
    }
	}
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

