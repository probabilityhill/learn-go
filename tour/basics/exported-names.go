/*
+ exported names: begin w/ a capital letter
+ can refer only to its exported names
+ unexported names: not accessible from outside the package
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}

/*
3.141592653589793
*/
