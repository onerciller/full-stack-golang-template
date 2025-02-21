
LOCAL_BIN := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/.bin
export PATH := ${LOCAL_BIN}:$(PATH)
