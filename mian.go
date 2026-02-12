package main

import (
	"fmt"
	"github.com/zan8in/masscan"
)

func main() {

	fmt.Printf("Welcome to M.E.F.!\nMASS EXPLOITATION FRAMEWORK\n\nThis framework automates the process of vulnerability scans and exploit automation on the internet.\n   1 - Add Module\n   2 - Balls\n   3 - Exit\n")
	var input string
	fmt.Scanln(&input)

	if input == "1" {

		fmt.Printf("Insert ports to search for, seperated by space and comma. Ex \"22, 80\"\n")

		var ports string
		fmt.Scanln(&ports)

		fmt.Printf("%s", ports)
	}
}

func ScanPort(selfIp string) {

	masscan.SetParamTargets("0.0.0.0/24")
	masscan.SetParamPorts("80")
	masscan.SetParamRate(1000000)
	masscan.SetParamExclude("255.255.255.255")

}
