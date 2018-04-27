generate-all: generate-raw-types generate-kubernetes-interfaces

clean-generation:
	rm -r pkg/apis/aws/v1/aws_*

generate-raw-types:
	cd type_generator && \
		go run main.go
generate-kubernetes-interfaces:
	./generate-ii.sh