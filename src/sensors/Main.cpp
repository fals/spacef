#include <stdio.h>
#include <string.h> //strlen
#include <sys/socket.h>
#include <arpa/inet.h> //inet_addr
#include <unistd.h>
#include <Sensor.h>

int main(int argc, char *argv[])
{
	Sensor wind = Sensor("wind1", "wind", "m_s");
	wind.start("172.16.238.10", 15002);
	wind.sendData();
	
	return 0;
}