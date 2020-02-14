package command

import (
	"fmt"
	"github.com/RobyFerro/go-web-framework"
	"io"
	"io/ioutil"
	"os"
	"path"
	"syscall"
)

type Install struct {
	Signature   string
	Description string
}

// Command registration
func (c *Install) Register() {
	c.Signature = "install"          // Change command signature
	c.Description = "install Go-Web" // Change command description
}

// Command business logic
func (c *Install) Run(kernel *gwf.HttpKernel, args string, console map[string]interface{}) {
	if err := Dir(gwf.GetDynamicPath("/"), args); err != nil {
		gwf.ProcessError(err)
	}
}

// Dir copies a whole directory recursively
func Dir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcInfo os.FileInfo

	if srcInfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = Dir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = File(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

// File copies a single file from src to dst
func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err := os.Chown(dst, syscall.Getuid(), syscall.Getgid()); err != nil {
		gwf.ProcessError(err)
	}

	return os.Chmod(dst, srcinfo.Mode())
}
