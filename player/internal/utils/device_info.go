package utils

import (
	"fmt"
	"net"
)

var Device *DeviceInfo

// DeviceInfo 设备信息
type DeviceInfo struct {
	Address string // 对应ip地址
	Mac     string // 对应mac地址
}

// GetDeviceMac 获取设备mac地址
// @return 		err		error  		错误信息
func (d *DeviceInfo) GetDeviceMac() (err error) {
	// 获取出口ip
	err = d.getOutboundIp()
	if err != nil {
		return
	}
	netIf, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, iFace := range netIf {
		// 跳过回环接口（127.0.0.1/::1 对应的 lo/loopback）
		if iFace.Flags&net.FlagLoopback != 0 {
			continue
		}
		// 跳过未启用的接口（必须是 UP 状态）
		if iFace.Flags&net.FlagUp == 0 {
			continue
		}
		addrs, _ := iFace.Addrs()
		for _, addr := range addrs {
			//fmt.Println(addr.String())
			if addr.(*net.IPNet).IP.String() == d.Address {
				d.Mac = iFace.HardwareAddr.String()
				return nil
			}
		}
	}
	return fmt.Errorf("没有找到出口地址")

}

// getOutboundIp 获取出口ip
// @return	string	ip出口
// @return 	error	错误信息
func (d *DeviceInfo) getOutboundIp() error {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP("8.8.8.8"),
		Port: 53,
	})
	if err != nil {
		return err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	d.Address = localAddr.IP.String()
	return nil
}
