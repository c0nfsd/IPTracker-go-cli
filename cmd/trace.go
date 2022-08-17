package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/color"
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
				showData(ip)
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

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := Ip{}
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		fmt.Println("Unable to unmarshal the response")
	}
	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
	c.Println("DATA FOUND:\n")
	fmt.Printf("IP: %s\nCITY: %s\nREGION: %s\nCOUNTRY: %s\nLOCATION: %s\nTIMEZONE: %s\nPOSTAL: %s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.TimeZone, data.Postal)
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
