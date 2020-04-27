#include "client.h"

int main(int argc, char * argv[]) {
    int client_socket_id = initialize_client_socket();

    // error in setting up the client socket
    if (client_socket_id == -1) {
        return -1;
    }

    generate_connection_thread(client_socket_id);

    // super loop for sending the messages to the server as the client inputs
    while(1) {
        send_to_server(client_socket_id);
    }

    return 0;
}
