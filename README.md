# BundleInspect CLI

![image](https://github.com/luis-tenorio-code/BundleInspect/assets/101147375/90d5115f-152f-468a-8a99-05e755260f5d)

## Descripción

BundleInspect es una herramienta de línea de comandos (CLI) que te permite obtener información detallada sobre el tamaño de los paquetes utilizando la API de BundlePhobia.

## Opciones de Instalación

### 1. Instalación desde el Código Fuente

Para instalar BundleInspect desde el código fuente, sigue estos pasos:

1.1. Clona el repositorio:

    git clone https://github.com/luis-tenorio-code/BundleInspect.git

1.2. Navega al directorio del proyecto:

    cd BundleInspect

1.3. Ejecuta el comando de construcción:

    go build -o BundleInspect main.go

1.4. Asegúrate de tener permisos de ejecución:

    chmod +x BundleInspect

1.5. Ahora puedes ejecutar la aplicación:

    ./BundleInspect <nombre del paquete>

### 2. Descarga del Binario Compilado

Puedes descargar la versión compilada de BundleInspect desde el siguiente enlace: [Descargar BundleInspect](https://github.com/luis-tenorio-code/BundleInspect/raw/main/infoBuild)

2.1. Después de la descarga, puedes ejecutar la aplicación directamente desde la línea de comandos:
  **Nota:** Asegúrate de tener permisos de ejecución para el archivo descargado:

    chmod +x BundleInspect
    
Reemplaza `<nombre del paquete>` con el nombre del paquete que deseas analizar.

    ./BundleInspect <nombre del paquete>

## Licencia

Este proyecto está bajo la licencia [MIT](LICENSE).
