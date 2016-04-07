package proxy

import (
	"io/ioutil"
)

func tmpFile(content []byte) (string, error) {

	tmpfile, err := ioutil.TempFile("", "tmp")
	if err != nil {
		return "", err
	}

	if _, err := tmpfile.Write(content); err != nil {
		return "", err
	}
	if err := tmpfile.Close(); err != nil {
		return "", err
	}

	return tmpfile.Name(), nil
}
