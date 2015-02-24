package main

/*
CPUUtilization
FreeableMemory
WriteIOPS
ReadIOPS
ReadLatency
WriteLatency
DiskQueueDepth
*/

type ReadLatency struct {
	Timestamp   string  `json:"@timestamp"`
	ReadLatency float64 `json:"ReadLatency"`
}

type DiskQueueDepth struct {
	Timestamp      string  `json:"@timestamp"`
	DiskQueueDepth float64 `json:"DiskQueueDepth"`
}

type FreeableMemory struct {
	Timestamp      string  `json:"@timestamp"`
	FreeableMemory float64 `json:"FreeableMemory"`
}

type WriteLatency struct {
	Timestamp    string  `json:"@timestamp"`
	WriteLatency float64 `json:"WriteLatency"`
}

type CPUUtilization struct {
	Timestamp      string  `json:"@timestamp"`
	CPUUtilization float64 `json:"CPUUtilization"`
}

type ReadIOPS struct {
	Timestamp string  `json:"@timestamp"`
	ReadIOPS  float64 `json:"ReadIOPS"`
}

type WriteIOPS struct {
	Timestamp string  `json:"@timestamp"`
	WriteIOPS float64 `json:"WriteIOPS"`
}
