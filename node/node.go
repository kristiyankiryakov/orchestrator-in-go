package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"orchestrator-in-go/stats"
	"orchestrator-in-go/utils"
)

type Node struct {
	Name            string
	Ip              string
	Api             string
	Memory          int64
	MemoryAllocated int64
	Disk            int64
	DiskAllocated   int64
	Stats           stats.Stats
	Role            string
	TaskCount       int
}

func NewNode(name string, api string, role string) *Node {
	return &Node{
		Name: name,
		Api:  api,
		Role: role,
	}
}

func (n *Node) GetStats() (*stats.Stats, error) {
	var resp *http.Response
	var err error

	url := fmt.Sprintf("%s/stats", n.Api)
	resp, err = utils.HTTPWithRetry(http.Get, url)
	if err != nil {
		msg := fmt.Sprintf("Unable to connect to %v. Permanent failure.\n",
			n.Api)
		fmt.Println(msg)
		return nil, errors.New(msg)
	}

	if resp.StatusCode != 200 {
		msg := fmt.Sprintf("Error retrieving stats from %v: %v\n",
			n.Api, err)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var stats stats.Stats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		msg := fmt.Sprintf("error decoding message while getting stats for node %s",
			n.Name)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	n.Memory = int64(stats.MemTotalKb())
	n.Disk = int64(stats.DiskTotal())

	n.Stats = stats

	return &stats, nil
}
