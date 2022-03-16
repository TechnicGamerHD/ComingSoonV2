# ComingSoon

Install on Linux with:
git clone  

Create Service File in /etc/systemd/system

Start on Port 20800 with:
`systemctl start comingsoon.service`

If you want to change the port, edit main.go on line 36. 
You have to rebuild the executable, just isntall go on linux, switch to /home/ComingSoon and run:
`GOOS=linux go build main.go`


# For Windows
Just download the comingsoon.exe and run it via console
`./comingsoon.exe`