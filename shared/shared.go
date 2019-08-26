package shared

import (
	"fmt"
	"os"
)

const UDP_PORT = 1200
const TCP_PORT = 7171

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}