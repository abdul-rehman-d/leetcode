package time_based_key_value_store

type Data struct {
	timestamps []int
	values     []string
	len        int
}

type TimeMap struct {
	data map[string]Data
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string]Data),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if data, has := this.data[key]; has {
		data.timestamps = append(data.timestamps, timestamp)
		data.values = append(data.values, value)
		data.len++
		this.data[key] = data
	} else {
		this.data[key] = Data{
			timestamps: []int{timestamp},
			values:     []string{value},
			len:        1,
		}
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	data, has := this.data[key]
	if !has {
		return ""
	}

	start := 0
	end := data.len - 1

	idx := -1
	for start <= end {
		middleIdx := start + ((end - start) / 2)
		middle := data.timestamps[middleIdx]

		if middle == timestamp {
			idx = middleIdx
			break
		} else if middle > timestamp {
			end = middleIdx - 1
		} else {
			idx = middleIdx // min so far
			start = middleIdx + 1
		}
	}

	if idx == -1 {
		return ""
	}
	return data.values[idx]
}
