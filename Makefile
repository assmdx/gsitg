.PHONY: test
test:
	rm -rf test/
	mkdir test/
	cd test && git clone https://github.com/assmdx/gsitg.git
	rm -rf ./test/gsitg/.git
	cd dep && go test