test-by-ip:
	siege -r 1000 -c 10 http://localhost:8080/

test-api-key:
	siege -r 1000 -c 10 --header='API_KEY:123456' http://localhost:8080/
