# /bin/bash

ps aux | grep go | awk '{print $2}' | xargs kill -9