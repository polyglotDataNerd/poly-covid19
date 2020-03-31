package main

import (
	"fmt"
	"github.com/polyglotDataNerd/poly-covid19/sources"
	uuid "github.com/satori/go.uuid"
	"time"
)

func main() {
	sources.GetCSV("poly-testing", fmt.Sprintf("%s%s%s%s%s", "covid/cds/", time.Now().Format("2006-01-02"), "/", uuid.NewV4().String(), ".gz"))
}
