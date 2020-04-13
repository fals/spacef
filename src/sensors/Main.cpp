#include <stdio.h>
#include <string.h> //strlen
#include <sys/socket.h>
#include <arpa/inet.h> //inet_addr

int main(int argc, char *argv[])
{
	int socket_desc;
	struct sockaddr_in server;

	//Create socket
	socket_desc = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
	if (socket_desc == -1)
	{
		printf("Could not create socket\n");
		return -1;
	}
	else
	{
		printf("Socket Created\n");
	}
	
	memset(&server, 0, sizeof(server));

	server.sin_family = AF_INET;
	server.sin_addr.s_addr = inet_addr("172.16.238.10");
	server.sin_port = htons(15002);

	//Connect to remote server
	int conn = connect(socket_desc, (struct sockaddr *)&server, sizeof(server));
	if (conn < 0)
	{
		printf("Connection failed! Error %d\n", conn);
		return -1;
	}
	else
	{
		printf("Connect successfully done.\n");
	}

	puts("Connected\n");

	//Send some data
	char const *message = "GET / HTTP/1.1\r\n\r\n";
	if (send(socket_desc, message, strlen(message), 0) < 0)
	{
		puts("Send failed");
		return 1;
	}
	puts("Data Send\n");

	return 0;
}