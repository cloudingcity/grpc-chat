.PHONY: client
client:
	go run main.go client -u Cloud -p 123456

.PHONY: client2
client2:
	go run main.go client -u Tifa -p 123456

.PHONY: server
server:
	go run main.go server --password 123456
