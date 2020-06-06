package sources

import (
	"encoding/csv"
	a "github.com/polyglotDataNerd/poly-Go-utils/aws"
	u "github.com/polyglotDataNerd/poly-Go-utils/utils"
	h "net/http"
)

func ReadCSV() {
	resp, err := h.Get("https://coronadatascraper.com/timeseries.csv")
	if err != nil {
		u.Error.Println(err, "Could not reach API")
	}
	defer resp.Body.Close()
	csvr := csv.NewReader(resp.Body)
	result, cerr := csvr.ReadAll()
	for _, v := range result {
		if cerr != nil {
			u.Error.Println(cerr)
		}
		u.Info.Println(v)
	}

}

func GetCSV(bucket string, key string) {
	resp, err := h.Get("https://coronadatascraper.com/timeseries.csv")
	if err != nil {
		u.Error.Println(err, "Could not reach API")
	}
	defer resp.Body.Close()
	a.S3Obj{
		Bucket: bucket,
		Key:    key,
	}.S3UploadGzip(resp.Body, a.SessionGenerator())

}
