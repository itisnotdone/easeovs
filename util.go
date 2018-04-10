package main

import (
	"fmt"
	"log"
	"os/exec"
	"bytes"
	"github.com/fatih/color"
  "reflect"
  "gopkg.in/yaml.v2"
)

// Run bash command
func runCommand(cmdstr string) (error, string) {
	color.Cyan(cmdstr)
	cmd := exec.Command("bash", "-c", cmdstr)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
    log.Printf("Failed running \"%s\" with %s\n", cmdstr, err)
    outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
    fmt.Printf("stdout:\n%s\nstderr:\n%s\n", outStr, errStr)
	}

  return err, string(stdout.Bytes())
}

// Return a slice with the differenct elements between two slices
func difference(slice1 []string, slice2 []string) ([]string) {
  diffStr := []string{}
  m :=  map[string]int{}

  for _, s1Val := range slice1 {
    m[s1Val] = 1
  }
  for _, s2Val := range slice2 {
    m[s2Val] = m[s2Val] + 1
  }

  for mKey, mVal := range m {
    if mVal==1 {
      diffStr = append(diffStr, mKey)
    }
  }

  return diffStr
}

// Check if a string is in a slice as an element
func stringInSlice(a string, list []string) bool {
  for _, v := range list {
    //fmt.Println("compare:", a, v)
    if v == a {
      return true
    }
  }
  return false
}

func structToYaml(data Network) {
  d, err := yaml.Marshal(data)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Marshaled nc, d is " + string(reflect.TypeOf(d).String()))
  fmt.Printf("--- nc dump:\n%s\n\n", string(d))

  // using map
  m := make(map[interface{}]interface{})

  err = yaml.Unmarshal([]byte(d), &m)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Unmarshaled data, m is " + string(reflect.TypeOf(m).String()))
  fmt.Printf("--- m:\n%+v\n\n", m)

  dd, err := yaml.Marshal(&m)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Marshaled m, dd is " + string(reflect.TypeOf(dd).String()))
  fmt.Printf("--- m dump:\n%s\n\n", string(dd))
  fmt.Println(data)
}

func mapToYaml(map[interface{}]interface{}) {
  d, err := yaml.Marshal(data)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Marshaled nc, d is " + string(reflect.TypeOf(d).String()))
  fmt.Printf("--- nc dump:\n%s\n\n", string(d))

  // using map
  m := make(map[interface{}]interface{})

  err = yaml.Unmarshal([]byte(d), &m)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Unmarshaled data, m is " + string(reflect.TypeOf(m).String()))
  fmt.Printf("--- m:\n%+v\n\n", m)

  dd, err := yaml.Marshal(&m)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Marshaled m, dd is " + string(reflect.TypeOf(dd).String()))
  fmt.Printf("--- m dump:\n%s\n\n", string(dd))
  fmt.Println(data)
}
