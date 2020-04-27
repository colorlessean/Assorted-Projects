// connection handler function for the client socket
// will be used for recieving messages from the server
void * receive_from_server(void * socket_desc);

// function to set up the client socket
int initialize_client_socket();

// 
void generate_connection_thread(int client_socket_id);

// sends messages from the client to the server
void send_to_server(int socket_id);
