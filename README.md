# easeovs
Will get you Openvswitch as you defined within a YAML file.

## Installation
```bash
go get github.com/itisnotdone/easeovs
```

## Usage
```bash
easeovs
NAME:
   easeovs - will help you create virtual networks using Openvswitch

USAGE:
   easeovs [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     clean    clean garbage ports shown in the result of 'ovs-vsctl show'.
              Note that this will only work on veth and vnet type ports.
     create   create networks as defined
     destroy  destroy networks as defined
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version


# to create
easeovs create --config src/github.com/itisnotdone/easeovs/sample.yml
sudo ovs-vsctl add-br argn-f01
sudo ovs-vsctl add-br argn-f01-10 argn-f01 10
sudo ovs-vsctl add-br argn-f01-11 argn-f01 11
sudo ovs-vsctl add-br argn-f01-13 argn-f01 13
sudo ovs-vsctl add-br argn-f01-16 argn-f01 16
sudo ovs-vsctl add-br argn-f01-20 argn-f01 20
sudo ovs-vsctl add-br argn-f01-21 argn-f01 21
sudo ovs-vsctl add-br argn-f01-24 argn-f01 24
sudo ovs-vsctl add-br argn-f01-200 argn-f01 200
sudo ovs-vsctl add-br argn-f01-201 argn-f01 201
sudo ovs-vsctl add-br brgn-f01
sudo ovs-vsctl add-br brgn-f01-10 brgn-f01 10
sudo ovs-vsctl add-br brgn-f01-11 brgn-f01 11
sudo ovs-vsctl add-br brgn-f01-13 brgn-f01 13
sudo ovs-vsctl add-br brgn-f01-16 brgn-f01 16
sudo ovs-vsctl add-br brgn-f01-20 brgn-f01 20
sudo ovs-vsctl add-br brgn-f01-21 brgn-f01 21
sudo ovs-vsctl add-br brgn-f01-24 brgn-f01 24
sudo ovs-vsctl add-br brgn-f01-200 brgn-f01 200
sudo ovs-vsctl add-br brgn-f01-201 brgn-f01 201

# to destroy
easeovs destroy --config ~/go/src/github.com/itisnotdone/easeovs/sample.yml
sudo ovs-vsctl del-br argn-f01-10
sudo ovs-vsctl del-br argn-f01-11
sudo ovs-vsctl del-br argn-f01-13
sudo ovs-vsctl del-br argn-f01-16
sudo ovs-vsctl del-br argn-f01-20
sudo ovs-vsctl del-br argn-f01-21
sudo ovs-vsctl del-br argn-f01-24
sudo ovs-vsctl del-br argn-f01-200
sudo ovs-vsctl del-br argn-f01-201
sudo ovs-vsctl del-br argn-f01
sudo ovs-vsctl del-br brgn-f01-10
sudo ovs-vsctl del-br brgn-f01-11
sudo ovs-vsctl del-br brgn-f01-13
sudo ovs-vsctl del-br brgn-f01-16
sudo ovs-vsctl del-br brgn-f01-20
sudo ovs-vsctl del-br brgn-f01-21
sudo ovs-vsctl del-br brgn-f01-24
sudo ovs-vsctl del-br brgn-f01-200
sudo ovs-vsctl del-br brgn-f01-201
sudo ovs-vsctl del-br brgn-f01
```

## Reference
- https://github.com/go-yaml/yaml
- https://mengzhuo.github.io/yaml-to-go/
- https://github.com/urfave/cli
- https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
