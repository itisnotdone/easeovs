package main

import (
	"fmt"
	"strings"
	"regexp"
	"os"
)

//type Openvswitch struct {
//	Bridge []struct {
//	  name string
//	  fake bool
//	  Port []struct {
//	    name  string
//	    vid   int
//    }
//  }
//}

type Port struct {
	name  string
	vid   int
  state string
}

type Bridge struct {
	name string
	fake bool
	port []Port
}

type Openvswitch struct {
	bridge []Bridge
}

func clean() {
  // First thing first, check if there is any port in error state
  _, out := runCommand("sudo ovs-vsctl show | grep error")

  if len(out) == 0 {
    fmt.Println("There is no port in error state.")
    os.Exit(0)
  }

  ports_in_error_state_in_ovs := strings.Split(strings.TrimSpace(out), "\n")

  for i, p := range ports_in_error_state_in_ovs {
    ports_in_error_state_in_ovs[i] = strings.TrimSpace(p)
  }

  for i, p := range ports_in_error_state_in_ovs {
    ports_in_error_state_in_ovs[i] = strings.Split(p, " ")[6]
  }

  //fmt.Println("ports_in_error_state_in_ovs: ", ports_in_error_state_in_ovs)
  //fmt.Println()

  _, out = runCommand("sudo ovs-vsctl list-br")
  br_list := strings.Split(strings.TrimSpace(out), "\n")

  //fmt.Println("br_list: ", br_list)
  //fmt.Println()

  matching_word := "veth|vnet"

  _, out = runCommand("ip link | egrep '" + matching_word + "'")
  var ifaces_being_used []string
  for _, v := range strings.Split(strings.TrimSpace(out), "\n") {
    ifaces_being_used = append(
      ifaces_being_used,
      strings.Split(strings.Split(v, " ")[1], "@")[0],
    )
  }

  //fmt.Println("ifaces_being_used: ", ifaces_being_used)
  //fmt.Println()

  ports_listed_in_ovs := Openvswitch{}
  bridge := Bridge{}
  port := Port{}

  var all_ports_in_ovs []string

  ports_listed_in_ovs.bridge = make([]Bridge, 0)

  for i, br := range br_list {
    _, out = runCommand("sudo ovs-vsctl list-ports " + br)

    bridge = Bridge{name: br}

    ports_listed_in_ovs.bridge = append(ports_listed_in_ovs.bridge, bridge)
    ports_listed_in_ovs.bridge[i].port = make([]Port, 0)

    for _, p := range strings.Split(strings.TrimSpace(out), "\n") {

      port = Port{name: p}
      ports_listed_in_ovs.bridge[i].port = append(ports_listed_in_ovs.bridge[i].port, port)
      matched, _ := regexp.MatchString(matching_word, port.name)
      if matched {
        all_ports_in_ovs = append(all_ports_in_ovs, port.name)
      }
    }
  }

  //fmt.Println("ports_listed_in_ovs: ", ports_listed_in_ovs)
  //fmt.Println()
  //fmt.Println("all_ports_in_ovs: ", all_ports_in_ovs)
  //fmt.Println()

  diffInOVS := difference(all_ports_in_ovs, ports_in_error_state_in_ovs)
  //fmt.Println("diffInOVS:", diffInOVS)
  //fmt.Println()

  diffBetweenIfaceAndDiffInOVS := difference(ifaces_being_used, diffInOVS)
  //fmt.Println("diffBetweenIfaceAndDiffInOVS:", diffBetweenIfaceAndDiffInOVS)
  //fmt.Println()

  // Update port state with 'ports_in_error_state_in_ovs'
  for bri, br := range ports_listed_in_ovs.bridge {
    for pi, p := range br.port {
      matched, _ := regexp.MatchString(matching_word, p.name)
      if matched {
        if stringInSlice(p.name, ports_in_error_state_in_ovs) {
          ports_listed_in_ovs.bridge[bri].port[pi].state = "error"
        }
      }
    }
  }

  for _, br := range ports_listed_in_ovs.bridge {
    for _, p := range br.port {
        fmt.Println(br.name, p.name, p.state)
    }
  }

  // Delete ports in error state
  if len(diffBetweenIfaceAndDiffInOVS) == 0 {
    for _, br := range ports_listed_in_ovs.bridge {
      for _, p := range br.port {
        matched, _ := regexp.MatchString(matching_word, p.name)
        if matched {
          if p.state == "error" {
            runCommand("sudo ovs-vsctl del-port " + br.name + " " + p.name)
          }
        }
      }
    }
  }
}
