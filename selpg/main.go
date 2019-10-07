package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"os"
	"os/exec"
)

var (
	startPage int
	endPage int
	inFileName string
	pageLen int /* default value, can be overriden by "-l number" on command line */
	pageType string /* 'l' for lines-delimited, 'f' for form-feed-delimited, default is 'l' */
	printDest string
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: selpg [OPTION]... [FILE]")
	fmt.Fprintln(os.Stderr, "Select pages from input [FILE] or stdin if [FILE] is not specified")
	fmt.Fprintln(os.Stderr, "")
	pflag.PrintDefaults()

	os.Exit(2)
}

func perror(err error) {
	if err != nil {
		if _, err2 := fmt.Fprintf(os.Stderr, "%s\n", err.Error()); err2 != nil {
			panic(err2)
		}
		os.Exit(1)
	}
}

func processArguments() error {
	pflag.IntVarP(&startPage, "start-page", "s", -1, "start page")
	pflag.IntVarP(&endPage, "end-page", "e", -1, "end page")
	pflag.IntVarP(&pageLen, "page-len", "l", 72, "page length")
	pageTypeFlag := pflag.BoolP("form-feed-delimited", "f", false, "form-feed-delimited")
	pflag.StringVarP(&printDest, "print-dest", "d", "", "printer destination")
	pflag.Parse()

	if *pageTypeFlag {
		pageType = "f"
	} else {
		pageType = "l"
	}

	if startPage == -1 || endPage == -1 {
		usage()
	}

	if startPage <= 0 || endPage <= 0 {
		return fmt.Errorf("%s", "start page and end page should be positive")
	}

	if startPage > endPage {
		return fmt.Errorf("%s", "End page should be greater then start page")
	}

	if pflag.NArg() == 1 {
		inFileName = pflag.Arg(0)
	} else if pflag.NArg() > 1 {
		return fmt.Errorf("%s", "Too many arguments")
	} else {
		inFileName = ""
	}

	return nil
}

func processInput() error {
	var reader *bufio.Reader
	var writer *bufio.Writer

	// set the input source
	if len(inFileName) > 0 {
		fin, err := os.Open(inFileName)
		if err != nil {
			return fmt.Errorf("selpg: could not open input file \"%s\". Reason:\n%s\n", inFileName, err.Error())
		}
		defer fin.Close()
		reader = bufio.NewReader(fin)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	// set the output destination
	if len(printDest) > 0 {
		cmd := exec.Command("lp", fmt.Sprintf("-d%s", printDest))
		stdinPipe, err := cmd.StdinPipe()
		if err != nil {
			return fmt.Errorf("selpg: could not open pipe to \"%s\". Reason:\n%s\n", printDest, err.Error())
		}
		defer stdinPipe.Close()
		cmd.Stdout = os.Stdout
		if err := cmd.Start(); err != nil {
			return err
		}
		writer = bufio.NewWriter(stdinPipe)
	} else {
		writer = bufio.NewWriter(os.Stdout)
	}
	defer writer.Flush()

	lineCtr := 0
	pageCtr := 1

	if pageType == "l" {
		line := bufio.NewScanner(reader)
		for line.Scan() {
			lineCtr++
			if lineCtr > pageLen {
				pageCtr++
				lineCtr = 1
			}
			if pageCtr >= startPage && pageCtr <= endPage {
				if _, err := writer.Write(line.Bytes()); err != nil {
					return err
				}
				if err := writer.WriteByte('\n'); err != nil {
					return err
				}
			}
		}
	} else {
		for {
			c, err := reader.ReadByte()
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}
			if c == '\f' {
				pageCtr++
			}

			if pageCtr >= startPage && pageCtr <= endPage {
				if err := writer.WriteByte(c); err != nil {
					return err
				}
			}
		}
	}

	if pageCtr < startPage {
		return fmt.Errorf("start_page (%d) greater than total pages (%d), no output written\n", startPage, pageCtr)
	} else if pageCtr < endPage {
		return fmt.Errorf("end_page (%d) greater than total pages (%d), less output than expected\n", endPage, pageCtr)
	}

	return nil
}

func main() {
	perror(processArguments())
	perror(processInput())
}