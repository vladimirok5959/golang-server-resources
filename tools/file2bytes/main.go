package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var ParamSrc string
var ParamPak string
var ParamVar string
var ParamDst string

func init() {
	flag.StringVar(&ParamSrc, "src", "", "file which will be converted to bytes")
	flag.StringVar(&ParamPak, "pak", "", "golan package name")
	flag.StringVar(&ParamVar, "var", "", "golan variable name")
	flag.StringVar(&ParamDst, "dst", "", "result file (if not set, stdout will be used)")
	flag.Parse()
}

func help() {
	fmt.Fprintf(os.Stdin, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stdin, "  -src string\n")
	fmt.Fprintf(os.Stdin, "    	file which will be converted to bytes\n")
	fmt.Fprintf(os.Stdin, "  -pak string\n")
	fmt.Fprintf(os.Stdin, "    	golan package name\n")
	fmt.Fprintf(os.Stdin, "  -var string\n")
	fmt.Fprintf(os.Stdin, "    	golan variable name\n")
	fmt.Fprintf(os.Stdin, "  -dst string\n")
	fmt.Fprintf(os.Stdin, "    	result file (if not set, stdout will be used)\n")
}

func main() {
	if len(os.Args) < 4 {
		help()
		return
	}

	ParamSrc = os.Args[1]
	ParamPak = os.Args[2]
	ParamVar = os.Args[3]
	if len(os.Args) >= 5 {
		ParamDst = os.Args[4]
	}

	if len(ParamSrc) <= 0 || len(ParamPak) <= 0 || len(ParamVar) <= 0 {
		help()
		return
	}

	st, err := os.Stat(ParamSrc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	if st.Mode().IsDir() {
		fmt.Fprintf(os.Stderr, "source file (%s) is a directory\n", ParamSrc)
		return
	}

	bytes, err := ioutil.ReadFile(ParamSrc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	str := fmt.Sprintf("%v", bytes)
	str = str[1 : len(str)-1]
	str = strings.Replace(str, " ", ", ", -1)
	str = "package " + ParamPak + "\n\n" + "var " + ParamVar + " = []byte{" + str + "}\n"

	if len(ParamDst) > 0 {
		f, err := os.Create(ParamDst)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			return
		}
		if _, err := f.Write([]byte(str)); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			return
		}
		if err := f.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			return
		}
	} else {
		fmt.Fprintf(os.Stdin, "%s", str)
	}
}
