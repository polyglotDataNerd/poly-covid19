# COVID19
**Tracks COVID19 Cases**

This project consolidates two of the major COVID19 data repositories and consolidates and standardizes two disparate sources.

Dependencies:
* [Corona Data Scraper](https://coronadatascraper.com/#home)
* [Johns Hopkins University](https://github.com/CSSEGISandData/COVID-19)
* [AWS S3](https://aws.amazon.com/s3/)
* [GoLang 1.14](https://golang.org/)
* [GO Utils library](https://github.com/polyglotDataNerd/zib-Go-utils)
* [Terraform](https://learn.hashicorp.com/terraform/getting-started/install.html)


Intention
-
* The intention of this repo is to clean and standardize a data model that can be put into visualization tools or pattern algorithms to find insight and trends for what's going on actively with the virus. 

Frequency
-  
* Data is updated daily but can be increased to real time depending on how fast this data gets to the API's 
    - Output of data can be found here:
        1. [Johns Hopkins University](output/JHU.csv)
        2. [Corona Data Scraper](output/CDS.csv)
        