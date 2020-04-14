#include <string>
#include <chrono>
using namespace std;

class SensorData
{
private:
    string id;
    string type;
    string unit;
    float value;

public:
    SensorData(/* args */);
    ~SensorData();
    std::chrono::system_clock::time_point SensorData::getTimestamp();
};

SensorData::SensorData(/* args */)
{
}

SensorData::~SensorData()
{
}

std::chrono::system_clock::time_point SensorData::getTimestamp()
{
    using namespace std::chrono;
    return system_clock::time_point{milliseconds{1437520382241}};
}
