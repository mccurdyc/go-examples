#!/bin/bash
# A one-liner for installing Go (and setting up the necessary environment variables)
# Credit: Steve Francia, Jess Frazelle and Chris Broadfoot
# Google Group Post: https://groups.google.com/forum/#!topic/golang-nuts/ZFqZkGYX4x8

curl -LO https://get.golang.org/$(uname)/go_installer && \
chmod +x go_installer && \
./go_installer && \
rm go_installer
