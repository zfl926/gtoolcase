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
	Format               int                     `format 1 dos2unix 2 unix2dos`
}

type CmdExcutor interface {
    Execute()
}


func (this *Cmd) Execute(){
    for _, fileName := range this.FileNames {
        if this.Format == 1 {
            dos2unix(fileName)
        } else {
            dos2unix(fileName)
        }
    }
}

func dos2unix(fileName string) int {
	rf, err := os.Open(fileName)
    if err != nil {
        return -1
    }
    wf, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
    if err != nil {
    	return -1
    }
    defer rf.Close()
    rd := bufio.NewReader(rf)
    wd := bufio.NewWriter(wf)
    var strlines string
    for {
    	line, err := rd.ReadString('\n')
    	if err != nil || io.EOF == err {
            break
        }
        line = strings.Replace(line, "\r\n", "\n", -1)
  		strlines = strlines + line
    }
    wd.WriteString(strlines)
    wd.Flush()
    return 1
}

func unix2dos(fileName string) int {
    rf, err := os.Open(fileName)
    if err != nil {
        return -1
    }
    wf, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
    if err != nil {
        return -1
    }
    defer rf.Close()
    rd := bufio.NewReader(rf)
    wd := bufio.NewWriter(wf)
    var strlines string
    for {
        line, err := rd.ReadString('\n')
        if err != nil || io.EOF == err {
            break
        }
        line = strings.Replace(line, "\n", "\r\n", -1)
        strlines = strlines + line
    }
    wd.WriteString(strlines)
    wd.Flush()
    return 1
}


func cmdController() {
    format := flag.String("format", "dos2unix", "convert format")
    fileNames := flag.String("files", "", "file name to convert")
    flag.Parse()
    if fileNames != nil  && format != nil {
    	strFileName := fmt.Sprintf("%s", *fileNames)
    	strFormat := fmt.Sprintf("%s", *format)
    	fileName := strings.Split(strFileName, " ")
        var intFormat int = 1
        if strFormat == "dos2unix" {
            intFormat = 1
        } else {
            intFormat = 2
        }
        var cmd Cmd = Cmd{
            FileNames : fileName,
            Format    : intFormat,
        }
        cmd.Execute()
    }
}


func main() {
    cmdController()
}