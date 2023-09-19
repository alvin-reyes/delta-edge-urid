package cmd

import (
	"context"
	"fmt"
	"github.com/alvin-reyes/edge-urid/api"
	"github.com/alvin-reyes/edge-urid/config"
	"github.com/alvin-reyes/edge-urid/core"
	"github.com/alvin-reyes/edge-urid/jobs"
	"github.com/alvin-reyes/edge-urid/utils"
	"github.com/urfave/cli/v2"
	"runtime"
	"strconv"
)

func DaemonCmd(cfg *config.EdgeConfig) []*cli.Command {
	// add a command to run API node
	var daemonCommands []*cli.Command

	daemonCmd := &cli.Command{
		Name:  "daemon",
		Usage: "EdgeURID gateway daemon that allows users to upload and download data to/from the Filecoin network.",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "repo",
			},
			&cli.StringFlag{
				Name: "host-ip",
			},
			&cli.StringFlag{
				Name: "port",
			},
		},

		Action: func(c *cli.Context) error {

			fmt.Println("OS:", runtime.GOOS)
			fmt.Println("Architecture:", runtime.GOARCH)
			fmt.Println("Hostname:", core.GetHostname())

			var ip string
			if c.String("host-ip") != "" {
				ip = c.String("host-ip")
			} else {
				publicIp, err := core.GetPublicIP()
				if err != nil {
					fmt.Println("Error getting public IP:", err)
				}
				ip = publicIp
			}

			fmt.Println("Public IP:", ip)
			fmt.Println(utils.Blue + "Starting Edge daemon..." + utils.Reset)

			repo := c.String("repo")
			port := c.String("port")
			if repo != "" {
				cfg.Node.Repo = repo
			}
			if port != "" {
				portInt, err := strconv.Atoi(port)
				if err != nil {
					fmt.Println("Error parsing port:", err)
				}
				cfg.Node.Port = portInt
			}

			fmt.Println(utils.Blue + "Setting up the Edge node... " + utils.Reset)
			ln, err := core.NewEdgeNode(context.Background(), *cfg)
			if err != nil {
				return err
			}
			fmt.Println(utils.Blue + "Setting up the Edge node... Done" + utils.Reset)

			core.ScanHostComputeResources(ln, cfg.Node.Repo)
			//	launch the jobs
			go rerunBucketCarGen(ln)

			// launch the API node
			fmt.Printf(`
 _______    ________   ________   _______                    ___  ___   ________     
|\  ___ \  |\   ___ \ |\   ____\ |\  ___ \                  |\  \|\  \ |\   __  \    
\ \   __/| \ \  \_|\ \\ \  \___| \ \   __/|    ____________ \ \  \\\  \\ \  \|\  \   
 \ \  \_|/__\ \  \ \\ \\ \  \  ___\ \  \_|/__ |\____________\\ \  \\\  \\ \   _  _\  
  \ \  \_|\ \\ \  \_\\ \\ \  \|\  \\ \  \_|\ \\|____________| \ \  \\\  \\ \  \\  \| 
   \ \_______\\ \_______\\ \_______\\ \_______\                \ \_______\\ \__\\ _\ 
    \|_______| \|_______| \|_______| \|_______|                 \|_______| \|__|\|__|
`)

			// default tagging.
			api.GetDefaultTagPolicy(ln)
			fmt.Println("Starting API server")
			api.InitializeEchoRouterConfig(ln)
			api.LoopForever()

			return nil
		},
	}

	// add commands.
	daemonCommands = append(daemonCommands, daemonCmd)

	return daemonCommands

}

func rerunBucketCarGen(ln *core.LightNode) {

	// query db for all "processing" jobs and retry
	var buckets []core.Bucket
	ln.DB.Model(&core.Bucket{}).Where("status = ?", "processing").Find(&buckets)
	dispatcher := jobs.CreateNewDispatcher()
	var numJobs int
	for _, bucket := range buckets {
		numJobs += 1
		dispatcher.AddJob(jobs.NewBucketCarGenerator(ln, bucket))
	}
	dispatcher.Start(numJobs)
}
