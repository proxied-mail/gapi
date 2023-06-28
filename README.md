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
        "password": "7Elrf2PZKsILfu38HS7J"
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