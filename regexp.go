package main

import (
	"fmt"
	"regexp"
)

func main() {
	// https://play.golang.org/p/_8HSgiNNmS

	re := regexp.MustCompile("[_]")
	// allow valid subdomain
	if CheckSubdomain := re.MatchString("manigandan"); CheckSubdomain != false {
		fmt.Println(CheckSubdomain)
		fmt.Println("Invalid subdomain. Don't use underscore on your subdomain.")
	} else {
		fmt.Println("Valid subdomain. Proceed!")
	}
	// don't allow invalid subdomain
	if CheckSubdomain := re.MatchString("manigandan_jeff"); CheckSubdomain != false {
		fmt.Println(CheckSubdomain)
		fmt.Println("Invalid subdomain. Don't use underscore on your subdomain.")
	} else {
		fmt.Println("Valid subdomain. Proceed!")
	}
	//other exapmles
	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)

	matched, err = regexp.MatchString("bar.*", "seafood")
	fmt.Println(matched, err)

	matched, err = regexp.MatchString("a(b", "seafood")
	fmt.Println(matched, err)
}
