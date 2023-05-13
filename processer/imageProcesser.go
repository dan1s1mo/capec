package processer

import (
	"capec/types"
	"capec/utils"
	"fmt"
)

type Imageprocesser interface {
	ProcessImage(path string) [](types.Box)
}

type HTTPImageProcesse struct{}

func (hip *HTTPImageProcesse) ProcessImage(path string) [](types.Box) {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	response := utils.MakeHttpRequest[[]types.Box](
		"http://localhost:8080/image",
		fmt.Sprintf(`{ "path": %q }`, path),
		"POST",
		headers,
	)
	return *response
}
