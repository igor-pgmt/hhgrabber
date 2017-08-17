package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
)

func Scrape(searchlink string, counter *int) {
	doc, err := goquery.NewDocument(searchlink)
	check(err)

	/*	f1, err := os.OpenFile(filename1, os.O_APPEND|os.O_WRONLY, 0600)
		check(err)
		defer f1.Close()
	*/
	// Find the review items
	// doc.Find(".search-result-item__head").Each(func(i int, s *goquery.Selection) {
	doc.Find(".search-result-description__item_primary").Each(func(i int, s *goquery.Selection) {

		// For each item found, get the band and title
		salary := s.ChildrenFiltered(".b-vacancy-list-salary").Text()
		companyresult := s.ChildrenFiltered(".search-result-item__company")
		companame := strings.Replace(companyresult.Text(), delimiter, "", -1)
		//cli := companyresult.ChildrenFiltered("a")
		companyLink, _ := companyresult.ChildrenFiltered(".bloko-link.bloko-link_secondary").Attr("href")

		//findresult := s.Find(".search-result-item__head")
		findresult := s.ChildrenFiltered(".search-result-item__head")

		name := findresult.Text()
		href, _ := findresult.ChildrenFiltered("a").Attr("href")
		href = strings.Split(href, "?")[0]

		href += ";" + salary + ";" + companame + ";"

		if companyLink != "" {
			href += baselink + companyLink
		}

		pagelinks = append(pagelinks, href)
		result[*counter] = map[string]string{name: href}
		*counter++
		/*		writeme := href + string("\n")
				_, err := f1.WriteString(writeme)
				check(err)*/
	})

	f2, err := os.OpenFile(pagelinksfile, os.O_APPEND|os.O_WRONLY, 0600)
	check(err)
	defer f2.Close()

	doc.Find(".b-pager__next").Each(func(i int, s *goquery.Selection) {
		ahref, _ := s.Find("a").Attr("href")

		newlink := baselink + ahref
		if newlink != baselink {
			nextlinks = append(nextlinks, ahref)
			writeme := newlink + string("\n")
			_, err := f2.WriteString(writeme)
			check(err)
			Scrape(newlink, *&counter)
		}

	})

}

func writeresult() {
	f3, err := os.OpenFile(resultfile, os.O_APPEND|os.O_WRONLY, 0600)
	check(err)
	defer f3.Close()

	fmt.Println(len(result))

	for i := 0; i < len(result); i++ {

		for key, value := range result[i] {
			fmt.Println("Key:", key, "Value:", value)
			writeme := key + delimiter + value + "\n"
			_, err := f3.WriteString(writeme)
			check(err)
		}
	}
}
