package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Trace the IP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
			}
		} else {
			fmt.Println("Please provide IP to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

}

type Ip struct {
	IP       string `json::"ip"`
	City     string `json::"city"`
	Region   string `json::"region"`
	Country  string `json::"country"`
	Loc      string `json::"loc"`
	TimeZone string `json::"timezone"`
	Postal   string `json::"postal"`
}

func showData() {
	url := "http://ipinfo.io/1.1.1.1/geo"
	responseByte := getData(url)

}

func getData(url string) []byte {

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Unable to read the response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Unable to read the response")
	}

	return responseByte

}
