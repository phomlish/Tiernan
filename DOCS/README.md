I use VS Code

# BUILD

To clone the git repository\
    git clone https://github.com/phomlish/Tiernan.git\

Navigate into tierman directory\
make sure the packages are up to date\
    go mod tidy\

To build the application tiernan.exe\
    go make \

# INSTALL

install the package creator from https://jrsoftware.org/isinfo.php
run INNO and open SUPPORT/t inno setup.iss
press PLAY in INNO and follow the prompts
Ignore the Support/Output directory created by the INNO installer

The installer package will install the application for WIndows. 
The application will start-on-boot.

Program files are installed at 
    C:\Program Files\Tiernan
    C:\Program Files\Tiernan\tiernan.exe
    C:\Program Files\Tiernan\unins000.dat
    C:\Program Files\Tiernan\unins000.exe

Auto start is enabled with registry entry at 
    Computer\HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run\Tiernan

# TESTING

go test

sample output
PS C:\dev\Tiernan> go test           
{"level":"info","time":"2024-11-05T14:33:59-05:00","message":"result true 11.9472ms"}
{"level":"info","time":"2024-11-05T14:33:59-05:00","message":"result false i1w fe80::9fe1:40f2:9cbc:b6fd%Ethernet\nfe80::c322:fecd:f6b:a8e4%vEthernet (WSL (Hyper-V firewall))\nfe80::66d2:c9c0:19cb:a6fd%vEthernet (Default Switch)\n10.11.1.51\n172.27.96.1\n172.25.80.1\n"}
PASS
ok      tiernan 0.234s

# TESTING WITH POSTMAN

open the postman file Tiernan.postman_collection.json in postman
make sure the app is running and try the requests:
    Ping
    Sysinfo

# App In Action

Although I didn't make an annimation, I did put some screenshots in the 'App In Action' directory

# Additional Notes

Ideas I had while devloping are in DOCS/TODO.md
I had questions about the implementation and put the questions and how I handled them in DOCS/Decisions.md

