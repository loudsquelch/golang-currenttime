#!/bin/sh
curl -i http://localhost:8000/api/time
echo ''
# HTTP/1.1 200 OK
# Content-Type: application/json
# Date: Mon, 31 Jan 2022 10:11:35 GMT
# Content-Length: 49

# {"current_time":"2022-01-31 10:11:35 +0000 UTC"}

curl -i http://localhost:8000/api/time?tz=US/Aleutian
echo ''
# HTTP/1.1 200 OK
# Content-Type: application/json
# Date: Mon, 31 Jan 2022 10:11:49 GMT
# Content-Length: 49

# {"current_time":"2022-01-31 00:11:49 -1000 HST"}

curl -i http://localhost:8000/api/time?tz=US/Aleutian,GMT
echo ''
# HTTP/1.1 200 OK
# Content-Type: application/json
# Date: Mon, 31 Jan 2022 10:12:55 GMT
# Content-Length: 86

# {"GMT":"2022-01-31 10:12:55 +0000 GMT","US/Aleutian":"2022-01-31 00:12:55 -1000 HST"}

curl -i http://localhost:8000/api/time?tz=US/Aleutian,NONSENSE
echo ''
# HTTP/1.1 400 Bad Request
# Content-Type: text/plain; charset=utf-8
# X-Content-Type-Options: nosniff
# Date: Mon, 31 Jan 2022 10:12:49 GMT
# Content-Length: 17

# invalid timezone