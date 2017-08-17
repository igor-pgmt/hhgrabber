package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

var vacname string                    // Vacancy name
var searchfield string                // Search by field
var items string                      // Items on page
var order_by string                   // Order by parameters
var search_period string              // Search period
var folder string = "./results"       // Folder name for results
var baselink string = "https://hh.ru" // Base website link
var delimiter string = ";"            // Delimiter for CSV saving

//var grablink string = baselink + "/search/vacname?no_magic=true&items_on_page=100&order_by=publication_time&search_period=&text=%D0%A1%D0%B8%D1%81%D1%82%D0%B5%D0%BC%D0%BD%D1%8B%D0%B9+%D0%B0%D0%B4%D0%BC%D0%B8%D0%BD%D0%B8%D1%81%D1%82%D1%80%D0%B0%D1%82%D0%BE%D1%80&salary=&experience=doesNotMatter&currency_code=RUR&search_field=name"
var searchlink string // search string for hh searching
//This file is for pagelinks
var pagelinksfile string = folder + "/" + "pagelinksfile.csv"
var resultfile string = folder + "/" + "resultfile.csv"
var pagelinks []string
var nextlinks []string
var result = make(map[int]map[string]string)
var counter int = 0 //just counter

// Function for error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage of -vacname argument is required\n\n")
		flag.PrintDefaults()
	}

	flag.StringVar(&vacname, "vacname", "", "vacancy name")
	flag.StringVar(&searchfield, "searchfield", "", "search field(name, company_name, description)")
	flag.StringVar(&items, "items", "100", "Items on page (20, 50, 100)")
	flag.StringVar(&order_by, "order_by", "publication_time", "Order by (publication_time, salary_desc, salary_asc, relevance)")
	flag.StringVar(&search_period, "search_period", "", "Search period (\"\", 7, 3, 1)")
	flag.Parse()

	if vacname == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(vacname)
	fmt.Println(searchfield)
	fmt.Println(items)
	fmt.Println(order_by)
	fmt.Println(search_period)
	fmt.Println(searchlink)

	// Creating directory for results
	err := os.Mkdir(folder, os.ModePerm)
	if err != nil {
		fmt.Println(err) // shows info if directory is exist
	}

	// Resulting files creation
	f2, err := os.Create(pagelinksfile)
	check(err)
	defer f2.Close()

	f3, err := os.Create(resultfile)
	check(err)
	defer f3.Close()
}

func main() {

	searchlink, err := url.Parse(baselink)
	check(err)
	searchlink.Path += "/search/vacancy"
	parameters := url.Values{}
	parameters.Add("text", vacname)
	parameters.Add("searchfield", searchfield)
	parameters.Add("items_on_page", items)
	parameters.Add("order_by", order_by)
	parameters.Add("search_period", search_period)
	searchlink.RawQuery = parameters.Encode()

	fmt.Println("link: " + searchlink.String())
	Scrape(searchlink.String(), &counter)
	writeresult()

}
