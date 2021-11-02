package testcli

import (
	"errors"
	"net"
)

func HostIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	bestScore := -1
	var bestIP net.IP
	// Select the highest scoring IP as the best IP.
	for _, eth := range interfaces {
		addrs, err := eth.Addrs()
		if err != nil {
			// Skip this interface if there is an error.
			continue
		}
		for _, addr := range addrs {
			score, ip := score(eth, addr)
			if score > bestScore {
				bestScore = score
				bestIP = ip
			}
		}
	}
	if bestScore == -1 {
		return "", errors.New("no addresses to listen on")
	}
	return bestIP.String(), nil
}
func score(eth net.Interface, addr net.Addr) (int, net.IP) {
	var ip net.IP
	if netAddr, ok := addr.(*net.IPNet); ok {
		ip = netAddr.IP
	} else if netIP, ok := addr.(*net.IPAddr); ok {
		ip = netIP.IP
	} else {
		return -1, nil
	}
	var score int
	if ip.To4() != nil {
		score += 300
	}
	if eth.Flags&net.FlagLoopback == 0 && !ip.IsLoopback() {
		score += 100
		if eth.Flags&net.FlagUp != 0 {
			score += 100
		}
	}
	return score, ip
}
