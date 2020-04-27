#include <stdlib.h>
#include <stdio.h>
#include <sys/socket.h>
#include <netinet/ip.h>
#include <pthread.h>

void * connection_handler(void * sock_id) {
    int client_socket = *((int *)sock_id);

    while(1) {
        char data[1024];
        int read = recv(client_socket, data, 1024, 0);
        data[read] = '\0'; // null terminated string
        printf("%s\n", data);
    }

    return NULL;
}

void generate_connection_thread(int client_socket_id) {
    pthread_t receiving_thread;
    pthread_create(&receiving_thread, NULL, connection_handler, (void *) &client_socket_id);
}

int initialize_client_socket() {
    // PF_INET sets the protocol to ivp42
    // SOCK_STREAM allows bidirectional stream of bytes
    // i.e. opens a TCP socket according to ip(7) linux manual
    int client_socket = socket(PF_INET, SOCK_STREAM, 0);

    struct sockaddr_in server_addr;

    server_addr.sin_family = AF_INET;
    // htons coverts the address array into an intenet address
    // INADDR_ANY lets the system pick an appropriate address 
    server_addr.sin_addr.s_addr = htons(INADDR_ANY);
    server_addr.sin_port = htons(8080);

    int client_connection = connect(client_socket, (struct sockaddr *)&server_addr, sizeof(server_addr));

    if (client_connection == -1) {
        return -1;
    }

    printf("Connection with server established\n");

    return client_socket;
}

void send_to_server(int socket_id) {
    char input[1024];
    scanf("%s", input);

    send(socket_id, input, 1024, 0);
}
