package azopenai

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func streamRequest(req *policy.Request, err error) (*policy.Request, error) {
	req.Raw().Header["Accpet"] = []string{"text/event-stream"}
	runtime.SkipBodyDownload(req) // skip downloading the body since we're streaming
	return req, err
}
