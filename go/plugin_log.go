package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/hpcloud/tail"
	log "github.com/sirupsen/logrus"
)

func popLine(f *os.File) ([]byte, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}
	line, err := buf.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	nw, err := io.Copy(f, buf)
	if err != nil {
		return nil, err
	}
	err = f.Truncate(nw)
	if err != nil {
		return nil, err
	}
	err = f.Sync()
	if err != nil {
		return nil, err
	}

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	return []byte(line), nil
}

func initLog() {
	// app started
	log.WithFields(log.Fields{"Started": 1}).Info("Order:Log")

	go watchNwnxeeLog()
}

func watchNwnxeeLog() {
	fname := `/logs/nwnx.txt`
	t, err := tail.TailFile("/logs/nwnx.txt", tail.Config{Follow: true})
	for line := range t.Lines {
		log.WithFields(log.Fields{"Line": line.Text}).Info("Order:Log:Nwnxee:")
		f, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
		defer f.Close()
		line, _ := popLine(f)
		fmt.Println("pop:", string(line))
	}

	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}