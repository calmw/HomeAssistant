import socket

from haservice.state import get_state

sk = socket.socket()

sk.connect(("我的服务器地址", 8888))


while True:
    accept_data = sk.recv(1024)
    if accept_data == "ping":
        sk.sendall(bytes('pong', encoding="utf8"))
    elif accept_data == "quite":
        sk.close()
        break
    elif accept_data == "XXX":
        get_state()
    elif accept_data == "XXX":
        pass
