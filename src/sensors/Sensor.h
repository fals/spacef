#include <stdio.h>
#include <string.h> //strlen
#include <sys/socket.h>
#include <arpa/inet.h> //inet_addr
#include <unistd.h>
#include <sstream>
#include <chrono>
#include <stdlib.h>     /* srand, rand */
#include <time.h>

using namespace std;

class Sensor
{
private:
    string id;
    string type;
    string unit;
    int socket_addr;
    struct sockaddr_in server;

public:
    Sensor(string id, string type, string unit);
    ~Sensor();
    int start(string address, int port);
    int sendData(int interval);
};