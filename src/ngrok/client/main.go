package client

import (
	"fmt"
	"github.com/inconshreveable/mousetrap"
	"math/rand"
	"ngrok/log"
	"ngrok/util"
	"os"
	"runtime"
	"time"
)

func init() {
	if runtime.GOOS == "windows" {
		if mousetrap.StartedByExplorer() {
			fmt.Println("不要双击执行文件!")
			fmt.Println("你得按照教程来启动，在快速上手文档就有!")
			time.Sleep(5 * time.Second)
			os.Exit(1)
		}
	}
}

func Main() {
	// parse options
	opts, err := ParseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set up logging
	log.LogTo(opts.logto, opts.loglevel)

	// read configuration file
	config, err := LoadConfiguration(opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// seed random number generator
	seed, err := util.RandomSeed()
	if err != nil {
		fmt.Printf("Couldn't securely seed the random number generator!")
		os.Exit(1)
	}
	rand.Seed(seed)

	NewController().Run(config)
}
