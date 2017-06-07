#!/bin/bash
rm boltdb/session.db
go build -o bin/commonwealth
bin/commonwealth
