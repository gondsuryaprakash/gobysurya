package https

import (
	"fmt"
	"net/http"
	"time"
)

func CallApi(url, method string) error {

	client := &http.Client{
		Timeout: time.Second * 100,
	}

	req, _ := http.NewRequest(method, url, nil)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	contentTypeValues := resp.Header.Values("content-type")
	fmt.Println(contentTypeValues)

	contentTypeValue := resp.Header.Get("content-type")
	fmt.Println("ContentTypeValue", contentTypeValue)

	headers := resp.Header
	fmt.Println("Headers", headers)
	return nil

}
