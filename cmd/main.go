package main

import (
	"github.com/charmbracelet/log"
	"github.com/levonxyz/sqlite-operator/config"
	"github.com/levonxyz/sqlite-operator/pkg/kubernetes"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	_, err = kubernetes.New(c.Clientset())
	if err != nil {
		log.Fatal(err)
	}

	log.Info("hello world from sqlite-operator")
}
