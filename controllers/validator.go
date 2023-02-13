package controllers

func validate(val any) (any, bool) {
	switch vType := val.(type) {

	case string:
		if val.(string) == "" {
			return vType, false
		}
	case int:
		if val.(int) == 0 {
			return vType, false
		}
	case float64:
		if val.(float64) == float64(0) {
			return vType, false
		}
	default:
		return val, true
	}

	return val, true
}
