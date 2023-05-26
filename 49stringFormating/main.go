package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	// Value of instance ~ %v
	fmt.Printf("struct 1: %v\n", p)

	// Value of struct with field names ~ %+v
	fmt.Printf("struct 2: %+v\n", p)

	// Instance with the source code snippet that produced the value ~ %#v
	fmt.Printf("struct 3: %#v\n", p)

	// For type ~ %t
	fmt.Printf("type: %T\n", p)

	// For Bool ~ %t
	fmt.Printf("bool: %t\n", true)

	// For base 10 integer representation ~ %d
	fmt.Printf("int: %d\n", 123)

	// For binary representation ~ %b
	fmt.Printf("bin: %b\n", 14)

	// For character corresponding to interger ~ %c
	fmt.Printf("char: %c\n", 65)

	// For hex encoding ~ %x
	fmt.Printf("hex: %x\n", 456)

	// For basic decimal formating ~ %f
	fmt.Printf("float1: %f\n", 78.9)

	// For float in scientific notation ~ %e | %E
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// For basic string ~ %s
	fmt.Printf("str1: %s\n", "\"string\"")

	// For double quote string ~ %q
	fmt.Printf("str2: %q\n", "\"string\"")

	// For str hex ~ %x
	fmt.Printf("str3: %x\n", "hex this")

	// For pointer ~ %p
	fmt.Printf("pointer: %p\n", &p)

	// For giving width integer ~%[width]d
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// For width and precision to float ~ %[width].[precision]f
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.457)

	// To justify left add "-"
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// For giving width to string ~ %[width]s
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// To justify left add "-"
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// To return a string with out printing
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")

}
