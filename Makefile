# https://github.com/kirby-101/dpxlabd.git

VERSION			!= git describe --tags --always --dirty
GO				!= which go
GO_ENV			= CGO_ENABLED=0
GO_FLAGS		:= -trimpath -ldflags "-s -w -X main.Version=${VERSION}"

DIST_DIR		= ./.dist
DIST_BIN		:= ${DIST_DIR}/dpxlabd
DIST_RC			= ./dpxlabd.rc.sh
DIST_CFG		= ./dpxlabd.yml

STAGE_DIR		:= ${DIST_DIR}/stage
STAGE_PREFIX	= /usr/local
STAGE_PLIST		:= ${STAGE_DIR}/+PLIST
STAGE_MANIFEST	:= ${STAGE_DIR}/+MANIFEST
STAGE_BIN		:= ${STAGE_DIR}${STAGE_PREFIX}/bin/dpxlabd
STAGE_RC		:= ${STAGE_DIR}${STAGE_PREFIX}/etc/rc.d/dpxlabd
STAGE_CFG		:= ${STAGE_DIR}${STAGE_PREFIX}/etc/dpxlabd.yml


.PHONY: all

all: clean test build stage plist manifest package

clean:
	rm -rf ${DIST_DIR}/*

test:
	${GO_ENV} ${GO} test -v ./...

build: clean test
	mkdir -p ${DIST_DIR}
	@echo "[Building] ..."
	${GO_ENV} ${GO} build ${GO_FLAGS} -o ${DIST_BIN} ./main/

stage: build
	mkdir -p ${STAGE_DIR}${STAGE_PREFIX}/bin
	install -m 755 ${DIST_BIN} ${STAGE_BIN}
	mkdir -p ${STAGE_DIR}${STAGE_PREFIX}/etc/rc.d
	install -m 755 ${DIST_RC} ${STAGE_RC}
	cp ${DIST_CFG} ${STAGE_CFG}

plist: stage
	@cd ${STAGE_DIR}${STAGE_PREFIX} && find . -type f -print > ../../+PLIST

manifest: plist
	@echo 'name: dpxlabd' > ${STAGE_MANIFEST}
	@echo 'version: "${VERSION}"' >> ${STAGE_MANIFEST}
	@echo 'origin: kirby-101/dpxlabd' >> ${STAGE_MANIFEST}
	@echo 'comment: "dpxlabd"' >> ${STAGE_MANIFEST}
	@echo 'desc: "dpxlabd"' >> ${STAGE_MANIFEST}
	@echo 'maintainer: kirby.41@signal.org' >> ${STAGE_MANIFEST}
	@echo 'www: https://github.com/kirby-101/dpxlab' >> ${STAGE_MANIFEST}
	@echo 'prefix: /usr/local' >> ${STAGE_MANIFEST}
	@echo 'licenselogic: single' >> ${STAGE_MANIFEST}
	@echo 'licenses: [MIT]' >> ${STAGE_MANIFEST}
	@echo 'arch: "FreeBSD:15:amd64"' >> ${STAGE_MANIFEST}
	@echo 'deps: {}' >> ${STAGE_MANIFEST}
	@echo 'scripts: {' >> ${STAGE_MANIFEST}
	@echo '  post-install: "true"' >> ${STAGE_MANIFEST}
	@echo '}' >> ${STAGE_MANIFEST}

package: manifest stage
	pkg create -v -M ${STAGE_MANIFEST} -r ${STAGE_DIR} -o ${DIST_DIR} -p ${STAGE_PLIST}

install:
	pkg add ${DIST_DIR}/dpxlabd-*.pkg

push-bullshit:
	git add .
	git commit
	git push

pull-bullshit:
	git pull

run-bullshit: pull-bullshit build
	./.dist/dpxlabd
