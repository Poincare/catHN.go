package main

import (
	"fmt"
	"http"
	"html"
	"io"
	"flag"
)

func printLenLine(str string) {
	l_s := len(str)

	for i := 0; i < l_s+2; i++ {
		fmt.Print("_")
	}

}

func printBox(str string) {
	printLenLine(str)

	fmt.Println("")
	fmt.Println("")

	fmt.Print("| ")
	fmt.Print(str)
	fmt.Print(" |")

	fmt.Println("")

	printLenLine(str)

	fmt.Println("")
}

func parseHTML(r io.Reader, np int) {

	doc, err := html.Parse(r)

	if err != nil {
		fmt.Println("Error occurred in parsing HTML!")
		return
	}

	var f func(*html.Node)

	i := 0
	max_stories := np 

	f = func(n *html.Node) {
		if i == max_stories {
			return
		}

		if n.Type == html.ElementNode && n.Data == "td" {
			if len(n.Attr) != 0 {
				if n.Attr[0].Key == "class" && n.Attr[0].Val == "title" {
					printBox(n.Child[0].Data)
					i++
				}
			}
		}

		for _, c := range n.Child {
			f(c)
		}
	}

	f(doc)
}

func catHN(URL string, np int) {
	res, url, error := http.Get(URL)

	if error == nil {
		parseHTML(res.Body, np)

	} else {
		fmt.Printf("%s\n", error)

	}
}

func parseArgs() (string, int) {
	URL := "http://news.ycombinator.com/"

	newposts := flag.Bool("new", false, "Show new posts")
	numposts := flag.Int("posts", 5, "How many posts to show")

  flag.Parse()

	if *newposts {
		fmt.Println("Newest Posts\n")
		URL = "http://news.ycombinator.com/newest"

	} else {
		fmt.Println("Top Posts\n")
	}

	return URL, *numposts
}

func main() {
	URL, numposts := parseArgs() 

	catHN(URL, numposts)
}
