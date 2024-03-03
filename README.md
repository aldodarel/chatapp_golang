# _Simple ChatApp with Database SQL in Golang_



## Tech
- XAMPP   (for SQL Database in phpmyadmin)
- Golang [![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)
go1.21.6 windows/amd64
- Code Editor ( VSCode [![Visual Studio Code](https://img.shields.io/badge/--007ACC?logo=visual%20studio%20code&logoColor=ffffff)](https://code.visualstudio.com/) )
- Browser (Chrome) 



## Installation
-> Clone my github to your local
```sh
git clone https://github.com/aldodarel/chatapp_golang.git
```

-> Open your Code Editor VSCode, ensure the folder structure:

📂simple_chatapp

 ┃ ┣ 📜Client.go

 ┃ ┣ 📜Server.go
 
 ┃ ┣ 📜go.mod
 
 ┃ ┣ 📜go.sum


-> Open your XAMPP and start the Apache and MySQL module

-> Open your phpmyadmin

-> Create a new database in phpmyadmin, if me , ex: "chatapp_db"

-> Run this query on your chatapp_db database:
```sh
CREATE TABLE messages ( id INT AUTO_INCREMENT PRIMARY KEY, message TEXT NOT NULL );
```

-> So there are you've made 1 table 'messages' with 2 column 'id' and 'message'

-> Save it

-> Back and demo the chatapp program, open your terminal and ensure the directory on the right path.

-> Open the first Command Prompt/terminal, type & run this command: 
```sh
go run Server.go
```

-> Waiting and if you see administrator prompt with .exe, just click Cancel.

-> Then, there is a message 
```sh
"Server mendengarkan pada Client..."
```

-> Open the second Command Prompt/terminal, type & run this command: 
```sh
go run Client.go
```

-> Wait for the reply:
```sh
"Masukkan pesan: "
```
If there is, so enter your message.

-> If you have done, there is message like: 
```sh
"Pesan terkirim: " 
```
It means the messages have delivered

-> Go to the first terminal and check to ensure the messages saved to your database. If delivered successfully, there is feedback: 
```sh
"Pesan disimpan: "
```

-> Check your database table and if it's success, your message will appear there.

```sh

```


** THANK YOU **

**KEEP CODING AND JUST TRY IT! 🔥🔥**



~ If you have any problem, just catch to me on darellaldo2004@gmail.com




