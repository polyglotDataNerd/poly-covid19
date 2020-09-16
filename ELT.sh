#!/usr/bin/env bash
# local RUN
# CURRENTDATE="$(date -v -1d +%m-%d-%Y)"
CURRENTDATE="$(date +%m-%d-%Y -d "$1 days ago")"
cd ~/external/COVID-19/csse_covid_19_data/csse_covid_19_daily_reports/
git pull
aws s3 cp ~/external/COVID-19/csse_covid_19_data/csse_covid_19_daily_reports/"$CURRENTDATE".csv s3://poly-testing/covid/jhu/raw/"$CURRENTDATE".csv --sse

# runs go package
cd ~/poly-covid19/
# run with profile and tracing
GODEBUG=gctrace=1,scheddetail=1,schedtrace=1000 go run Entry.go
# https://golang.org/doc/articles/race_detector.html
# go run -race Entry.go
cd ~/

# run EMR spark job
# source ~/poly-spark-covid/EMR.sh
# cd ~/

AWS_ACCESS_KEY_ID=$(aws ssm get-parameters --names /s3/polyglotDataNerd/admin/AccessKey --query Parameters[0].Value --with-decryption --output text)
AWS_SECRET_ACCESS_KEY=$(aws ssm get-parameters --names /s3/polyglotDataNerd/admin/SecretKey --query Parameters[0].Value --with-decryption --output text)


# https://aws.amazon.com/about-aws/whats-new/2020/04/amazon-emr-now-available-los-angeles/
aws emr create-cluster --profile default --security-configuration EMRSecConfig --applications Name=Spark --tags 'VPC=Default' 'name=poly-covid' 'Environment=production' 'Name=poly-covid' --ec2-attributes '{"KeyName":"gerard-bartolome-dev","AdditionalSlaveSecurityGroups":["sg-933774ce"],"InstanceProfile":"admin","SubnetId":"subnet-01c6a60a78a29d1da","EmrManagedSlaveSecurityGroup":"sg-933774ce","EmrManagedMasterSecurityGroup":"sg-933774ce","AdditionalMasterSecurityGroups":["sg-933774ce"]}' --release-label emr-5.30.0 --log-uri 's3n://poly-hadoop/' --instance-groups '[{"InstanceCount":1,"EbsConfiguration":{"EbsBlockDeviceConfigs":[{"VolumeSpecification":{"SizeInGB":20,"VolumeType":"io1","Iops":1000},"VolumesPerInstance":1}],"EbsOptimized":true},"InstanceGroupType":"MASTER","InstanceType":"c5.xlarge","Name":"Master - 1"},{"InstanceCount":3,"EbsConfiguration":{"EbsBlockDeviceConfigs":[{"VolumeSpecification":{"SizeInGB":20,"VolumeType":"io1","Iops":1000},"VolumesPerInstance":3}]},"InstanceGroupType":"CORE","InstanceType":"c5.4xlarge","Name":"Core - 2"}]' --configurations '[{"Classification":"spark","Properties":{"authenticate":"true","maximizeResourceAllocation":"true"}},
{"Classification":"emrfs-site",
"Properties": {
    "fs.s3.serverSideEncryptionAlgorithm":"AES256",
    "fs.s3.enableServerSideEncryption":"true",
    "fs.s3a.access.key":"'"$AWS_ACCESS_KEY_ID"'",
    "fs.s3a.secret.key":"'"$AWS_SECRET_ACCESS_KEY"'"
  }
}]' --service-role EMR_DefaultRole --enable-debugging --name 'poly-covid' --scale-down-behavior TERMINATE_AT_TASK_COMPLETION --region us-west-2 \
--steps \
Type=Spark,Name=Spark-Converter,ActionOnFailure=CONTINUE,Args=[--class,com.poly.covid.Loader,s3n://bigdata-utility/spark/poly-spark-covid-1.0-production.jar] \
--auto-terminate

# runs spark package using scala in .jar
# http://biercoff.com/how-to-install-scala-on-mac-os/
#cd ~/solutions/poly-spark-covid/
#mvn clean package
#spark-submit --class com.poly.covid.Loader \
#--master local["*"] \
#target/poly-spark-covid-1.0-production.jar
#mvn clean
#cd ~/
