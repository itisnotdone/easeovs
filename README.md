# easeovs
Will get you Openvswitch as you define within a YAML file.

## Installation
```bash
go get github.com/itisnotdone/easeovs
```

## Steps to use easeovs
- Create a yaml file and define a network or networks referring 'sample.yml'
- Run 'create' command to create virtual networks and gateways
- Run 'generate' command to create network and device configurations for MAAS

## Requirments

### /etc/network/interfaces
has to contain `source /etc/network/interfaces.d/*.cfg` to be able to recognize more network configurations.

### /etc/sudoers
`Defaults  secure_path` has to contains $GOBIN which is $GOPATH/bin. Following script might help you out.
```bash
sudo sed -i "s \(.*secure_path=.*\)\" \1:$HOME/go/bin\" " /etc/sudoers
```


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


# Make sure that you have following line in your /etc/network/interfaces file.
source /etc/network/interfaces.d/*.cfg

# to build virtual network as defined
$ sudo easeovs create --config $GOPATH/src/github.com/itisnotdone/easeovs/template/single_region.yml

# to destroy the virtual network
$ sudo easeovs destroy --config $GOPATH/src/github.com/itisnotdone/easeovs/template/single_region.yml

# `generate` command will generate network and device configuration for MAAS container and XML network definition for libvirt
easeovs generate --config $GOPATH/src/github.com/itisnotdone/easeovs/template/single_region.yml --host-id 2
ls
cloudinit_net_argn.yml  cloudinit_net_default.yml  virsh_net_argn_f01.xml  virsh_net_default_f01.xml

# to create a lxc container to build a MAAS
gogetit create argn-maas --no-maas -f cloudinit_net_argn.yml --maas-on-lxc

# to define a virtual network for libvirt
virsh net-define virsh_net_argn_f01.xml
virsh net-start argn-f01
virsh net-autostart argn-f01
```

## Reference
- https://github.com/go-yaml/yaml
- https://mengzhuo.github.io/yaml-to-go/
- https://github.com/urfave/cli
- https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
