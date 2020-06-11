# 网络编程之TCP接口方法整理

## net.Listen()

```go
func Listen(network, address string) (Listener, error) {
	var lc ListenConfig
	return lc.Listen(context.Background(), network, address)
}

```
### 参数：
- network: string类型，表示选用的协议，如TCP/UDP， "tcp", "udp"

- address: string类型，表示 IP地址+端口号，如"127.0.0.1:8080", ":8080"
### 返回值
- error: error类型，表示错误

- Listener: interface 类型，是面向流协议的通用网络侦听器。本质是一个 socket,用于监听的socket

Listener 定义如下：
```go
type Listener interface {
	Accept() (Conn, error)

	Close() error

	Addr() Addr
}
```

### Accept() (Conn, error) 
该方法表示：接受等待，并将下一个连接返回给侦听器

### 参数：
- error: error类型，表示错误

- Conn: interface 类型，表示面向流的通用网络连接。本质也是一个socket，用于和客户端进行通信
```go

type Conn interface {

	Read(b []byte) (n int, err error)

	Write(b []byte) (n int, err error)

	Close() error


	LocalAddr() Addr

	RemoteAddr() Addr

	SetDeadline(t time.Time) error

	SetReadDeadline(t time.Time) error

	SetWriteDeadline(t time.Time) error
}
```


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