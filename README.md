# gapi

## Installation

```bash
make build2
make up
```




## Curls


### send mail
```bash
curl --location --request POST 'http://localhost:9900/gapi/send-mail' \
--header 'authority: proxiedmail.com' \
--header 'accept: application/json' \
--header 'accept-language: en-GB,en-US;q=0.9,en;q=0.8' \
--header 'authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIyIiwianRpIjoiMThjN2YxZmE3NzE2ZjE4NDE1MGU2NDY4MWQzNjhmYjY5N2VkODA0ZDVlNDgxMzk5NWZkYzNlY2NhOTg5MzZjN2ZiYmMyYjUwYjEwMDVmM2QiLCJpYXQiOjE2ODYxMzUzODQsIm5iZiI6MTY4NjEzNTM4NCwiZXhwIjoxNzE3NzU3NzgzLCJzdWIiOiI3Iiwic2NvcGVzIjpbXX0.DUmvhuF9Zi3pncxfCueuMS9A30cB1q2-_y2cMy0J_3VM0t8cDBYCR3f2fufSmZvdwp9AJMq1lQZ9kKZk5UodV8HjDcgvQH2rjTZeCVLi0pzNp8QNnHxdO9f-_yqG-rEtpk9CEWJnzBmFFLF62bHhcPKHZV_JW-dkidaA936n675ZAl6gCI102geDnNcDpDjjEWw1Vz4lcGGli-uV25NTHwW8wChO-yIDAEIhm-5jrBNhRfsuWh-CZybN5QSJzX86XnuvyEQjvclj-o7umCQ0g-vgNij29hBlL920sakmkpHNJdQYivht8kxTk27rUUdcO-zMR2aeHvkK7qq9GFwonXGW8DCSrSYU1p3WzILJ02FN8gNTs4jE709B9tKxeanmtJu75ReOaNyXPAUGh4s01eCgcWRu39HAQXJLlKYaXo4BpJy0JogqcOuLMT4TV51FbPCd-LVGVs-UjEmu5OSbLIacZ5dvaFvjfKL3xNlhqJTp9GwOHBNuzYbqZNFKvm6qIzLBRaM_023jk4rExxM9fzxevLw82-1oYMWoQo7YyKpnUgGPSrE4uoFs9ZiTNUZf5u3eKSaiBKs0Szj38fnc7qrkLBBCPo37aD07OTcHOiM2TvmLRyoHoZGCNQMC7V5H4eVq-R2ON51Ex86HStf5LXoempaXgdm8dKpsaFWCW5g' \
--header 'content-type: application/json' \
--header 'referer: https://proxiedmail.com/en/board' \
--header 'sec-ch-ua: "Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"' \
--header 'sec-ch-ua-mobile: ?0' \
--header 'sec-ch-ua-platform: "macOS"' \
--header 'sec-fetch-dest: empty' \
--header 'sec-fetch-mode: cors' \
--header 'sec-fetch-site: same-origin' \
--header 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36' \
--data-raw '{
    "auth":{
        "host": "mx.proxiedmail.com",
        "port": 587,
        "username":"catchall@emailsharevialink.com",
        "password": "not-valid-anymore-7Elrf2PZKsILfu38HS7J"
    },
    "mail": {
        "from": "catchall@emailsharevialink.com",
        "to":"webfay1@gmail.com",
        "subject": "Your pull request on Github",
        "type":"text/html",
        "body": "Please. If you'\''re authenticated - don'\''t go to spam"
    }
}'
```


### Add domain

```bash
curl --location --request POST 'http://localhost:9900/gapi/domains' \
--header 'authority: proxiedmail.com' \
--header 'accept: application/json' \
--header 'accept-language: en-GB,en-US;q=0.9,en;q=0.8' \
--header 'authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIyIiwianRpIjoiMThjN2YxZmE3NzE2ZjE4NDE1MGU2NDY4MWQzNjhmYjY5N2VkODA0ZDVlNDgxMzk5NWZkYzNlY2NhOTg5MzZjN2ZiYmMyYjUwYjEwMDVmM2QiLCJpYXQiOjE2ODYxMzUzODQsIm5iZiI6MTY4NjEzNTM4NCwiZXhwIjoxNzE3NzU3NzgzLCJzdWIiOiI3Iiwic2NvcGVzIjpbXX0.DUmvhuF9Zi3pncxfCueuMS9A30cB1q2-_y2cMy0J_3VM0t8cDBYCR3f2fufSmZvdwp9AJMq1lQZ9kKZk5UodV8HjDcgvQH2rjTZeCVLi0pzNp8QNnHxdO9f-_yqG-rEtpk9CEWJnzBmFFLF62bHhcPKHZV_JW-dkidaA936n675ZAl6gCI102geDnNcDpDjjEWw1Vz4lcGGli-uV25NTHwW8wChO-yIDAEIhm-5jrBNhRfsuWh-CZybN5QSJzX86XnuvyEQjvclj-o7umCQ0g-vgNij29hBlL920sakmkpHNJdQYivht8kxTk27rUUdcO-zMR2aeHvkK7qq9GFwonXGW8DCSrSYU1p3WzILJ02FN8gNTs4jE709B9tKxeanmtJu75ReOaNyXPAUGh4s01eCgcWRu39HAQXJLlKYaXo4BpJy0JogqcOuLMT4TV51FbPCd-LVGVs-UjEmu5OSbLIacZ5dvaFvjfKL3xNlhqJTp9GwOHBNuzYbqZNFKvm6qIzLBRaM_023jk4rExxM9fzxevLw82-1oYMWoQo7YyKpnUgGPSrE4uoFs9ZiTNUZf5u3eKSaiBKs0Szj38fnc7qrkLBBCPo37aD07OTcHOiM2TvmLRyoHoZGCNQMC7V5H4eVq-R2ON51Ex86HStf5LXoempaXgdm8dKpsaFWCW5g' \
--header 'content-type: application/json' \
--header 'referer: https://proxiedmail.com/en/board' \
--header 'sec-ch-ua: "Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"' \
--header 'sec-ch-ua-mobile: ?0' \
--header 'sec-ch-ua-platform: "macOS"' \
--header 'sec-fetch-dest: empty' \
--header 'sec-fetch-mode: cors' \
--header 'sec-fetch-site: same-origin' \
--header 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36' \
--data-raw '{
    "domain":"pxdmail.net"
}'
```


### Domain status 
```bash
curl --location --request GET 'https://proxiedmail.com/gapi/custom-domains?domain=abcddd.net' \
--header 'authority: proxiedmail.com' \
--header 'accept: application/json' \
--header 'accept-language: en-GB,en-US;q=0.9,en;q=0.8' \
--header 'authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIzIiwianRpIjoiZjg0ZTBmM2M0ZDI0OWUyNjU2MDViZWU3NTBiOTAyOTBjOTJjMjM3YWMwZTA3Mjk1NDJjYzRiOGZkNWI0NjM3YjcwNzhmZjY2ZGZhYTUwZTEiLCJpYXQiOjE2ODg4NTA1OTgsIm5iZiI6MTY4ODg1MDU5OCwiZXhwIjoxNzIwNDcyOTk3LCJzdWIiOiI5MDY2Iiwic2NvcGVzIjpbXX0.Q3xgZyWMtzgBTlyJfhsXhpNkJ6orXPFdWqh1XBGgkXXpegehv1IuRuItVj3gm-RyvuPspN5Vz86NVDmngR-VTNp3am-R7pRKw1uh2gtPTJswI8GoHgFQaztnpqsQlbEsbmxz-_LSxlKa2TnYVtLTmnkO3k9MvHCgRvD2CyUy4QYaaBN9K_TsioJOj1C5xX7u-KfnvYw6t8xfmE0owItJvTGtLuaXMxCyIvyxsHh7Bb_pr-6GqTdPhu5ZedBcVQc0y2fuSx2XzGU8BnYwDgZgbKC_Az--MMnphWYZMY2zvoMvk2Ap0Ncs6v8udKOcLXYaPzZ762AlSAmvRB1rs0eHiExwQzU8KRp_M3lGHz98T1NlVofNDv7mRm_nax4D8B4cFXrss1uKyy5OeJ9BVEhe9a-kbup2zhVbuMAto1qhb5fY3k5VWT-N3Tpk__H2Q8lhaa48mEV2JtA-F8FPWj25Gv9hGS6dp9Q6a1BFsA4wzVGRTiPB_4Ad52rb32-PXF-rv8Rwamh0l8lmnPYT1ZJ0Kl5gdFKP2XvVL4r3jkWDq-cqI1OjkxBuMIZEBe_USq793Fk2hNXl0QxM6NCgHaJp1ZK7ZhT_TGMx-nGvThY-JgkJW06kzepV3u1au6uxU0247vG99b2CoZXOVDoOnzZVFZV80mWyjZTP5tqcqDxX7JU' \
--header 'content-type: application/json' \
--header 'referer: https://proxiedmail.com/en/board' \
--header 'sec-ch-ua: "Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"' \
--header 'sec-ch-ua-mobile: ?0' \
--header 'sec-ch-ua-platform: "macOS"' \
--header 'sec-fetch-dest: empty' \
--header 'sec-fetch-mode: cors' \
--header 'sec-fetch-site: same-origin' \
--header 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36' \
--data-raw '{
    "domain":"abcddd.net"
}'
```



### Used on create

```
curl --location --request PATCH 'http://localhost:9900/gapi/used-on' \
--header 'authority: localhost:904' \
--header 'accept: */*' \
--header 'accept-language: en-GB,en-US;q=0.9,en;q=0.8' \
--header 'authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIyIiwianRpIjoiNjE1YjQ0NmMxZjUyMzZlMDFmMjIyYzc4N2Y5NzFiYTA4MGNlYTZiYjk2ZjZiZDI5ZjdkODNiMTJhMDAzNGM2MzhiODAzODk5MjJkNGYwZTEiLCJpYXQiOjE2OTY4NTE2MTAsIm5iZiI6MTY5Njg1MTYxMCwiZXhwIjoxNzI4NDc0MDA5LCJzdWIiOiI3Iiwic2NvcGVzIjpbXX0.boCXtkt5JUKSmeGpT7X95LZHqUPp8QHuqwMp4xsEpBXGE5pUDVGb4y2_OzW1-ywCVh0c_gDEmbv33HUJA2OCj-rw7yfyFh4vFHk_5npGHZZuWuH4JRL3zsjjFE4dmVFE-28eu0m5a6ayGwMBWvEOKbXA9tPZc9ZEPKu125jOtXbCdraXbsB1Dr22pmxW5nxSYQ06hTeraAY1icah4S0MPKdItzYqPLzf1eu-5_NdyeIsS1DJyG--bkEt89uxf_ZD4o6bNSnr-cldtf7NR1iL5Mn-_LGkSAWg1cRmXcDb98sSMRaSu3mSH4Uj6JXycQMFmk8xOJaV8Zs_JKPkyaJ-014AMlueXRdDm6z5OZaufO1hggzxJdzXXvz6VTaENidBpn1ZK5rJPfldIcb9WPyWsbl3prgj8ZtmFW4IsepZa_9TJEGNuIaTxhV8wS8Pw6TKSHVUFianJlEs3nLMUPe6R9q95oBRDlwoWEGbJyOFFxHTOIQF9kg87OtO7Ei6wm_4a76_0ie5AGdhCkQCEwLGl0IMEDJvOGasRqHrAUtM2MHNtN3uATzNcy9v0lgApgVDA7p3k-CGVRxrTqyO3HxUryeA5AKYfiStIy4YztWARnDKkTc2aqKKE0ZWF0FUC8JO8_R2RYlPGiO1C6Vtth1BM5L5A6x_c0ZQBSDU_Q_Mv50' \
--header 'content-type: application/json' \
--header 'cookie: _ga=GA1.1.1967188408.1637453570; _ga_0KY68T96T3=GS1.1.1637453583.1.0.1637454027.56; _ga_YCN8Z4TCV1=GS1.1.1645566552.41.0.1645566552.0; drift_aid=03899190-2f16-4628-9a90-011736c6f8ad; driftt_aid=03899190-2f16-4628-9a90-011736c6f8ad; wasAuthenticated=1; Phpstorm-a93d77e=6fc3f500-e2b5-4bed-9b66-170af6eb737d; fs_cid=1.0; _fbp=fb.0.1683802020203.90645981; _ga_M750FJFSKP=GS1.1.1685549399.132.1.1685552482.0.0.0; _ga_TV4KBG10CJ=GS1.1.1690450962.8.1.1690454212.0.0.0; pll_language=en; _tt_enable_cookie=1; _ttp=EHsGw57W1NNpDvhXryDab4jD2Vb; _ga_TG7D7SWNQ0=GS1.1.1690469178.1.1.1690469706.55.0.0; __stripe_mid=b3eb7bf8-2613-4ac4-a2cb-efbaf2f8b3bb59167d; _gcl_au=1.1.1904167837.1691248370.980981524.1696851289.1696851288; fs_lua=1.1696851289309; fs_uid=#TPK4D#22a4bea8-00b2-4800-8f17-dd13b1983000:f66b2bda-253a-4ae5-b1b8-88648d8b5371:1696851289309::1#/1718018934; pm_frontend_app_session=seNTjlbFB3dCfxgzaYhl3sHC3P2u3neztAWHGDE7; token=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIyIiwianRpIjoiNjE1YjQ0NmMxZjUyMzZlMDFmMjIyYzc4N2Y5NzFiYTA4MGNlYTZiYjk2ZjZiZDI5ZjdkODNiMTJhMDAzNGM2MzhiODAzODk5MjJkNGYwZTEiLCJpYXQiOjE2OTY4NTE2MTAsIm5iZiI6MTY5Njg1MTYxMCwiZXhwIjoxNzI4NDc0MDA5LCJzdWIiOiI3Iiwic2NvcGVzIjpbXX0.boCXtkt5JUKSmeGpT7X95LZHqUPp8QHuqwMp4xsEpBXGE5pUDVGb4y2_OzW1-ywCVh0c_gDEmbv33HUJA2OCj-rw7yfyFh4vFHk_5npGHZZuWuH4JRL3zsjjFE4dmVFE-28eu0m5a6ayGwMBWvEOKbXA9tPZc9ZEPKu125jOtXbCdraXbsB1Dr22pmxW5nxSYQ06hTeraAY1icah4S0MPKdItzYqPLzf1eu-5_NdyeIsS1DJyG--bkEt89uxf_ZD4o6bNSnr-cldtf7NR1iL5Mn-_LGkSAWg1cRmXcDb98sSMRaSu3mSH4Uj6JXycQMFmk8xOJaV8Zs_JKPkyaJ-014AMlueXRdDm6z5OZaufO1hggzxJdzXXvz6VTaENidBpn1ZK5rJPfldIcb9WPyWsbl3prgj8ZtmFW4IsepZa_9TJEGNuIaTxhV8wS8Pw6TKSHVUFianJlEs3nLMUPe6R9q95oBRDlwoWEGbJyOFFxHTOIQF9kg87OtO7Ei6wm_4a76_0ie5AGdhCkQCEwLGl0IMEDJvOGasRqHrAUtM2MHNtN3uATzNcy9v0lgApgVDA7p3k-CGVRxrTqyO3HxUryeA5AKYfiStIy4YztWARnDKkTc2aqKKE0ZWF0FUC8JO8_R2RYlPGiO1C6Vtth1BM5L5A6x_c0ZQBSDU_Q_Mv50; _ga_KNKJPRR29L=GS1.1.1696851288.253.1.1696851615.18.0.0' \
--header 'origin: https://localhost:904' \
--header 'referer: https://localhost:904/en/board' \
--header 'sec-ch-ua: "Google Chrome";v="117", "Not;A=Brand";v="8", "Chromium";v="117"' \
--header 'sec-ch-ua-mobile: ?0' \
--header 'sec-ch-ua-platform: "macOS"' \
--header 'sec-fetch-dest: empty' \
--header 'sec-fetch-mode: cors' \
--header 'sec-fetch-site: same-origin' \
--header 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36' \
--data-raw '{"proxy_binding_id":"A5653B00-0000-0000-00000BAE", "list": ["zalupa"]}'
```

Response:

```
{
    "status": true
}
```

### Used on GET

Same cURL as PATCH, but GET. Response:


```[
    {
        "proxy_binding_id": "BDF13B00-0000-0000-00000BAE",
        "list": [
            "zalupa"
        ]
    }
]
```

