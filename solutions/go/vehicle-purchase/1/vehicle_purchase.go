package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	//panic("NeedsLicense not implemented")
    if kind == "car" || kind == "truck" {
        return true
    } else {
        return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	//panic("ChooseVehicle not implemented")
    var ans string
    if option1 > option2 {
        ans = option2
    } else {
        ans = option1
    }
    return ans + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	//panic("CalculateResellPrice not implemented")
    var dc float64
    if age <3 {
        dc = 0.8
    } else if age >= 10 {
        dc = 0.5
    } else {
        dc = 0.7
    }
    return originalPrice * dc
}
