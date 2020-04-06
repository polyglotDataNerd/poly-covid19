package sources

import (
	"fmt"
	"github.com/polyglotDataNerd/zib-Go-utils/aws"
	"github.com/polyglotDataNerd/zib-Go-utils/reader"
	"github.com/polyglotDataNerd/zib-Go-utils/scanner"
	"github.com/polyglotDataNerd/zib-Go-utils/utils"
	"strings"
	"sync"
	"time"
)

type DataModel struct {
	FIPS          string
	Admin         string
	ProvinceState string
	CountryRegion string
	LastUpdate    string
	Latitude      string
	Longitude     string
	Confirmed     string
	Deaths        string
	Recovered     string
	Active        string
	CombinedKey   string
}

type JHU struct {
	DataModel
	ChannelLine chan string
	ChannelOut  chan string
	S3Bucket    string
	S3key       string
	Wg          sync.WaitGroup
}

func (j *JHU) Munge(bucket string, key string) {
	start := time.Now()
	var builder strings.Builder
	/* producer */
	go scanner.ProcessDir(j.ChannelLine, j.S3Bucket, j.S3key, "flat")
	/* consumer */
	go reader.ReadObj(j.ChannelLine, j.ChannelOut)

	for line := range j.ChannelOut {
		if !strings.Contains(strings.Split(line, ",")[0], "FIPS") {
			j.Wg.Add(1)
			time.Sleep(200 * time.Microsecond)
			go func() {
				defer j.Wg.Done()
				builder.WriteString(fmt.Sprintf("%s%s", line, "\n"))
				/* checks to see if previous day file is processing */
				if strings.Contains(strings.Split(line, ",")[4], time.Now().Add(-time.Hour*24).Format("2006-01-02")) {
					utils.Info.Println(line)
				}
			}()
		}
	}
	j.Wg.Wait()
	aws.S3Obj{Bucket: bucket, Key: key,}.S3WriteGzip(builder.String(), aws.SessionGenerator())
	utils.Info.Println("Runtime took ", time.Since(start))
}
