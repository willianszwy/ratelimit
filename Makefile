test-by-ip:
	ab -t 60 -c 1 http://localhost:8080/

test-api-key:
	ab -H 'API_KEY:123456' -t 60 -c 1 http://localhost:8080/
