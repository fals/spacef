#include <iostream>
#include <thread>
#include <vector>
#include <string>
#include <Sensor.h>

using namespace std;

void sendData(Sensor sensor, int interval)
{
	sensor.sendData(interval);
}

int main(int argc, char *argv[])
{
	std::vector<std::thread> sensors;
	for (int i = 0; i < 3; ++i)
	{
		string name = "fuel_" + std::to_string(i);
		Sensor fuel = Sensor(name, "pressure", "Pa");
		fuel.start("172.16.238.10", 15002);
		sensors.emplace_back(sendData, fuel, 1);
	}

	for (int i = 0; i < 3; ++i)
	{
		string name = "external_" + std::to_string(i);
		Sensor external = Sensor(name, "temp", "C");
		external.start("172.16.238.10", 15002);
		sensors.emplace_back(sendData, external, 2);
	}

	for (int i = 0; i < 3; ++i)
	{
		string name = "wind_" + std::to_string(i);
		Sensor wind = Sensor(name, "wind", "m_s");
		wind.start("172.16.238.10", 15002);
		sensors.emplace_back(sendData, wind, 5);
	}

	for (auto &t : sensors)
	{
		t.join();
	}

	return 0;
}