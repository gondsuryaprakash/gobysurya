package http

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func GetIp() {
	fmt.Println("GetIp")
	handlerFunc := http.HandlerFunc(handleRequest)
	http.Handle("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ip, err := getIp(r)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("No Valid IP"))
	}

	w.WriteHeader(200)
	w.Write([]byte(ip))

}

func getIp(r *http.Request) (ip string, err error) {

	// Get IP from X-REAL-IP header
	ip = r.Header.Get("X-REAL-IP")
	fmt.Println("Ip from X-REAL-IP", ip)
	netIp := net.ParseIP(ip)
	if netIp != nil {
		return ip, nil
	}

	// Get IP from X-FORWARDED-FOR header

	ips := r.Header.Get("X-FORWARDED-FOR")
	fmt.Println("Ip from X-FORWARDED-FOR", ip)
	splitIps := strings.Split(ips, ",")

	for _, val := range splitIps {

		netIp := net.ParseIP(val)
		if netIp != nil {
			return val, nil
		}
	}

	// Get IP from RealAddr

	ip, _, err = net.SplitHostPort(r.RemoteAddr)
	fmt.Println("Ip from RemoteAddr", ip)
	if err != nil {
		return "", err
	}
	netIp = net.ParseIP(ip)
	if netIp != nil {
		return ip, nil
	}

	return "", errors.New("Not valid IP")

}
