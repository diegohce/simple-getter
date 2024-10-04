package simplegetter

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
)

/* func DownloadProcessor(processor ProcessorFile, targetDir string) error {
	opts := []getter.ClientOption{}
	client := &getter.Client{
		Ctx:     context.Background(),
		Src:     processor.URL,
		Dst:     filepath.Join(targetDir, filepath.Base(processor.Filename)),
		Mode:    getter.ClientModeFile,
		Options: opts,
	}

	if err := client.Get(); err != nil {
		return err
	}

	return nil
} */

const (
	ClientModeFile = 0
)

type ClientOption struct{}

type Client struct {
	Ctx     context.Context
	Src     string
	Dst     string
	Mode    int
	Options []ClientOption
}

func (c *Client) Get() error {

	exists := false
	isZero := false

	exists, err := fileExists(c.Dst)
	if err != nil {
		return err
	}

	if exists {
		isZero, err = zeroSize(c.Dst)
		if err != nil {
			return err
		}
	}

	if !exists || (exists && isZero) {
		err = downloadFile(c)
	}

	return err
}

func downloadFile(c *Client) error {

	res, err := http.Get(c.Src)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	f, err := os.OpenFile(c.Dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)

	return err
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return false, err
		} else {
			return false, nil
		}
	}
	return true, nil
}

func zeroSize(path string) (bool, error) {
	fin, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fin.Size() == 0, nil
}
