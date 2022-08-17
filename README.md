# geoip2-convert

This application converts Maxmind geoip2 CSV format to legacy geoip CSV format. 
This is for legacy systems where updating is a bit tricky, it's for IPV4 only. 

Input files (in same dir). 
- GeoLite2-Country-Locations-en.csv
- GeoLite2-Country-Blocks-IPv4.csv

Output :
- stdout in Legacy GeoLite format

## New GeoLite2 format (input)

|geoname_id|locale_code|continent_code|continent_name|country_iso_code|country_name|is_in_european_union|
|---|---|---|---|---|---|---|
|49518|en|AF|Africa|RW|Rwanda|0|

|network|geoname_id|registered_country_geoname_id|represented_country_geoname_id|is_anonymous_proxy|is_satellite_provider|
|---|---|---|---|---|---|
|1.0.0.0/24|2077456|2077456||0|0|

## Legacy GeoLite format (output)

|start IP|end IP|start IP int|end IP int|iso code|country name|
|---|---|---|---|---|---|
|"1.0.0.0" | "1.0.0.255" | "16777216" | "16777471" | "AU" | "Australia" |

## Prereq 
You will need go installed 

You will need to download the GeoIP2 CSV country file from Maxmind. Free, but an account is required. 
Unzip the downloaded file and place GeoLite2-Country-Locations-en.csv and 
GeoLite2-Country-Blocks-IPv4.csv in the same folder you are running the application from. 

## Install 
```
go install github.com/footfish/geoip2-convert@latest
```
## Run it
```
geoip2-convert > geoip.csv

#you can check it with a line count (should be 1 less as there is no header )
wc geoip.csv
wc GeoLite2-Country-Blocks-IPv4.csv
```
## Credits
CIDR conversion from https://gist.github.com/3c54bacef489499e2b44a075fdab6af0.git