#!/usr/bin/expect
geth account new
cd ~/.ethereum/keystore/
echo "Your recovery file :"
ls -t UTC* | head -1
