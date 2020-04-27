# Compilation Details
In order to compile the server files need to run the following
```
gcc -pthread -o server.exe main.c server.c
```
This is due to the fact that the gcc complier on linux needs to know about the library of pthread