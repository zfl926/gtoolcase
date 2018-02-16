package main

import (
	"os"
    "io"
	"archive/zip"
	//"flag"
    //"io/ioutil"
    "fmt"
    "regexp"
)

type Compress interface {
	Compress(files []*os.File, dst string) error
}

type Decompress interface {
	Decompress(zipfile string, dst string) error
}

type ZipCompress struct {
    IgnoreFiles            []string            // the file to ignore， use semicoin to sepit multi files
    IgnoreRegular            string            // it suporrt regular express to ignore file
    Regular           *regexp.Regexp
}

type ZipDecompress struct {

}

// compress
func (this *ZipCompress)Compress(files []*os.File, dst string) error {
	d, _ := os.Create(dst)
	w := zip.NewWriter(d)
    defer func () {
        w.Close()
    	d.Close()
    }()
    if this.IgnoreRegular != "" {
        r, e := regexp.Compile(this.IgnoreRegular)
        if e != nil {
            return e
        }
        this.Regular = r
    }
    for _, file := range files {
        err := this.addToZip(file, "", w)
        if err != nil {
            return err
        }
    }
    return nil
}

func (this *ZipCompress)addToZip(file *os.File, prefix string, zw *zip.Writer) error {
    info, err := file.Stat()
    if err != nil {
        return err
    }
    defer func(){
        file.Close()
    }()
    if info.IsDir() {
        if prefix == "" {
           prefix = info.Name() 
        } else {
            prefix = prefix + "/" + info.Name()
        }
        fileInfos, err := file.Readdir(-1)
        if err != nil {
            return err
        }
        for _, fi := range fileInfos {
            f, err := os.Open(file.Name() + "/" + fi.Name())
            if err != nil {
                return err
            }
            err = this.addToZip(f, prefix, zw)
            if err != nil {
                return err
            }
        }
    } else {
        // check ignore file
        if this.IgnoreFiles != nil && len(this.IgnoreFiles) > 0 {
            for _,fileName := range this.IgnoreFiles {
                if fileName == info.Name() {
                    return nil
                }
            }
        }
        // regular express
        if this.Regular != nil {
            if this.Regular.MatchString(info.Name()) {
                // skip this file
                return nil
            }
        }
        header, err := zip.FileInfoHeader(info)
        header.Name = prefix + "/" + header.Name
        writer, err := zw.CreateHeader(header)
        if err != nil {
            return err
        }
        _, err = io.Copy(writer, file)
        if err != nil {
            return err
        }
    }
    return nil
}
// decompress
func (this *ZipDecompress) Decompress(zipfile string, dst string) error{
    r, err := zip.OpenReader(zipfile)
    if err != nil {
        return err
    }
    defer r.Close()
    for _,f := range r.File {
        fmt.Printf("file name %s:\n", f.Name)
        rc, err := f.Open()
        if err != nil {
            return err
        }
        // _, err = io.CopyN(os.Stdout, rc, int64(f.UncompressedSize64))
        // if err != nil {
        //     return err
        // }
        rc.Close()
        //fmt.Println()
    }
    return nil
}
////////////////////////////////////////////////////////////
// public function


func main(){
	// files := flag.String("files",".","zip files")
 //    zipfile := flag.String("zipfile",".","zip file name")

 //    flag.Parse()

    // fmt.Printf(" cmd   : %s\n", flag.Arg(0))
    // fmt.Printf(" flies : %s\n", *files)
    // fmt.Printf(" zipfile : %s\n", *zipfile)

    var folder string = "./test"
    var output string = "test.zip"
    file, err := os.Open(folder)
    if err != nil {
        fmt.Println(err)    
        return
    }
    var zippedFiles []*os.File = []*os.File{file}
    compress := &ZipCompress{}
    compress.IgnoreFiles = []string{"ignore.txt", "text.pg"}
    compress.IgnoreRegular = `(?i:^hello).*Go`
    e := compress.Compress(zippedFiles, output)
    if e != nil {
        fmt.Println(e)
    }
    // decompress := &ZipDecompress{}
    // decompress.Decompress("./test.zip", "")
}