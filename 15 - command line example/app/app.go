package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate vai retornar a aplicação de linha de comando pronta para ser executada
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search IPs and server names at web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs address at web",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Search server name at web",
			Flags:  flags,
			Action: searchHosts,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IPs do host %s\n", host)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchHosts(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Servers do host %s\n", host)
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
