# easeovs
Will get you Openvswitch as you defined within a YAML file.

## Installation
```bash
go get github.com/itisnotdone/easeovs
```

## Steps to use easeovs
- Create a yaml file and define a network or networks referring 'sample.yml'
- Run 'create' command to create virtual networks and gateways
- Run 'generate' command to create network and device configurations for MAAS

## Usage
```bash
easeovs 
NAME:
   easeovs - will help you create virtual networks using Openvswitch

USAGE:
   easeovs [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     clean  clean garbage ports shown in the result of 'ovs-vsctl show'.
              Note that this will only work on veth and vnet type ports.
     create  create networks and gateways as defined.
              Note that this command needs sudo privilege.
     destroy  destroy networks and gateways as defined.
              Note that this command needs sudo privilege.
     generate  generate network configuration for MAAS in cloud-init v1 format
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version


# to create
sudo easeovs create --config src/github.com/itisnotdone/easeovs/sample.yml
ovs-vsctl add-br argn-f01
ovs-vsctl add-br argn-f01-10 argn-f01 10
===============>> network_config for bm-switch
Generating /etc/network/interfaces.d/argn-f01-10.cfg
ifup argn-f01-10
ovs-vsctl add-br argn-f01-11 argn-f01 11
===============>> network_config for ipmi
Generating /etc/network/interfaces.d/argn-f01-11.cfg
ifup argn-f01-11
===============>> network_config for bm-server
Generating /etc/network/interfaces.d/argn-f01.cfg
ifup argn-f01
ovs-vsctl add-br argn-f01-13 argn-f01 13
===============>> network_config for infra-svc
Generating /etc/network/interfaces.d/argn-f01-13.cfg
ifup argn-f01-13
ovs-vsctl add-br argn-f01-14 argn-f01 14
===============>> network_config for k8s
Generating /etc/network/interfaces.d/argn-f01-14.cfg
ifup argn-f01-14
ovs-vsctl add-br argn-f01-18 argn-f01 18
===============>> network_config for os-admin
Generating /etc/network/interfaces.d/argn-f01-18.cfg
ifup argn-f01-18
ovs-vsctl add-br argn-f01-19 argn-f01 19
===============>> network_config for os-public
Generating /etc/network/interfaces.d/argn-f01-19.cfg
ifup argn-f01-19
ovs-vsctl add-br argn-f01-116 argn-f01 116
===============>> network_config for os-floating
Generating /etc/network/interfaces.d/argn-f01-116.cfg
ifup argn-f01-116
ovs-vsctl add-br argn-f01-1100 argn-f01 1100
===============>> network_config for os-tunnel
Generating /etc/network/interfaces.d/argn-f01-1100.cfg
ifup argn-f01-1100
ovs-vsctl add-br argn-f01-1101 argn-f01 1101
===============>> network_config for ceph-cluster
Generating /etc/network/interfaces.d/argn-f01-1101.cfg
ifup argn-f01-1101
ovs-vsctl add-br brgn-f01
ovs-vsctl add-br brgn-f01-164 brgn-f01 164
===============>> network_config for bm-switch
Generating /etc/network/interfaces.d/brgn-f01-164.cfg
ifup brgn-f01-164
ovs-vsctl add-br brgn-f01-165 brgn-f01 165
===============>> network_config for ipmi
Generating /etc/network/interfaces.d/brgn-f01-165.cfg
ifup brgn-f01-165
===============>> network_config for bm-server
Generating /etc/network/interfaces.d/brgn-f01.cfg
ifup brgn-f01
ovs-vsctl add-br brgn-f01-167 brgn-f01 167
===============>> network_config for infra-svc
Generating /etc/network/interfaces.d/brgn-f01-167.cfg
ifup brgn-f01-167
ovs-vsctl add-br brgn-f01-168 brgn-f01 168
===============>> network_config for k8s
Generating /etc/network/interfaces.d/brgn-f01-168.cfg
ifup brgn-f01-168
ovs-vsctl add-br brgn-f01-172 brgn-f01 172
===============>> network_config for os-admin
Generating /etc/network/interfaces.d/brgn-f01-172.cfg
ifup brgn-f01-172
ovs-vsctl add-br brgn-f01-173 brgn-f01 173
===============>> network_config for os-public
Generating /etc/network/interfaces.d/brgn-f01-173.cfg
ifup brgn-f01-173
ovs-vsctl add-br brgn-f01-180 brgn-f01 180
===============>> network_config for os-floating
Generating /etc/network/interfaces.d/brgn-f01-180.cfg
ifup brgn-f01-180
ovs-vsctl add-br brgn-f01-1200 brgn-f01 1200
===============>> network_config for os-tunnel
Generating /etc/network/interfaces.d/brgn-f01-1200.cfg
ifup brgn-f01-1200
ovs-vsctl add-br brgn-f01-1201 brgn-f01 1201
===============>> network_config for ceph-cluster
Generating /etc/network/interfaces.d/brgn-f01-1201.cfg
ifup brgn-f01-1201

# to destroy
sudo easeovs destroy --config ~/go/src/github.com/itisnotdone/easeovs/sample.yml
ifdown argn-f01-10
ovs-vsctl del-br argn-f01-10
Removing /etc/network/interfaces.d/argn-f01-10.cfg
ifdown argn-f01-11
ovs-vsctl del-br argn-f01-11
Removing /etc/network/interfaces.d/argn-f01-11.cfg
ifdown argn-f01-13
ovs-vsctl del-br argn-f01-13
Removing /etc/network/interfaces.d/argn-f01-13.cfg
ifdown argn-f01-14
ovs-vsctl del-br argn-f01-14
Removing /etc/network/interfaces.d/argn-f01-14.cfg
ifdown argn-f01-18
ovs-vsctl del-br argn-f01-18
Removing /etc/network/interfaces.d/argn-f01-18.cfg
ifdown argn-f01-19
ovs-vsctl del-br argn-f01-19
Removing /etc/network/interfaces.d/argn-f01-19.cfg
ifdown argn-f01-116
ovs-vsctl del-br argn-f01-116
Removing /etc/network/interfaces.d/argn-f01-116.cfg
ifdown argn-f01-1100
ovs-vsctl del-br argn-f01-1100
Removing /etc/network/interfaces.d/argn-f01-1100.cfg
ifdown argn-f01-1101
ovs-vsctl del-br argn-f01-1101
Removing /etc/network/interfaces.d/argn-f01-1101.cfg
ifdown argn-f01
ovs-vsctl del-br argn-f01
Removing /etc/network/interfaces.d/argn-f01.cfg
ifdown brgn-f01-164
ovs-vsctl del-br brgn-f01-164
Removing /etc/network/interfaces.d/brgn-f01-164.cfg
ifdown brgn-f01-165
ovs-vsctl del-br brgn-f01-165
Removing /etc/network/interfaces.d/brgn-f01-165.cfg
ifdown brgn-f01-167
ovs-vsctl del-br brgn-f01-167
Removing /etc/network/interfaces.d/brgn-f01-167.cfg
ifdown brgn-f01-168
ovs-vsctl del-br brgn-f01-168
Removing /etc/network/interfaces.d/brgn-f01-168.cfg
ifdown brgn-f01-172
ovs-vsctl del-br brgn-f01-172
Removing /etc/network/interfaces.d/brgn-f01-172.cfg
ifdown brgn-f01-173
ovs-vsctl del-br brgn-f01-173
Removing /etc/network/interfaces.d/brgn-f01-173.cfg
ifdown brgn-f01-180
ovs-vsctl del-br brgn-f01-180
Removing /etc/network/interfaces.d/brgn-f01-180.cfg
ifdown brgn-f01-1200
ovs-vsctl del-br brgn-f01-1200
Removing /etc/network/interfaces.d/brgn-f01-1200.cfg
ifdown brgn-f01-1201
ovs-vsctl del-br brgn-f01-1201
Removing /etc/network/interfaces.d/brgn-f01-1201.cfg
ifdown brgn-f01
ovs-vsctl del-br brgn-f01
Removing /etc/network/interfaces.d/brgn-f01.cfg

easeovs generate --config /home/ubuntu/go/src/github.com/itisnotdone/easeovs/sample.yml --host-id 2
find . -maxdepth 1 -name "*.yml"
./network_argn.yml
./device_argn.yml
./device_brgn.yml
./network_brgn.yml
```

## Reference
- https://github.com/go-yaml/yaml
- https://mengzhuo.github.io/yaml-to-go/
- https://github.com/urfave/cli
- https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
