run-go:
	go run main.go
	
run-web:
	cd web && yarn && yarn build &&yarn start
