module github.com/xuxife/azopenai

go 1.18

require github.com/Azure/azure-sdk-for-go/sdk/azcore v1.4.0

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.2.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)

replace github.com/Azure/azure-sdk-for-go/sdk/azcore v1.4.0 => github.com/xuxife/azure-sdk-for-go/sdk/azcore v0.0.0-20230323180737-52abf5d35916
