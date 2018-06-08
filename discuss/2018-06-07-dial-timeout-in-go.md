问题：
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error) {
func DialTimeout(network, address string, timeout time.Duration) (Conn, error) { 
这两个方法返回的对象类型不一样，想得到TCPConn对象，同时可以设置连接超时，应该怎么写？有谁可以帮帮忙吗？

回答：
Use net.Dialer with either the Timeout or Deadline fields set.

d := net.Dialer{Timeout: timeout}
conn, err := d.Dial("tcp", addr)
if err != nil {
   // handle error
}
A variation is to call Dialer.DialContext with a deadline or timeout applied to the context.

Type assert to *net.TCPConn if you specifically need that type instead of a net.Conn:

tcpConn, ok := conn.(*net.TCPConn)
