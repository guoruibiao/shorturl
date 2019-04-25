package model

type Request struct {
	OriginURL string
}
/*
{
    "urls": [
        {
            "result": true,
            "url_short": "http://t.cn/RxnlTYR",
            "url_long": "https://github.com",
            "object_type": "",
            "type": 0,
            "object_id": ""
        }
    ]
}
*/
type ResponseEntry struct {
	Success bool `json:"result"`
	ShortURL string `json:"url_short"`
	OriginURL string `json:"url_long"`
}

type Response struct {
	URLS []ResponseEntry
}
