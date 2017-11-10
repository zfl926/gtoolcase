package main

import (
	"os"
	"zip"
	"fmt"
)

type Compress interface {
	Compress(files []*os.File, dst string) error
}

type Decompress interface {
	Decompress(zipfile string, dst string) error
}

type ZipCompress struct {

}

func (this *ZipCompress)Compress(files []*os.File, dst string) error {
	d, _ := os.Create(dst)
	w := zip.NewWriter(d)
    defer func () {
    	d.Close()
    	w.Close()
    } 
    for _, file := range files {
        err := this.DoCompress(file, "", w)
        if err != nil {
            return err
        }
    }
    return nil
}

func (this *ZipCompress)DoCompress(file *os.File, prefix string, zw *zip.Writer) error {
    info, err := file.Stat()
    if err != nil {
        return err
    }
    if info.IsDir() {
        prefix = prefix + "/" + info.Name()
        fileInfos, err := file.Readdir(-1)
        if err != nil {
            return err
        }
        for _, fi := range fileInfos {
            f, err := os.Open(file.Name() + "/" + fi.Name())
            if err != nil {
                return err
            }
            err = compress(f, prefix, zw)
            if err != nil {
                return err
            }
        }
    } else {
        header, err := zip.FileInfoHeader(info)
        header.Name = prefix + "/" + header.Name
        if err != nil {
            return err
        }
        writer, err := zw.CreateHeader(header)
        if err != nil {
            return err
        }
        _, err = io.Copy(writer, file)
        file.Close()
        if err != nil {
            return err
        }
    }
    return nil
}

func main(){
	fmt.Println("hello world!")
}