<h1 align="center">Golang group listeners and one sender for Artemis and protocol STOMP</h1>

## Description
This program in a separate goroutine listens to messages using the stomp protocol and sends messages synchronously. The listener prints received messages.

## Deploy Artemis in Docker
```
docker run -d --name artemis -p 8161:8161 -p 61616:61616 vromero/activemq-artemis
```

## Install library github.com/go-stomp/stomp
```
go get github.com/go-stomp/stomp
```

## Execution result
![image](https://user-images.githubusercontent.com/59051004/219353627-1fbf28a0-0476-49a4-9bee-514539dc2471.png)
