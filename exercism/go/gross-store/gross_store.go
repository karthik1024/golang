package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	m := make(map[string]int)
	m["quarter_of_a_dozen"] = 3
	m["half_of_a_dozen"] = 6
	m["dozen"] = 12
	m["small_gross"] = 120
	m["gross"] = 144
	m["great_gross"] = 1728
	return m
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	u, e := units[unit]
	if !e {
		return false
	}

	_, e = bill[item]
	if !e {
		bill[item] = u
	} else {
		bill[item] += u
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	v, e := bill[item]
	if !e {
		return false
	}

	u, e := units[unit]
	if !e {
		return false
	}

	if v-u < 0 {
		return false
	}

	if v == u {
		delete(bill, item)
	} else {
		bill[item] -= u
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, e := bill[item]
	return v, e
}
