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

# runs spark package using scala in .jar
# http://biercoff.com/how-to-install-scala-on-mac-os/
cd ~/solutions/poly-spark-covid/
mvn clean package
scala -classpath "target/poly-spark-covid-1.0-production.jar" "com.poly.covid.Loader"
mvn clean
cd ~/