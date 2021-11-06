from django.test import TestCase

# Create your tests here.
import socket
import base64
import hashlib

# 正常的socket代码
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
# 防止linux/mac报错
sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
sock.bind(('127.0.0.1', 8080))
sock.listen(5)

conn, address = sock.accept()
data = conn.recv(1024)  # 获取客户端发送的消息


def get_headers(data):
    """
    将请求头格式化成字典
    :param data:
    :return:
    """
    header_dict = {}
    data = str(data, encoding='utf-8')

    header, body = data.split('\r\n\r\n', 1)
    header_list = header.split('\r\n')
    for i in range(0, len(header_list)):
        if i == 0:
            if len(header_list[i].split(' ')) == 3:
                header_dict['method'], header_dict['url'], header_dict['protocol'] = header_list[i].split(' ')
        else:
            k, v = header_list[i].split(':', 1)
            header_dict[k] = v.strip()
    return header_dict


# 想将http协议的数据处理成字典的形式方便后续取值
header_dict = get_headers(data)  # 将一大堆请求头转换成字典数据  类似于wsgiref模块
client_random_string = header_dict['Sec-WebSocket-Key']  # 获取浏览器发送过来的随机字符串

# magic string拼接
magic_string = '258EAFA5-E914-47DA-95CA-C5AB0DC85B11'  # 全球共用的随机字符串 一个都不能写错
value = client_random_string + magic_string  # 拼接

# 算法加密
ac = base64.b64encode(hashlib.sha1(value.encode('utf-8')).digest())  # 加密处理

# 将处理好的结果再发送给客户端校验
tpl = "HTTP/1.1 101 Switching Protocols\r\n" \
      "Upgrade:websocket\r\n" \
      "Connection: Upgrade\r\n" \
      "Sec-WebSocket-Accept: %s\r\n" \
      "WebSocket-Location: ws://127.0.0.1:8080\r\n\r\n"
response_str = tpl % ac.decode('utf-8')  # 处理到响应头中

# 将随机字符串给浏览器返回回去
print(f"建立连接,加密验证key{ac}")
conn.send(bytes(response_str, encoding='utf-8'))


def get_data(info):
    """
    前后端进行通信,对前端发生消息进行解密
    """
    payload_len = info[1] & 127
    if payload_len == 126:
        extend_payload_len = info[2:4]
        mask = info[4:8]
        decoded = info[8:]
    elif payload_len == 127:
        extend_payload_len = info[2:10]
        mask = info[10:14]
        decoded = info[14:]
    else:
        extend_payload_len = None
        mask = info[2:6]
        decoded = info[6:]
    bytes_list = bytearray()
    for i in range(len(decoded)):
        chunk = decoded[i] ^ mask[i % 4]  # 异或运算
        bytes_list.append(chunk)

    body = str(bytes_list, encoding='utf-8')
    return body


# 基于websocket通信
while True:
    # ws.send("info")
    data = conn.recv(1024)  # 数据是加密处理的
    # print(data)
    # 对data进行解密操作
    value = get_data(data)
    print(value)

