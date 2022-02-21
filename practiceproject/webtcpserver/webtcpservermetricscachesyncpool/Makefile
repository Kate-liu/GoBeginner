all: server client

server: cmd/server/main.go
	go build github.com/Kate-liu/GoBeginner/practiceproject/webtcpserver/webtcpservermetricscachesynpool/cmd/server
client: cmd/client/main.go
	go build github.com/Kate-liu/GoBeginner/practiceproject/webtcpserver/webtcpservermetricscachesynpool/cmd/client

clean:
	rm -fr ./server
	rm -fr ./client
