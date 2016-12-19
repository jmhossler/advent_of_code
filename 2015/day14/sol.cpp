#include <iostream>
#include <string>
#include <regex>
#include <map>
#include <algorithm>
#include <list>

class reindeer {
    private:
        enum state {RUN, REST};
        int speed;
        int runtime;
        int resttime;
        state doing;
        int doingtime;

    public:
    int distance;
    int points;
    explicit reindeer() {}
    reindeer(int _s, int _run, int _rest) : speed(_s), runtime(_run), resttime(_rest),
        distance(0), doing(RUN), doingtime(0), points(0) {}
    void advance(void);
};

void reindeer::advance(void) {
    doingtime += 1;
    if (doing == RUN)
        distance += speed;
    if (doing == RUN && doingtime == runtime) {
        doing = REST;
        doingtime = 0;
    } else if (doing == REST && doingtime == resttime) {
        doing = RUN;
        doingtime = 0;
    }
}

int main(void) {
    std::string line;
    std::regex speedre{R"((\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.)"};
    std::map<std::string, reindeer> deer;

    while (std::getline(std::cin, line)) {
        std::smatch fields;
        if (std::regex_match(line, fields, speedre)) {
            deer.emplace(fields[1], reindeer{std::stoi(fields[2]), std::stoi(fields[3]), std::stoi(fields[4])});
        } else {
            std::cerr << "Unknown line  << line << \n";
        }
    }
    for (int s = 0; s < 2503; s++) {
        for (auto &d : deer)
            d.second.advance();
        std::list<std::string> leaders;
        int leading_distance = 0;
        for (auto &d : deer) {
            if (d.second.distance > leading_distance) {
                leaders.clear();
                leaders.push_back(d.first);
                leading_distance = d.second.distance;
            } else if (d.second.distance == leading_distance)  {
                leaders.push_back(d.first);
            }
        }
        for (auto &name : leaders)
            deer[name].points += 1;
    }

    int max_distance = 0;
    int max_points = 0;
    for (auto &d : deer) {
        max_distance = std::max(max_distance, d.second.distance);
        max_points = std::max(max_points, d.second.points);
    }
    std::cout << "Max distance: " << max_distance << " km\n";
    std::cout << "Max points: " << max_points << '\n';
    return 0;
}
