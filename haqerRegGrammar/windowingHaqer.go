package main

import "fmt"

// What if we windowed out the 545bp of DNA into 200bp windows? What about 150bp?
// If we walk the windows by 1bp at a time.
// These windows of the total 545bp would be tested in an assay (STARR-seq).
// This command determines the total possible combinations to be tested with the windows of 150 to 200bp.

func main() {
low, high := 150, 200
s := make([]int, high-low+1)
for i := range s {
	s[i] = i + low
}

sum := 0
bp := 545
for i := range s {
	sum += bp-i
}


fmt.Println(s)
fmt.Println(sum)


}
