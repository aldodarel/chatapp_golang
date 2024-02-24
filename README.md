# chatapp_golang
Simple ChatApp with Database SQL in Golang

Tools you need before start:
1. XAMPP (for SQL Database in phpmyadmin)
2. Golang
3. Code Editor VSCode
   

**How to Setup?**

1. git clone https://github.com/aldodarel/chatapp_golang.git
2. Open your Code Editor VSCode, ensure there are files: Client.go, Server.go, go.mod, and go.sum
3. Open your XAMPP and start the Apache and MySQL module
4. Open your phpmyadmin
5. Create a new database, if me , ex: "chatapp_db"
   
6. Run this query on your chatapp_db database:
   CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    message TEXT NOT NULL
   );

7. So there are you make 1 table 'messages' with 2 column 'id' and 'message'
8. Save it.
9. Back and demo the chatapp program, open your terminal and ensure the directory on the right path.
10. Open the first Command Prompt/terminal and run this: 'go run Server.go'
11. Waiting and if you see administrator prompt with .exe, just click Cancel.
12. Then, there is a message "Server mendengarkan pada Client..."
13. Open the second Command Prompt/terminal and run this: 'go run Client.go'
14. Wait for the reply "Masukkan pesan: ". If there is, so enter your message.
15. If you have done, there is message like:
    "Pesan terkirim: "
    It means the messages have delivered
16. Go to the first terminal and check to ensure the messages saved to your database. If delivered successfully, there is feedback: "Pesan disimpan: "
17. Check your database table and if it's success, your message will appear there.

    THANK YOU
    KEEP CODING AND JUST TRY IT!

    If you have any problem, just catch to me on **darellaldo2004@gmail.com**
