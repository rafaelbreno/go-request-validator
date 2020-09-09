package validation

import "fmt"

func Required(value string, field string) (bool, string) {
	resp := true
	msg := ""
	if(len(value) <= 0) {
		resp = false
		msg = fmt.Sprintf("The field %s is required!", field) 
	}
	return resp, msg
}

func Max(value interface{}, field string, maxValue int) (bool, string) {
	switch value.(type) {
        case string:
    		valStr, _ := value.(string)
            return maxString(valStr, field, maxValue)
        default:
    		valInt, _ := value.(int)
            return maxInt(valInt, field, maxValue)
	}
}

func maxInt(value int, field string, maxValue int) (bool, string) {
	resp := true
	msg := ""
	if(value > maxValue) {
		resp = false
		msg = fmt.Sprintf("The field %s must be less or equal to %d!", field, maxValue)
	}
	return resp, msg
}

func maxString(value string, field string, maxValue int) (bool, string) {
	resp := true
	msg := ""
	if(len(value) > maxValue) {
		resp = false
		msg = fmt.Sprintf("The field %s must be less or equal to %d!", field, maxValue)
	}
	return resp, msg
}


func Min(value interface{}, field string, minValue int) (bool, string) {
	switch value.(type) {
        case string:
    		valStr, _ := value.(string)
			return minString(valStr, field, minValue)
        default:
    		valInt, _ := value.(int)
            return minInt(valInt, field, minValue)
	}
}

func minInt(value int, field string, minValue int) (bool, string) {
	resp := true
	msg := ""
	if(value < minValue) {
		resp = false
		msg = fmt.Sprintf("The field %s must be greater or equal to %d!", field, minValue)
	}
	return resp, msg
}

func minString(value string, field string, minValue int) (bool, string) {
	resp := true
	msg := ""
	if(len(value) < minValue) {
		resp = false
		msg = fmt.Sprintf("The field %s must be greater or equal to %d!", field, minValue)
	}
	return resp, msg
}