// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	//panic("Please implement NewResident.")
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	//panic("Please implement HasRequiredInfo.")
	if r.Name == "" {
		return false
	}
	if len(r.Address) == 0 {
		return false
	}
	if addr, ok := r.Address["street"]; !ok {
		return false
	} else {
		if addr == "" {
			return false
		}
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	//panic("Please implement Delete.")
	r.Name = ""
	r.Address = nil
	r.Age = 0
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	//panic("Please implement Count.")
	sum := 0
	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			sum += 1
		}
	}
	return sum
}
