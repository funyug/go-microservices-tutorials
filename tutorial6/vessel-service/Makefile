build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/funyug/go-microservices-tutorials/tutorial6/vessel-service \
    proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .

run:
	docker run -d --net="host" \
        		-p 50053 \
        		-e MICRO_SERVER_ADDRESS=:50053 \
        		-e MICRO_REGISTRY=mdns \
        		-e DISABLE_AUTH=true \
        		vessel-service