package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
)

func main() {

	var yamlFile string
	var hostId int

	app := cli.NewApp()
	app.Name = "easeovs"
	app.Version = "0.1.0"
	app.Usage = "will help you create virtual networks using Openvswitch"

	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "create networks as defined",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config, c",
					Usage:       "Load configuration from `FILE`",
					Destination: &yamlFile,
				},
			},
			Action: func(c *cli.Context) error {
				vn := createVirtualNetworkObject(yamlFile)
				vn.createNetworks()
				return nil
			},
		},
		{
			Name:  "destroy",
			Usage: "destroy networks as defined",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config, c",
					Usage:       "Load configuration from `FILE`",
					Destination: &yamlFile,
				},
			},
			Action: func(c *cli.Context) error {
				vn := createVirtualNetworkObject(yamlFile)
				vn.destroyNetworks()
				return nil
			},
		},
		{
			Name: "clean",
			Usage: `clean garbage ports shown in the result of 'ovs-vsctl show'.
              Note that this will only work on veth and vnet type ports.`,
			Action: func(c *cli.Context) error {
				clean()
				return nil
			},
		},
		{
			Name: "generate",
			Usage: "generate network configuration in cloud-init v1 format " +
				"for network interfaces of MAAS",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config, c",
					Usage:       "Load configuration from `FILE`",
					Destination: &yamlFile,
				},
				cli.IntFlag{
					Name:        "host-id, i",
					Usage:       "host ID to assign on each network interface",
					Destination: &hostId,
				},
			},
			Action: func(c *cli.Context) error {
				vn := createVirtualNetworkObject(yamlFile)
				createNetworkConfigWithString(vn, hostId)
				//createNetworkConfigWithMap(vn, hostId)
				//createNetworkConfigWithStruct(vn, hostId)
				//mainOfTranslate()
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	clierr := app.Run(os.Args)
	if clierr != nil {
		log.Fatal(clierr)
	}
}
