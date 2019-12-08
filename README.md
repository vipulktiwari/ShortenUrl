# ShortenUrl
A URL shortening service written in go.


Exposes 2 apis CreateURL and AccessURL

Example Output : 

Vipuls-MacBook-Pro:ShortenUrl vipul$ curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" http://127.0.0.1:8080/createurl -d '{"url":"www.fb.com"}'
"OKUK68tuS8wH"
Vipuls-MacBook-Pro:ShortenUrl vipul$ curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" http://127.0.0.1:8080/accessurl -d '{"url":"OKUK68tuS8wH"}'
"www.fb.com"
