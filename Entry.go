package main

import (
	"fmt"
	"github.com/polyglotDataNerd/poly-covid19/sources"
	"sync"
	"time"
)

func main() {
	chLine := make(chan string)
	chOut := make(chan string)
	var wg sync.WaitGroup

	jhu := sources.JHU{
		ChannelLine: chLine,
		ChannelOut:  chOut,
		S3Bucket:    "poly-testing",
		S3key:       "covid/jhu/raw",
		Wg:          wg,
	}
	jhu.Munge("poly-testing", fmt.Sprintf("%s%s%s%s%s", "covid/jhu/transformed/", time.Now().Format("2006-01-02"), "/jhu_", time.Now().Format("2006-01-02"), ".gz"))
	sources.GetCSV("poly-testing", fmt.Sprintf("%s%s%s%s%s", "covid/cds/", time.Now().Format("2006-01-02"), "/cds_", time.Now().Format("2006-01-02"), ".gz"))
}
