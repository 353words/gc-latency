#!/bin/bash

ls -l users | awk '{print $NF}' | sed -e 's/users_//'
