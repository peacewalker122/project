#!/bin/sh
# this script is used to start the ent 

go run -mod=mod entgo.io/ent/cmd/ent init --target ./db/ent/schema $1
