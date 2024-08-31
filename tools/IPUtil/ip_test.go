package IPUtil

import "testing"

func TestIpv4ToLong(t *testing.T) {
	ip := "153.3.238.110"
	l := Ipv4ToLong(ip)

	if LongToIpv4(l) != ip {
		t.Errorf("address %s long %d", ip, l)
	}
	t.Logf("address %s long %d", ip, l)
}
