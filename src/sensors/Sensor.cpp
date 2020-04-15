#include <stdio.h>
#include <string.h> //strlen
#include <sys/socket.h>
#include <arpa/inet.h> //inet_addr
#include <unistd.h>
#include <sstream>
#include <algorithm>
#include <chrono>
#include <stdlib.h> /* srand, rand */
#include <time.h>
#include <Sensor.h>

using namespace std;

Sensor::Sensor(string id, string type, string unit)
{
	this->id = id;
	this->type = type;
	this->unit = unit;
}

Sensor::~Sensor()
{
}

int Sensor::start(string address, int port)
{
	this->socket_addr = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
	if (this->socket_addr == -1)
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
	server.sin_addr.s_addr = inet_addr(address.c_str());
	server.sin_port = htons(port);

	//Connect to remote server
	int conn = connect(this->socket_addr, (struct sockaddr *)&server, sizeof(server));
	if (conn < 0)
	{
		printf("Connection failed! Error %d\n", conn);
		return -1;
	}
	else
	{
		printf("Connect successfully done.\n");
	}

	return 0;
}

int Sensor::sendData()
{
	using namespace std::chrono;

	while (1 == 1)
	{
		srand(time(NULL));
		int value = rand() % 100 + 1;
		uint64_t timestamp =  duration_cast<milliseconds>(system_clock::now().time_since_epoch()).count();

		//char const *message = "id=sensor_number&tp=temperature&un=C&vl=23.4&ts=1586896417"
		std::stringstream stream;
		stream 
			<< "id=" << id 
			<< "&tp=" << type 
			<< "&un=" << unit
			<< "&vl=" << value 
			<< "&ts=" << timestamp;

		string temp = stream.str();
		char const *message = temp.c_str();
		printf(message);

		if (send(socket_addr, message, strlen(message), 0) < 0)
		{
			printf("Send failed\n");
			return 1;
		}
		printf("Data Send\n");
		sleep(1);
	}
}
