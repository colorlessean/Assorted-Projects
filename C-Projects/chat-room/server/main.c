#include "server.h"

// handles the calling of the server functions
int main(int argc, char * argv[]) {
    // setup the server
    int server_socket = setup_server_socket();

    // error in server setup
    if (server_socket == -1) {
        return -1;
    }

    // super loop for accepting and handling the client connections
    while (1) {
        generate_connection_thread(server_socket);
    }

    return 0;
}
