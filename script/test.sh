#!/bin/bash

for i in {1..10}; do curl localhost:2112/fast_response && echo; done
curl localhost:2112/slow_response && echo
curl localhost:2112/error_response && echo
