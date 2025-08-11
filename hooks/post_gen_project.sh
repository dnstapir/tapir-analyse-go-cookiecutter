#!/bin/bash
git init -b main

git config --local user.name {{cookiecutter.user}}
git config --local user.email {{cookiecutter.email}}

git add .
git commit -m "Initial commit using template revision {{cookiecutter.__revision}} of github.com/dnstapir/tapir-analyse-go-cookiecutter"
