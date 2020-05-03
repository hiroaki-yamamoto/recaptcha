#!/bin/env sh
# -*- coding: utf-8 -*-

set -e

go test -coverprofile=recaptcha.cov ./...
go tool cover -html=recaptcha.cov
