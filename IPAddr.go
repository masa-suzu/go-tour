package tour

import "fmt"

// IPAddr keeps IPV4.
type IPAddr [4]byte

// String returns ip address like "192.168.0.1".
// If address is "127.0.0.1", returns "localhost".
func (i *IPAddr) String() string {

	v := fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
	if v == "127.0.0.1" {
		return "localhost"
	}
	return v
}
