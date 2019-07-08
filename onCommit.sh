#!/bin/bash

git pull
snap run go build
systemctl restart gitHook
