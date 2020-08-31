@echo off
del *.exe
del go.sum

go build -ldflags "-s -w"