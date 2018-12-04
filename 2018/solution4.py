import sys
import datetime
from common import read_strings


def guard_minute_optimizer_naive(guard_logs):
    guard_map = parse_guard_logs(guard_logs)
    def most_minutes_asleep(guard_map):
        max_guard_id = 0
        for guard_id in guard_map:
            if (get_total_hours_slept(guard_id, guard_map)
                    > get_total_hours_slept(max_guard_id, guard_map)):
                max_guard_id = guard_id
        return max_guard_id

    best_guard = most_minutes_asleep(guard_map)
    return best_guard, get_most_slept_minute(best_guard, guard_map)

def parse_guard_logs(guard_logs):
    shifts = []
    current_log = []
    for log in guard_logs:
        if 'Guard' in log:
            if current_log != []:
                shifts.append(Shift(current_log))
                current_log = []
        current_log.append(log)
    shifts.append(Shift(current_log))
    guard_map = {x: [] for x in set(x.guard_id for x in shifts)}
    for shift in shifts:
        guard_map[shift.guard_id].append(shift.sleep_cycle)
    return guard_map


def get_total_hours_slept(guard_id, guard_map):
    if guard_id not in guard_map:
        return 0
    return sum(len([x for x in y if not y[x]]) for y in guard_map[guard_id])

def get_most_slept_minute(guard_id, guard_map):
    times = [0] * 60
    for cycle in guard_map[guard_id]:
        for time in cycle:
            if not cycle[time]:
                times[time] += 1
    return times.index(max(times))


class Shift:
    def __init__(self, shift_array):
        self.guard_id = _parse_guard_id(shift_array[0])
        self.sleep_cycle = {}
        self._parse_sleep_cycle(shift_array)

    def _parse_sleep_cycle(self, shifts):
        state = {'awake': True, 'time': _parse_time(shifts[0])}
        for shift in shifts[1:]:
            for time in _time_diff(state['time'], _parse_time(shift)):
                self.sleep_cycle[time] = state['awake']
            state = {'awake': not state['awake'], 'time': _parse_time(shift)}
        return shifts

def _parse_guard_id(state_string):
    return next(int(item[1:]) for item in state_string.split(' ') if '#' in item)


def _parse_time(shift):
    values = [x for x in shift.split(' ')]
    return values[1][:-1]

def _time_diff(a, b):
    cursor_hour, cursor_min = [int(x) for x in a.split(':')]
    end_hour, end_min = [int(x) for x in b.split(':')]
    while (cursor_hour, cursor_min) != (end_hour, end_min):
        if cursor_hour == 0:
            yield cursor_min
        else:
            break
        cursor_hour, cursor_min = _increment(cursor_hour, cursor_min)

def _increment(hour, minute):
    if minute == 59:
        return (hour+1, 0)
    return (hour, minute+1)

def _extract_time(time):
    return time.split(' ')[0][1:] + ' ' + time.split(' ')[1][:-1]


def guard_frequency_optimizer(guard_logs):
    guard_map = parse_guard_logs(guard_logs)
    def get_most_frequent_minute(guard):
        hours = {time: 0 for time in range(0, 60)}
        for shift in guard:
            for hour in shift:
                if not shift[hour]:
                    hours[hour] += 1
        print(hours)
        max_minute_freq = 0
        for hour in hours:
            if hours[hour] > max_minute_freq:
                max_minute_freq = hours[hour]
        return next((hour, hours[hour]) for hour in hours if hours[hour] == max_minute_freq)
    most_frequent_minutes = {
        guard_id: get_most_frequent_minute(guard_map[guard_id])
        for guard_id in guard_map
    }
    max_f = (0, 0)
    max_z = 0
    for z in most_frequent_minutes:
        if most_frequent_minutes[z][1] > max_f[1]:
            max_f = most_frequent_minutes[z]
            max_z = z
    return max_z, max_f[0]


if __name__ == '__main__':
    GUARD_LOGS = read_strings(sys.argv[1])
    SORTED_GUARD_LOGS = sorted(
        GUARD_LOGS,
        key=lambda x: datetime.datetime.strptime(_extract_time(x), '%Y-%m-%d %M:%S'))
    x, y = guard_minute_optimizer_naive(SORTED_GUARD_LOGS)
    print(x * y)
    x, y = guard_frequency_optimizer(SORTED_GUARD_LOGS)
    print(x * y)
