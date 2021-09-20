package util

import (
	"errors"
	"fmt"

	"io/ioutil"
	"net"
	"os"

	"github.com/fatih/color"
	"github.com/jhoonb/archivex"
)

func GetIPAddr() (string, error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return "", errors.New("util.getIPAddr(): Failed to get net interfaces")
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return "", errors.New("util.getIPAddr(): Failed to get IP addresses of interface")
		}

		// Each interface can have both IPv4 and IPv6
		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			}

			if !ip.IsLoopback() {
				return ip.String(), nil
			}
		}
	}
	return "", errors.New("util.getIPAddr(): No local IP found")

}

func getAllExistingFiles(files []string) []string {
	existingFiles := []string{}
	for _, file := range files {
		_, err := os.Stat(file)
		if err == nil {
			existingFiles = append(existingFiles, file)
		} else {
			bold := color.New(color.Attribute(color.Bold))
			fmt.Printf("%v %v\n", color.YellowString("ignoring"), bold.Sprint(file))
		}
	}
	return existingFiles
}

func shouldZip(validFiles []string) bool {
	zipRequired := len(validFiles) >= 2

	for _, file := range validFiles {
		metadata, _ := os.Stat(file)
		if metadata.IsDir() {
			zipRequired = true
		}
	}

	return zipRequired
}

func zipFiles(files []string) (string, error) {
	zip := new(archivex.ZipFile)
	tmpFile, err := ioutil.TempFile("", "wft")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	zip.Create(tmpFile.Name())
	for _, fileName := range files {
		metadata, err := os.Stat(fileName)
		if err != nil {
			return "", err
		}
		if metadata.IsDir() {
			zip.AddAll(fileName, true)
		} else {
			file, err := os.Open(fileName)
			if err != nil {
				return "", err
			}
			defer file.Close()
			if err := zip.Add(fileName, file, metadata); err != nil {
				return "", err
			}
		}
	}
	if err := zip.Close(); err != nil {
		return "", err
	}
	return zip.Name, nil

}

func GetPayloadFromArgs(fileArgs []string) (string, error) {
	files := getAllExistingFiles(fileArgs)
	zipRequired := shouldZip(files)

	if zipRequired {
		zip, err := zipFiles(files)
		if err != nil {
			return "", err
		}
		return zip, nil
	}

	if len(files) != 0 {
		return fileArgs[0], nil
	}

	return "", nil
}
