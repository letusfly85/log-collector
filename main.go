package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

var cfg Config = Config{}
var logDir string

func main() {
	var ymd string
	var hostName string
	flag.StringVar(&ymd, "ymd", "20150101", "yyyymmdd")
	flag.StringVar(&hostName, "hostname", "your host", "hostname")
	flag.Parse()

	cfg, _ := GetConfig("conf.gcfg")
	logDir = cfg.LogInfo.LogDir

	indexName := hostName + "_" + ymd
	documentNames := []string{"CPUUtilization", "FreeableMemory", "WriteIOPS", "ReadIOPS", "ReadLatency", "WriteLatency", "DiskQueueDepth"}

	for _, documentName := range documentNames {
		files := matchedPathList(hostName, documentName, ymd)
		params := make(map[int][]byte, 0)
		index := 0
		for _, file := range files {
			index, params = generateParams(file, documentName, index, params)
		}
		println(documentName)
		put2ElasticSearch(indexName, documentName, params)
	}
}

func put2ElasticSearch(indexName string, documentName string, params map[int][]byte) {
	uri := "http://localhost/es" + "/" + indexName + "/" + documentName

	for index, value := range params {
		_uri := uri + "/" + strconv.Itoa(index)
		binary := bytes.NewReader(value)
		req, _ := http.NewRequest("PUT", _uri, binary)

		client := &http.Client{}
		_, err := client.Do(req)
		if err != nil {
			println(err.Error())
		}
	}
}

func matchedPathList(hostName string, documentName string, ymd string) []string {
	pattern := "*" + hostName + "*" + documentName + "*.dat"
	path := filepath.Join(logDir, "log", "target_log", ymd, pattern)
	files, _ := filepath.Glob(path)

	return files
}

func generateParams(file string, documentName string, ind int, wa map[int][]byte) (index int, result map[int][]byte) {
	_contents, _ := ioutil.ReadFile(file)
	contents := string(_contents)
	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		ary := strings.Split(line, "\t")

		if len(ary) > 1 {
			value, _ := strconv.ParseFloat(ary[2], 32)
			_time := ary[0][0:4] + "-" + ary[0][4:6] + "-" + ary[0][6:8] + " " +
				ary[1][0:2] + ":" + ary[1][2:4] + ":" + ary[1][4:6]

			param := make([]byte, 0)
			switch documentName {
			case "WriteIOPS":
				target := WriteIOPS{_time, value}
				param, _ = json.Marshal(target)

			case "ReadIOPS":
				target := ReadIOPS{_time, value}
				param, _ = json.Marshal(target)

			case "WriteLatency":
				target := WriteLatency{_time, value}
				param, _ = json.Marshal(target)

			case "ReadLatency":
				target := ReadLatency{_time, value}
				param, _ = json.Marshal(target)

			case "DiskQueueDepth":
				target := DiskQueueDepth{_time, value}
				param, _ = json.Marshal(target)

			case "CPUUtilization":
				target := CPUUtilization{_time, value}
				param, _ = json.Marshal(target)

			case "FreeableMemory":
				target := FreeableMemory{_time, value}
				param, _ = json.Marshal(target)

			}

			ind = ind + 1
			wa[ind] = param
		}
	}
	result = wa
	index = ind

	return index, result
}
