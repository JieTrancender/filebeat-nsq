package main

import (
	"os"

	"github.com/elastic/beats/v7/filebeat/cmd"
	inputs "github.com/elastic/beats/v7/filebeat/input/default-inputs"

	_ "github.com/JieTrancender/filebeat-nsq/output/nsq"
)

func main() {
	if err := cmd.Filebeat(inputs.Init).Execute(); err != nil {
		os.Exit(1)
	}
}
