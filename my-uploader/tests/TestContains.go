package main

import "strings"

func main()  {
	content := "data:image/png;base64,iVBORw0KGgoA"
	println(strings.Contains(content, "image/png"))
}
