package main

import "fmt"

/*
Состояние — это поведенческий паттерн, позволяющий динамически изменять поведение объекта при смене его состояния.
Преимущества и недостатки:
+ Избавляет от множества больших условных операторов машины состояний.
+ Упрощает код контекста.
+ Концентрирует в одном месте код, связанный с определённым состоянием.
- Может неоправданно усложнить код, если состояний мало и они редко меняются.

Пример ТCP соединение
оно может быть в трех состояних:
открыто(opened), прослушивание(listening), закрыто(closed)
*/

// TCPState - интерфейс различных состояний
type TCPState interface {
	Open()
	Close()
	Listen()
}

type TCPConnection struct {
	opened    TCPState
	closed    TCPState
	listening TCPState
	curState  TCPState
}

func New() *TCPConnection {
	conn := &TCPConnection{}
	tcpOpened := &TCPOpened{
		conn: conn,
	}
	tcpClosed := &TCPClosed{
		conn: conn,
	}
	tcpListening := &TCPListening{
		conn: conn,
	}
	conn.opened = tcpOpened
	conn.listening = tcpListening
	conn.closed = tcpClosed
	conn.SetState(conn.closed)
	return conn
}

func (tConn *TCPConnection) Open() {
	tConn.curState.Open()
}

func (tConn *TCPConnection) Close() {
	tConn.curState.Close()
}

func (tConn *TCPConnection) Listen() {
	tConn.curState.Listen()
}

func (tConn *TCPConnection) SetState(state TCPState) {
	tConn.curState = state
}

type TCPOpened struct {
	conn *TCPConnection
}

func (t *TCPOpened) Open() {
	fmt.Println("Cannot open already opened connection")
}

func (t *TCPOpened) Close() {
	fmt.Println("Transition from open to closed")
	t.conn.SetState(t.conn.closed)
}

func (t *TCPOpened) Listen() {
	fmt.Println("Transition from open to listening")
	t.conn.SetState(t.conn.listening)
}

type TCPClosed struct {
	conn *TCPConnection
}

func (t *TCPClosed) Open() {
	fmt.Println("Transition from close to opened")
	t.conn.SetState(t.conn.opened)
}

func (t *TCPClosed) Close() {
	fmt.Println("Cannot close already closed connection")
}

func (t *TCPClosed) Listen() {
	fmt.Println("Cannot listening on this connection because connection is closed")
}

type TCPListening struct {
	conn *TCPConnection
}

func (t *TCPListening) Open() {
	fmt.Println("Cannot open connection because connection is listening")
}

func (t *TCPListening) Close() {
	fmt.Println("Transaction from listening state to closed state")
	t.conn.SetState(t.conn.closed)
}

func (t *TCPListening) Listen() {
	fmt.Println("Cannot listening on this connection because connection is already listening")
}

func main() {
	conn := New()
	// Моделируем теперь работу соединение
	conn.Open()
	conn.Listen()
	conn.Close()

	// Пробуем сделать различные аварийные состояния

	conn.Close()
	conn.Listen()

	conn.Open()
	conn.Listen()
	conn.Open()
}
