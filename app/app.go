package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Vai retornar a app de linha de comando pronta para ser executada
func Gerar() *cli.App { // G maiusculo para outros pacotes usarem
	app := cli.NewApp()
	app.Name = "App de linha de comando"
	app.Usage = "Busca de IP e Nome de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com", // Padrão, caso nada seja escrita na linha de comando
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca Ip de algum endereço na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "server",
			Usage: "Busca nome de servers de algum endereço na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servidores {
		fmt.Println(server.Host)
	}

}