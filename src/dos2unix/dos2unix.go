package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "io"
    "strings"
)

type Cmd struct {
	FileNames            []string               `file names`
	Format               int                    `format 1 dos2unix 2 unix2dos`
}

func dos2unix(fileName string) int {
	f, err := os.Open(fileName)
    if err != nil {
        return -1
    }
    defer f.Close()
    rd := bufio.NewReader(f)
    wd := bufio.NewWriter(f)
    for {
    	line, err := rd.ReadString('\n')
    	if err != nil || io.EOF == err {
            break
        }
        line = strings.Replace(line, "\n", "\r\n", -1)
        fmt.Println(line)
        fmt.Fprintln(wd, line)
    }
    wd.Flush()
    return 1
}


func cmdController() {
    format := flag.String("format", "dos2unix", "convert format")
    fileNames := flag.String("files", "", "file name to convert")
    // fmt.Println("%s", *format)
    if fileNames != nil  && format != nil {
    	strFileName := fmt.Sprintf("%s", *fileNames)
    	strFormat := fmt.Sprintf("%s", *format)
    	fmt.Println(strFileName)
    	fileName := fmt.Sprintf(strFileName, " ")
    	for _, file := range fileName {
    		fmt.Println(file)
    		if strFormat == "dos2unix" {
    			dos2unix(fmt.Sprintf("%s", file))
    		}
    	}
    }
}


func main() {
    cmdController()
}