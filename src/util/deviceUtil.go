package util

func CreateInitialDeviceState(deviceIp string) string {
	if deviceIp == "" {
		return "Pending"
	} else {
		return "Onboarding"
	}
}
