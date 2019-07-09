#!/bin/bash

git -C /root/go/src/goProjects/ pull
snap run go build goProjects
mv /goProjects /root/
systemctl restart gitHook
