windres.exe -i app.rc -o defaultRes_windows_amd64.syso

go build -tags tempdll -ldflags="-s -w -H windowsgui" -o bin/wawa_pitaya.exe .

pause