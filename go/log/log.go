package log

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

func watchNwnxeeLog() {
	t, err := tail.TailFile("/logs/nwnx.txt", tail.Config{Follow: true})
	for line := range t.Lines {
		// this is where the nwnx log parsing will go.

		//
		log.WithFields(log.Fields{"Line": line.Text, "Time": line.Time}).Info("Order:Log:Nwnxee:")
		f, _ := os.OpenFile("/logs/nwnx.txt", os.O_RDWR|os.O_CREATE, 0666)
		defer f.Close()
		popLine(f)
	}

	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

// InitLog func
func InitLog() {
	// app started
	log.WithFields(log.Fields{"Started": 1}).Info("Order:Log")
	go watchNwnxeeLog()
}
