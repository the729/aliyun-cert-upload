package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cas"
)

func main() {
	certFile := flag.String("cert", "", "Path to the certificate file")
	keyFile := flag.String("key", "", "Path to the private key file")
	name := flag.String("name", "", "Name of the certificate, overrides base-name")
	basename := flag.String("base-name", "", "Name of the certificate will be generated as <base-name>-YYYYMMDD")
	resourceGroupID := flag.String("resource-group-id", "", "Resource group ID (optional)")
	flag.Parse()

	client, err := cas.NewClientWithAccessKey("cn-hangzhou", os.Getenv("Ali_Key"), os.Getenv("Ali_Secret"))
	if err != nil {
		panic(err)
	}

	req := cas.CreateUploadUserCertificateRequest()

	cert, err := os.ReadFile(*certFile)
	if err != nil {
		fmt.Println("Error reading certificate file:", err)
		os.Exit(1)
	}

	key, err := os.ReadFile(*keyFile)
	if err != nil {
		fmt.Println("Error reading private key file:", err)
		os.Exit(1)
	}
	req.Cert = string(cert)
	req.Key = string(key)

	if *name != "" {
		req.Name = *name
	} else if *basename != "" {
		req.Name = fmt.Sprintf("%s-%s", *basename, time.Now().Format("20060102"))
	} else {
		fmt.Println("Either name or base-name is required")
		os.Exit(1)
	}

	req.ResourceGroupId = *resourceGroupID

	response, err := client.UploadUserCertificate(req)
	if err != nil {
		fmt.Println("Error uploading certificate:", err)
		os.Exit(2)
	}
	fmt.Print(response.GetHttpContentString())
}
