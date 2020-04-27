#include <sys/socket.h>
#include <netinet/ip.h>
#include <stdio.h>
#include <pthread.h>

int client_count = 0;

struct client {
    int index;
    int sock_id;
    struct sockaddr_in client_addr;
    int length;
};

// arrays to contain the clients and connection threads
struct client active_clients[16];
pthread_t connection_threads[16];

// function to open a socket for the server
// if sucessful will return the file descriptor for the server socket
// if failure will return -1
int setup_server_socket() {
    // PF_INET sets the protocol to ivp4
    // SOCK_STREAM allows bidirectional stream of bytes
    // i.e. opens a TCP socket according to ip(7) linux manual
    int soc = socket(PF_INET, SOCK_STREAM, 0);

    struct sockaddr_in server_addr;

    server_addr.sin_family = AF_INET;
    // htons coverts the address array into an intenet address
    // INADDR_ANY lets the system pick an appropriate address
    server_addr.sin_addr.s_addr = htons(INADDR_ANY);
    server_addr.sin_port = htons(8080);

    // bind will assign the address specified by the sockaddr to the file described by the socket
    int successfully_bound = bind(soc, (struct sockaddr *) &server_addr, sizeof(server_addr));

    if (successfully_bound == -1) {
        return -1;
    }

    // must listen on the actual open socket that has been bound
    // allow for 16 connections
    int listening = listen(soc, 16);

    if (listening == -1) {
        return -1;
    }

    printf("Server started listening on port 8080\n");

    return soc;
}

// sends a message to all the clients except the one sending it
void send_all_clients(int sender_sockfd, const void * msg, size_t length) {
    for (int i = 0; i < client_count; i++) {
        // TODO add thread safe access to the list of clients
        if (active_clients[i].sock_id != sender_sockfd){
            send(active_clients[i].sock_id, msg, length, 0);
        }
    }
}

// function to be attached to each thread
void * connection_handler(void * client_detail) {
    // get both of the values of the socketfd and index from the client_detail
    int client_socket = ((struct client *)client_detail)->sock_id;
    int index = ((struct client *)client_detail)->index;

    printf("Client %d connected on socket %d.\n",index + 1, client_socket);

    // super loop for each client handler thread to recieve and push messages
    while (1){
        // allow for 1024 char (byte) messages
        char data[1024];

        // calls the recieve function on the socket setup for the client
        // returns the number of bytes recieved
        int read = recv(client_socket, (void*)&data, 1024, 0);

        // null terminates the "string"
        data[read] = '\0';

        // push message to all clients (except the current client)
        send_all_clients(client_socket, (const void *)&data, 1024);
    }

    return NULL;
}

void generate_connection_thread(int socket_id) {
    // accepts a connection on the socket_id that has been bound and is listening
    // pass the accept function a pointer to the sockaddr structure to fill and a socklen_t structure to fill
    // hangs onto accept until a new client connection comes thru (i.e. blocks)
    // NOTE the id and index are for the socket set up by the client side
    active_clients[client_count].sock_id = accept(socket_id, (struct sockaddr *)&active_clients[client_count].client_addr, (socklen_t *) &active_clients[client_count].length);
    active_clients[client_count].index = client_count;

    printf("Socket: %d\n", active_clients[client_count].sock_id);
    printf("Index(client count): %d\n", active_clients[client_count].index);

    // creates a new POSIX (linux) thread
    // first argument is a location to store the thread id
    // second argument is the attributes for the thread NULL selects defaults
    // third argument is the function the thread will run on start
    // fourth argument is the argument passed to the function given as third argument
    pthread_create(&connection_threads[client_count], NULL, connection_handler, (void *) &active_clients[client_count]);

    client_count++;

}
