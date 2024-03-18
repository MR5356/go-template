package main

import (
	"github.com/MR5356/go-template/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.NewApplication().Execute(); err != nil {
		logrus.Fatal(err)
	}
}
