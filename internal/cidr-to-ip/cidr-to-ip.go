//from https://gist.github.com/3c54bacef489499e2b44a075fdab6af0.git
package cidrtoip

import (
	"fmt"
	"strconv"
	"strings"
)

// Convert CIDR to IPv4 range
func CIDRRangeToIPv4Range(cidrs []string) (ipStart string, ipEnd string, err error) {

	var ip uint32 // ip address

	var ipS uint32 // Start IP address range
	var ipE uint32 // End IP address range

	for _, CIDR := range cidrs {

		cidrParts := strings.Split(CIDR, "/")

		ip = IPv4ToUint32(cidrParts[0])
		bits, _ := strconv.ParseUint(cidrParts[1], 10, 32)

		if ipS == 0 || ipS > ip {
			ipS = ip
		}

		ip = ip | (0xFFFFFFFF >> bits)

		if ipE < ip {
			ipE = ip
		}

	}

	ipStart = uInt32ToIPv4(ipS)
	ipEnd = uInt32ToIPv4(ipE)

	return ipStart, ipEnd, err
}

// Convert IPv4 to uint32
func IPv4ToUint32(iPv4 string) uint32 {

	ipOctets := [4]uint64{}

	for i, v := range strings.SplitN(iPv4, ".", 4) {
		ipOctets[i], _ = strconv.ParseUint(v, 10, 32)
	}

	result := (ipOctets[0] << 24) | (ipOctets[1] << 16) | (ipOctets[2] << 8) | ipOctets[3]

	return uint32(result)
}

// Convert uint32 to IP
func uInt32ToIPv4(iPuInt32 uint32) (iP string) {
	iP = fmt.Sprintf("%d.%d.%d.%d",
		iPuInt32>>24,
		(iPuInt32&0x00FFFFFF)>>16,
		(iPuInt32&0x0000FFFF)>>8,
		iPuInt32&0x000000FF)
	return iP
}
