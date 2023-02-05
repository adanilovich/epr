package main

import (
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/adanilovich/epr/pkg/browser"
)

const defaultLogPath = "stdout"
const outputDir = "./reports"

var usageMessage = `
[ Examples ]
pe https://domain.com/path 	Extract only virtual DOM and direct the result to stdout
pe -ao myreports https://domain.com/path 	 Write all resources to myreports folder 
pe -a https://domain.com/path 	 		   Write all resources to default folder 
`

// pe - page extractor
func main() {
	var targetURL, resFolder string
	var ok, flagAllResources, flagDisableHeadless bool
	pflag.ErrHelp = errors.New(usageMessage)
	pflag.Usage = usage

	pflag.BoolVarP(&flagAllResources, "all-resources", "a", false, "Extract all resources from target page")
	pflag.BoolVarP(&flagDisableHeadless, "disable-headless", "d", false, "Show browser")
	pflag.StringVarP(&resFolder, "output-directory", "o", outputDir, "Resources directory")
	pflag.Parse()

	// 	std := bufio.NewScanner(os.Stdin)
	b, err := browser.NewBrowser(flagDisableHeadless)
	if err != nil {
		log.Println(err)
	}

	if targetURL, ok = parseURL(); !ok {
		fmt.Fprintf(os.Stderr, "[!] Not found target url \n\n")
		pflag.Usage()
		fmt.Fprintf(os.Stderr, "%s\n", usageMessage)
		os.Exit(1)
	}

	urlText := strings.ToLower(targetURL)
	urlInstance, err := url.Parse(urlText)
	if err != nil {
		fmt.Printf("Bad URL: %v\n", err)
		os.Exit(1)
	}

	// make the output directory
	dirPath := resFolder + "/" + urlInstance.Hostname()
	err = os.MkdirAll(dirPath+"/resources", 0750)
	if err != nil {
		fmt.Printf("mkdirall: %v\n", err)
		os.Exit(1)
	}

	if urlText[len(urlText)-1] == '/' {
		urlText = urlText[:len(urlText)-1]
	}

	if flagAllResources {
		b.ExtractAllResources(dirPath, urlText)
		return
	}
	b.ExtractVirtualDOM(urlText)
}

func usage() {
	fmt.Fprintf(os.Stderr, "-----\n")
	fmt.Fprintf(os.Stderr, "[EPR] Extractor of page resources\n")
	fmt.Fprintf(os.Stderr, "-----\n")
	fmt.Fprintf(os.Stderr, "EPR allows to retrieve and save page resources to folder in grepable form for following processing ones through appropriate parsers.\n")
	fmt.Fprintf(os.Stderr, "Types of resources:\n * Virtual DOM\n * Local/Session storage\n * js/css/svg files\n * intercepted data of requests/responses performed by page scripts. It works like devtools > network.\n\nIt means that the data are kept in one place. It's useful for parsing e.g. to find all api-keys, jwt tokens, potential vulnerable input data like http parameters, cookies, headers etc.\n")
	fmt.Fprintf(os.Stderr, "EPR works using chromium engine.\n\n")
	fmt.Fprintf(os.Stderr, "-----\n")
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	pflag.PrintDefaults()
}

func parseURL() (url string, ok bool) {
	for _, el := range os.Args {
		if strings.Contains(el, "http://") || strings.Contains(el, "https://") {
			return el, true
		}
	}
	return "", false
}


