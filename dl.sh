#!/bin/sh
#export http_proxy=http://127.0.0.1:7890
#export https_proxy=http://127.0.0.1:7890
go build .
sudo ./dlm3u8 -i $1
sudo chown -R jellyfin:jellyfin /home/jellyfin/films
