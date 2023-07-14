// Echo2 prints its commandline
// arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Println(i, arg)
	}
}

// TODO: Exercise 1.3: Experiment to measure the dif ference in running time bet ween our potential ly
// inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates par t of the
// time package, and Sec tion 11.4 shows how to write benchmark tests for systematic performance
// evaluation.)
