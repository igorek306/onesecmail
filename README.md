# 1SECMAIL simple intergration for go
this package isn't official, official documentation you can find on [www.1secmail.com/api](https://www.1secmail.com/api/)

## Installation:
```bash
go get github.com/igorek306/onesecmail
```
## Usage:
```go
package main

import "github.com/igorek306/onesecmail"

func main() {
	// 1secmail client initialization
	client := onesecmail.NewClient()

	// generating random email addresses
	addresses, err := client.GenerateRandomEmailAddresses(10)

	// getting list of active domains
	domains, err := client.GetAllActiveDomains()

	// checking your mailbox
	messages, err := client.CheckMailbox("example@vjuum.com")

	//fetching single message
	detailed_message, err := client.ReadEmail("example@vjuum.com", 123)

	//getting attachment download url
	url, err := client.DownloadAttachmentUrl("example@vjuum.com", 123, "image.jpg")
}

```
