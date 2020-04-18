#!/usr/bin/env bash
# local RUN
CURRENTDATE="$(date -v -1d +%m-%d-%Y)"
cd ~/external/COVID-19/csse_covid_19_data/csse_covid_19_daily_reports/
git pull
aws s3 cp ~/external/COVID-19/csse_covid_19_data/csse_covid_19_daily_reports/"$CURRENTDATE".csv s3://poly-testing/covid/jhu/raw/"$CURRENTDATE".csv --sse

# runs go package
cd ~/solutionsgo/poly-covid19/
go run Entry.go
cd ~/

# runs spark package (.jar)
cd cd ~/solutions/poly-spark-covid/
mvn clean package
java -cp poly-spark-covid-1.0-production.jar com.poly.covid.Loader