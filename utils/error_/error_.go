package error_

import (
	"fmt"
	"os"
)

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
