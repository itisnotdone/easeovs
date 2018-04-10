package main

import (
  "fmt"
  "log"
  "gopkg.in/yaml.v2"
  "reflect"
	"github.com/fatih/color"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
  A string
  B struct {
    RenamedC int   `yaml:"c"`
    D        []int `yaml:",flow"`
  }
  Fabric []struct {
    Network []struct {
    } `yaml:"network"`
  } `yaml:"fabric"`
}

func blah() {
  // using T as struct
  t := T{}

  fmt.Print("data:")
  fmt.Println(data)
  color.Cyan("data is " + string(reflect.TypeOf(data).String()))
  err := yaml.Unmarshal([]byte(data), &t)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Unmarshaled data, t is " + string(reflect.TypeOf(t).String()))
  fmt.Printf("--- t:\n%+v\n\n", t)

  d, err := yaml.Marshal(&t)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Marshaled t, d is " + string(reflect.TypeOf(d).String()))
  fmt.Printf("--- t dump:\n%s\n\n", string(d))

  // using map
  m := make(map[interface{}]interface{})

  err = yaml.Unmarshal([]byte(data), &m)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Unmarshaled data, m is " + string(reflect.TypeOf(m).String()))
  fmt.Printf("--- m:\n%+v\n\n", m)

  d, err = yaml.Marshal(&m)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Marshaled m, d is " + string(reflect.TypeOf(d).String()))
  fmt.Printf("--- m dump:\n%s\n\n", string(d))

  // using d, unmarshal and marshal as struct
  tt := T{}
  err = yaml.Unmarshal(d, &tt)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Unmarshaled data, tt is " + string(reflect.TypeOf(tt).String()))
  fmt.Printf("--- tt:\n%+v\n\n", tt)

  dd, err := yaml.Marshal(&tt)
  if err != nil {
    log.Fatalf("error: %v", err)
  }
  color.Cyan("Marshaled tt, dd is " + string(reflect.TypeOf(dd).String()))
  fmt.Printf("--- tt dump:\n%s\n\n", string(dd))
}
