package sources

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	a "github.com/polyglotDataNerd/poly-Go-utils/aws"
	u "github.com/polyglotDataNerd/poly-Go-utils/utils"
	"io/ioutil"
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

func GetZip(bucket string, key string) {
	resp, err := h.Get("https://coronadatascraper.com/timeseries.csv.zip")
	if err != nil {
		u.Error.Println(err, "Could not reach API")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	czip, cerr := zip.NewReader(bytes.NewReader(body), resp.ContentLength)
	if cerr != nil {
		u.Error.Println(cerr)
	}

	for _, f := range czip.File {
		src, serr := f.Open()
		if serr != nil {
			u.Error.Println(serr)
		}
		defer src.Close()
		a.S3Obj{
			Bucket: bucket,
			Key:    key,
		}.S3UploadGzip(src, a.SessionGenerator())
	}
}
