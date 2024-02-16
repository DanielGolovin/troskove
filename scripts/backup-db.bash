#!/bin/bash

DB_PATH=~/troskove/volumes/sqlite/troskove.db
BACKUP_PATH=~/troskove/backups/sqlite
BACKUP_NAME=troskove_$(date +%Y-%m-%d_%H-%M-%S).db

mkdir -p $BACKUP_PATH
cp $DB_PATH $BACKUP_PATH/$BACKUP_NAME
find $BACKUP_PATH -name "troskove_*.db" -type f -mtime +30 -exec rm {} \;
