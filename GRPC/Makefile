gen:
	protoc --proto_path=./proto/ --go-grpc_out=pb/ --go_out=pb/ ./proto/*.proto
	protoc --proto_path=./proto/ --go-grpc_out=pb/ ./proto/*.proto
	
clean:
	rm pb/*.go	

run:
	go run main.go	