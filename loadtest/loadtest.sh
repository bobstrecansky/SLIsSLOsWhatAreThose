#!/bin/bash
hey -n 1000 -c 5 http://goapp:2112/fast_response
hey -n 100 -c 1 http://goapp:2112/error_response
