# 网络编程之UDP服务端接口

## net.ResolveUDPAddr()
创建一个 udp 地址结构。  指定服务器ip地址 端口号
```go
func ResolveUDPAddr(network, address string) (*UDPAddr, error) {
	switch network {
	case "udp", "udp4", "udp6":
	case "": // a hint wildcard for Go 1.0 undocumented behavior
		network = "udp"
	default:
		return nil, UnknownNetworkError(network)
	}
	addrs, err := DefaultResolver.internetAddrList(context.Background(), network, address)
	if err != nil {
		return nil, err
	}
	return addrs.forResolve(network, address).(*UDPAddr), nil
}

```
### 参数：
- network: string类型，表示选用的协议， "udp", "udp4", "udp6"

- address: string类型，表示 IP地址+端口号，如"127.0.0.1:8080", ":8080"

### 返回值
- error: error类型，表示错误

- *UDPAddr: 结构体指针类型，表示UDP端点的地址，封装了服务端的地址+端口

UDPAddr 定义如下：
```go
// UDPAddr represents the address of a UDP end point.
type UDPAddr struct {
	IP   IP
	Port int
	Zone string // IPv6 scoped addressing zone
}
```

## net.ListenUDP()
创建用户通信的socket

```go
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error) {
	switch network {
	case "udp", "udp4", "udp6":
	default:
		return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: UnknownNetworkError(network)}
	}
	if laddr == nil {
		laddr = &UDPAddr{}
	}
	sl := &sysListener{network: network, address: laddr.String()}
	c, err := sl.listenUDP(context.Background(), laddr)
	if err != nil {
		return nil, &OpError{Op: "listen", Net: network, Source: nil, Addr: laddr.opAddr(), Err: err}
	}
	return c, nil
}
```

### 参数：
- network: string类型，表示选用的协议， "udp", "udp4", "udp6"

- laddr: 结构体指针类型，表示UDP端点的地址，封装了服务端的地址+端口

### 返回值
- error: error类型，表示错误

- *UDPConn: 结构体指针类型，是用于UDP网络连接的Conn和PacketConn接口的实现

### udpConn.ReadFromUDP()
读取客户端数据

函数定义：
```go
// ReadFromUDP acts like ReadFrom but returns a UDPAddr.
func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error) {
	if !c.ok() {
		return 0, nil, syscall.EINVAL
	}
	n, addr, err := c.readFrom(b)
	if err != nil {
		err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
	}
	return n, addr, err
}
```

#### 参数：
- b: 字节切片类型，读取数据的缓冲区

#### 返回值
- int: int类型，表示读取到字节数

- *UDPAddr: 结构体指针类型，表示客户端 udp 地址结构

- error: error类型，表示错误


### udpConn.WriteToUDP()
向客户端写入数据
```go
// WriteToUDP acts like WriteTo but takes a UDPAddr.
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error) {
	if !c.ok() {
		return 0, syscall.EINVAL
	}
	n, err := c.writeTo(b, addr)
	if err != nil {
		err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
	}
	return n, err
}
```


#### 参数：
- b: 字节切片类型，读取数据的缓冲区

- *UDPAddr: 结构体指针类型，表示客户端 udp 地址结构
#### 返回值
- int: int类型，表示写入成功的字节数

- error: error类型，表示错误









# 客户端接口

## net.Dial()
```go
// Dial connects to the address on the named network.
//
// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
// "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4"
// (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and
// "unixpacket".
//
// For TCP and UDP networks, the address has the form "host:port".
// The host must be a literal IP address, or a host name that can be
// resolved to IP addresses.
// The port must be a literal port number or a service name.
// If the host is a literal IPv6 address it must be enclosed in square
// brackets, as in "[2001:db8::1]:80" or "[fe80::1%zone]:80".
// The zone specifies the scope of the literal IPv6 address as defined
// in RFC 4007.
// The functions JoinHostPort and SplitHostPort manipulate a pair of
// host and port in this form.
// When using TCP, and the host resolves to multiple IP addresses,
// Dial will try each IP address in order until one succeeds.
//
// Examples:
//	Dial("tcp", "golang.org:http")
//	Dial("tcp", "192.0.2.1:http")
//	Dial("tcp", "198.51.100.1:80")
//	Dial("udp", "[2001:db8::1]:domain")
//	Dial("udp", "[fe80::1%lo0]:53")
//	Dial("tcp", ":80")
//
// For IP networks, the network must be "ip", "ip4" or "ip6" followed
// by a colon and a literal protocol number or a protocol name, and
// the address has the form "host". The host must be a literal IP
// address or a literal IPv6 address with zone.
// It depends on each operating system how the operating system
// behaves with a non-well known protocol number such as "0" or "255".
//
// Examples:
//	Dial("ip4:1", "192.0.2.1")
//	Dial("ip6:ipv6-icmp", "2001:db8::1")
//	Dial("ip6:58", "fe80::1%lo0")
//
// For TCP, UDP and IP networks, if the host is empty or a literal
// unspecified IP address, as in ":80", "0.0.0.0:80" or "[::]:80" for
// TCP and UDP, "", "0.0.0.0" or "::" for IP, the local system is
// assumed.
//
// For Unix networks, the address must be a file system path.
func Dial(network, address string) (Conn, error) {
	var d Dialer
	return d.Dial(network, address)
}
```
### 参数：
- network: string类型，表示选用的协议，如TCP/UDP， "tcp", "udp"

- address: string类型，表示 IP地址+端口号，如"127.0.0.1:8080", ":8080"

### 返回值：
- conn:  同 net.Accept() 的参数
- error: 同 net.Accept() 的参数