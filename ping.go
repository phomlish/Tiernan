package main

import probing "github.com/prometheus-community/pro-bing"

func (c *commander) Ping(host string) (PingResult, error) {

	var pingResult PingResult

	pinger, err := probing.NewPinger(host)
	if err != nil {
		pingResult.Successful = false
		return pingResult, err
	}
	pinger.SetPrivileged(true)
	pinger.Count = 1
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		pingResult.Successful = false
		return pingResult, err
	}

	stats := pinger.Statistics()
	pingResult.Successful = true
	pingResult.Time = stats.AvgRtt

	return pingResult, nil
}
