package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"li.lan/labs/plasma/lnrpc"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("rpcserver", "localhost:10000", "The server address in the format of host:port")
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "[lncli] %v\n", err)
	os.Exit(1)
}

func getClient(ctx *cli.Context) lnrpc.LightningClient {
	conn := getClientConn(ctx)
	return lnrpc.NewLightningClient(conn)
}

func getClientConn(ctx *cli.Context) *grpc.ClientConn {
	// TODO(roasbeef): macaroon based auth
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial(ctx.GlobalString("rpcserver"), opts...)
	if err != nil {
		fatal(err)
	}

	return conn
}

func main() {
	app := cli.NewApp()
	app.Name = "lncli"
	app.Version = "0.1"
	app.Usage = "control plane for your LN daemon"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rpcserver",
			Value: "localhost:10000",
			Usage: "host:port of plasma daemon",
		},
	}
	app.Commands = []cli.Command{
		NewAddressCommand,
		SendManyCommand,
		ShellCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fatal(err)
	}

	// ctx := context.Background()
}
