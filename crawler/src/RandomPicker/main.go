package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var codeforceBase string = "http://codeforces.com/"
var bojBase string = "https://www.acmicpc.net/"

func main() {
	channel := make(chan string)

	codeforceTarget := getRandomPage(getPageNum(codeforceBase+"problemset"), codeforceBase+"problemset/page")
	go getPage(codeforceTarget, "td", channel, 0)

	bojTarget := getRandomPage(getPageNum(bojBase+"problemset"), bojBase+"problemset")
	go getPage(bojTarget, ".list_problem_id", channel, 1)

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getRandomPage(endPage int, baseURL string) string {
	return baseURL + "/" + strconv.Itoa(rand.Intn(endPage)+1)
}

func getPageNum(baseURL string) int {
	res, err := http.Get(baseURL)
	checkErr(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	ret := 0

	doc.Find("[class*=\"pagination\"]").Each(func(i int, pages *goquery.Selection) {
		pages.Find("a").Each(func(j int, page *goquery.Selection) {
			num, err := strconv.Atoi(page.Text())
			if err == nil {
				if ret < num {
					ret = num
				}
			}
		})
	})

	return ret
}

func getPage(targetURL string, selector string, channel chan<- string, stype int) {
	res, err := http.Get(targetURL)
	checkErr(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	ret := []string{}

	doc.Find(selector).Each(func(i int, problems *goquery.Selection) {
		if stype == 1 {
			ret = append(ret, problems.Text())
		} else if stype == 0 {
			problems.Find("a").Each(func(j int, problem *goquery.Selection) {
				name, isExist := problem.Attr("href")
				if isExist {
					ret = append(ret, name)
				}
			})
		}
	})
	msg := ret[rand.Intn(len(ret))]
	if stype == 0 {
		tmp := strings.Split(msg, "/")
		msg = strings.Join(tmp[len(tmp)-2:len(tmp)], "/")
		channel <- codeforceBase + "problemset/problem/" + msg
	} else if stype == 1 {
		channel <- bojBase + "problem/" + msg
	}
}
