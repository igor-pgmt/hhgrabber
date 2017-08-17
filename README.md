# hhgrabber
This script grabs search results from hh.ru and saves them to the CSV.
The result contains columns: vacancy name, link to the vacancy, salary, employer's name and link to the emloyer.

## Installation
Download the script from Github or clone it:
```bash
git clone https://github.com/igor-pgmt/hhgrabber.git
```
Build it:
```bash
go build -v hhgrabber.go grabber.go
```

## Usage
Run the script from the terminal or cmd:
```bash
./hhgrabber -vacname "System administrator"
```
The "vacname" argument is required.<br />
You can specify additional arguments:
```bash
./hhgrabber -vacname "System administrator" -order_by salary_desc -search_period 3 -searchfield name
```
This means you'll grab vacansies with the name "System administrator" in the past three days, ordered by salary.

 Or, execute the script without arguments to read the help:
```bash
 ./hhgrabber 
Usage of -vacname argument is required

  -items string
    	Items on page (20, 50, 100) (default "100")
  -order_by string
    	Order by (publication_time, salary_desc, salary_asc, relevance) (default "publication_time")
  -search_period string
    	Search period ("", 7, 3, 1)
  -searchfield string
    	search field(name, company_name, description)
  -vacname string
    	vacancy name
```
## Result
You can see a result at the folder named "result" at the script's directory.

## Dependencies
The script uses goquery (https://github.com/PuerkitoBio/goquery) package for websites scraping.