#!/bin/bash

git -C /root/go/src/goProjects/ pull
snap run go build goProjects
systemctl restart gitHook
