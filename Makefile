build_and_install:
	go build -o finder .
	sudo mv finder /usr/bin/finder