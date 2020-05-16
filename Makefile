.PHONY: client
client:
	go run main.go client -u "John Doe" -p 123456

.PHONY: server
server:
	go run main.go server --password 123456
