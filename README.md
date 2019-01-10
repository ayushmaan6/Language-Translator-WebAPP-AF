# Create a simple Language Translator Web-App using Go and Watson Language Translator service

[![](https://img.shields.io/badge/IBM%20Cloud-powered-blue.svg)](https://bluemix.net)
![Platform](https://img.shields.io/badge/platform-go-lightgrey.svg?style=flat)

### Table of Contents

* [Summary](#summary)
* [Requirements](#requirements)
* [Configuration](#configuration)
* [Deploy](#Deploy)

<a name="summary"></a>
### Summary

This is a Web Application which is built on GO programming language to translate text in english language to other supported by the IBM Watson Language Translator service. 


#### 1. Architecture

#### 2. Flow

#### 3. Included Components
*[Watson Language Translator](https://www.ibm.com/cloud/watson-language-translator): Language Translator translates text from one language to another. Take news from across the globe and present it in your language. Communicate with your customers in their own language, and more.

#### 4. Included Technologies
*[Go](https://golang.org/): Go is an open source programming language that makes it easy to build simple, reliable, and efficient software
*HTML & CSS

<a name="requirements"></a>
### Requirements
#### IBM Cloud tools setup (optional)

Install [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/ibmcloud/download_cli.html#install_use) 



<a name="configuration"></a>
### Configuration

The project contains IBM Cloud specific files that are used to deploy the application as part of an IBM Cloud.

#### 1. Clone the repo
Clone the Language-Translator-WebAPP-Using-GO locally. In a terminal, run:

'git clone https://github.com/ayushmaan6/Language-Translator-WebAPP-Using-GO.git'

#### 2. Create Watson services with IBM Cloud
Create the following services:

* [Language Translator](https://console.bluemix.net/catalog/services/language-translator)

#### 3. Configure credentials
Collect the credentials for the IBM Cloud Language Translator service.

Find the service in your IBM Cloud Dashboard.
Click on the service.
Hit Manage in the left sidebar menu.
Copy the API Key and URL.


Change to the directory where your code is located.
Edit the server.go file Line 71 & 73 with the API key and the URL.

```
cd Language-Translator-WebAPP-Using-GO
nano server.go
``` 
### Deploy





