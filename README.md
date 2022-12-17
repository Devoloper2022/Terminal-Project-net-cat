# Net-cat

![image](https://www.myfreax.com/content/images/2022/11/linux-nc-command-send-directory.webp)
The pproject consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

   - NetCat, nc system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open TCP connections, send UDP packages, listen on arbitrary TCP and UDP ports and many more.

  -  To see more information about NetCat inspect the manual man nc.



![image](https://highload.today/wp-content/uploads/2021/12/golang.jpeg)

## Program architecture

- service  // Back-end (server)
	- function //  additional function 
    - handle // main 2 function, core of the program
    - module // all structs which use program
    - server // for lounch program 
- main.go // input point for hole program 

---

## Program architecture
Run program: "go run . [IP]" or "go run . " by default 
Example 1:go run . 2525
Example 2:go run .

