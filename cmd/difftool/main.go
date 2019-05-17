package main

import (
	"fmt"

	"github.com/joincivil/recruitment-backend-go/pkg/diff"
)

func differImplementation() diff.Differ {
	// TODO: Init and return differ implementation here.
	return nil
}

func main() {
	d := differImplementation()
	if d == nil {
		fmt.Println("needs implementation")
		return
	}

	err := d.Run("sludge")
	if err != nil {
		fmt.Printf("error = %v\n", err)
		return
	}

	err = d.Run("documented")
	if err != nil {
		fmt.Printf("error = %v\n", err)
		return
	}

	err = d.Run("ecowurd")
	if err != nil {
		fmt.Printf("error = %v\n", err)
		return
	}

	fmt.Println("Done.")
}
