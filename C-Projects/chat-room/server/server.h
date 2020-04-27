#include <stdlib.h>

// Structure to hold client information
struct client;

// sends a message to all the clients except the one sending it
int send_all_clients(int sender_sockfd, const void * msg, size_t length);

// handles the generation of each of the threads for each of the connecting clients
void generate_connection_thread(int socket_id);

// thread attached function to handle the connection of new client
void * connection_handler();

// Setup function
int setup_server_socket();
