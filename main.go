package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	cidrtoip "github.com/footfish/geoip2-convert/internal/cidr-to-ip"
)

const geoLite2Locations = "GeoLite2-Country-Locations-en.csv"
const geoLite2Blocks = "GeoLite2-Country-Blocks-IPv4.csv"

type Location struct {
	//	locale_code          string
	//	continent_code       string
	//	continent_name       string
	country_iso_code string
	country_name     string
	//	is_in_european_union string
}

func main() {

	//Load complete country location CSV file to map
	locations, err := readLocations(geoLite2Locations)
	if err != nil {
		log.Fatal(err)
	}

	//Load blocks CSV file
	f, err := os.Open(geoLite2Blocks)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	// now process CSV records per line
	var n int
	for {
		record, err := r.Read()
		if err == io.EOF { //finished
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if n > 0 { //skip headers
			// convert cidr to ip range
			startIP, endIP, err := cidrtoip.CIDRRangeToIPv4Range([]string{record[0]})
			if err != nil {
				log.Fatal(err)
			}
			// convert ip range to int
			startIPint := cidrtoip.IPv4ToUint32(startIP)
			endIPint := cidrtoip.IPv4ToUint32(endIP)
			// lookup country and iso
			geonameId := record[1]
			if geonameId == "" {
				geonameId = record[2] //fallback to registered_country_geoname_id
			}
			country, ok := locations[geonameId]
			if !ok {
				log.Fatal(fmt.Errorf("error: country index %s does not exist line %d", geonameId, n+1))
			}
			// csv
			fmt.Printf("\"%s\",\"%s\",\"%d\",\"%d\",\"%s\",\"%s\"\n", startIP, endIP, startIPint, endIPint, country.country_iso_code, country.country_name)

		}
		n++
	}
}

//readLocations() reads country CSV file into map
func readLocations(fileName string) (map[string]Location, error) {
	locations := map[string]Location{}

	f, err := os.Open(fileName)
	if err != nil {
		return locations, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	// skip first line
	if _, err := r.Read(); err != nil {
		return locations, err
	}

	records, err := r.ReadAll()
	if err != nil {
		return locations, err
	}

	for _, record := range records {
		location := Location{
			//			locale_code:          record[1],
			//			continent_code:       record[2],
			//			continent_name:       record[3],
			country_iso_code: record[4],
			country_name:     record[5],
			//			is_in_european_union: record[6],
		}
		locations[record[0]] = location //geoname_id index
	}
	return locations, nil
}
