BUILD_NAME=ddns
DIST_DIR=dist

install: clean build-osx build-linux build-win
	mkdir -p ${DIST_DIR}/scripts  && mkdir -p ${DIST_DIR}/conf
	cp -r conf/* ${DIST_DIR}/conf
	chmod +x ${DIST_DIR}/scripts/* && chmod -x ${DIST_DIR}/scripts/*.exe
	cd ${DIST_DIR} && zip -r ddns-golang.zip ./*
	mv ${DIST_DIR}/ddns-golang.zip ${DIST_DIR}/../
#	rm -rfv ${DIST_DIR}/conf/*
#	rm -rfv ${DIST_DIR}/scripts/${BUILD_NAME}_*
	rm -rfv ${DIST_DIR}

build-osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${DIST_DIR}/scripts/${BUILD_NAME}_darwin_amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ${DIST_DIR}/scripts/${BUILD_NAME}_darwin_arm64

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=386   go build -o ${DIST_DIR}/scripts/${BUILD_NAME}_linux_386
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${DIST_DIR}/scripts/${BUILD_NAME}_linux_amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm   go build -o ${DIST_DIR}/scripts/${BUILD_NAME}_linux_arm
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${DIST_DIR}/scripts/${BUILD_NAME}_linux_arm64

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=386   go build -o ${DIST_DIR}/scripts/${BUILD_NAME}_386.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${DIST_DIR}/scripts/${BUILD_NAME}_amd64.exe

clean:
#	mkdir -p ${DIST_DIR}/scripts  && mkdir -p ${DIST_DIR}/conf
#	rm -rfv ${DIST_DIR}/conf/*
#	rm -rfv ${DIST_DIR}/scripts/${BUILD_NAME}_*
	rm -rfv ${DIST_DIR}
	rm -rf  ./ddns-golang.zip
