
all: run

run: src/*
	@bash run.sh

depend: src/*
	brew install mercurial
	go get code.google.com/p/goprotobuf/{proto,protoc-gen-go}
	go get -u github.com/christopherhesse/rethinkgo

