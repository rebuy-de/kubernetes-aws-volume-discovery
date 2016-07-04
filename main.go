package main

import (
	"fmt"
	"os"
	"time"
)

const (
	PAUSE = time.Minute
)

var (
	version = "unknown"
)

func main() {
	for {
		err := run()

		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}

		time.Sleep(PAUSE)
	}
}

func run() error {
	awsVolumes, err := awsGetVolumes()
	if err != nil {
		return err
	}

	for _, awsVol := range awsVolumes {
		kubeVol := aws2kube(awsVol)
		if kubeVol == nil {
			continue
		}

		err := kubeApply(*kubeVol)
		if err != nil {
			return err
		}
	}

	return nil
}
