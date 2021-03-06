MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := zip
.ONESHELL:



GIT_COMMIT ?= $(shell git rev-list -1 HEAD)
BUILD_FLAGS := \
	-ldflags "-X github.com/mhristof/alfred-boxpn/cmd.GitCommit=$(GIT_COMMIT)"
PNG := $(subst svg,png,$(wildcard icons/*))

fast-test:  ## Run fast tests
	go test ./... -tags fast

test:	## Run all tests
	go test ./...

alfred-boxpn: $(shell find ./ -name '*.go')
	go build $(BUILD_FLAGS) -o alfred-boxpn main.go

rand: alfred-boxpn
	./alfred-boxpn $(shell ./alfred-boxpn alfred  | jq '.items[].arg' -r | grep openvpn | shuf -n1)

zip: alfred-boxpn info.plist
	zip -r alfred-boxpn.alfredworkflow FBD01B59-817D-4289-92FB-937ABBD14F70.png boxpn-openvpn-configs info.plist alfred-boxpn alfred-boxpn.sh icons icon.png

.PHONY: fix
fix:  $(PNG)
	echo '$(PNG)'

%.png:
	svg2png '$(subst png,svg,$@)' > '$@'
	rm -rf '$(subst png,svg,$@)'

go-fix:
	go run main.go fix

minor:
	sed -i "" 's/$(shell semver current | tr -d 'v' )/$(shell semver -n | rev | cut -d ' ' -f1 | rev | tr -d 'v')/' info.plist
	make commitVersion
	semver
	git push --tags

patch:
	sed -i "" 's/$(shell semver current | tr -d 'v' )/$(shell semver -p -n | rev | cut -d ' ' -f1 | rev | tr -d 'v')/' info.plist
	make commitVersion
	semver --patch
	git push --tags

.PHONY: commitVersion
commitVersion:
	git add info.plist
	git commit -m 'bumped version' info.plist


.PHONY:
help:           ## Show this help.
	@grep '.*:.*##' Makefile | grep -v grep  | sort | sed 's/:.* ##/:/g' | column -t -s:

