// +build linux

package ioctl

// BlkPart send blkpart ioctl to fd
func BlkRRPart(fd uintptr) error {
	return IOCTL(fd, IO(0x12, 95), uintptr(0))
}

// BlkPg send blkpg ioctl to fd
func BlkPg(fd, data uintptr) error {
	return IOCTL(fd, IO(0x12, 105), data)
	return nil
}

// Fitrim send fitrim ioctl to fd
func Fitrim(fd, data uintptr) error {
	return IOCTL(fd, IOWR('X', 121, uintptr(0)), data)
}

// Firfreeze send firfreeze ioctl to fd
func Firfreeze(fd, data uintptr) error {
	return IOCTL(fd, IOWR('X', 119, uintptr(0)), data)
}
