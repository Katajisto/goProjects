#!/bin/bash

git pull
go build
systemctl restart gitHook
