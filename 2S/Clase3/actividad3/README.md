Actividad 3 - Datos y consultas Prolog

Contenido:
- datos.pl : base de hechos y reglas
- consultas.txt : consultas de ejemplo con explicación breve

Instrucciones rápidas:
1. Abrir `datos.pl` en el entorno (Tau Prolog en la web o la app Fyne si la tienes).
2. Ejecutar las consultas listadas en `consultas.txt`.
3. Tomar capturas de las consultas exitosas y escribir una breve explicación (2 líneas) por consulta.

Ejemplos de consultas útiles:
- persona(X, 28, D).
- edad_mayor(carlos, ana).
- persona(Nombre, Edad, direccion(Calle, Numero)).
- amigo(X, Y). (usar ; para seguir con más resultados)
- abuelo(carlos, Y).

Checklist de la consigna (estado):

- Crear `datos.pl` con hechos (>=5): DONE
- Implementar `maximo/3` con corte verde: DONE
- Implementar `clasifica/2` (ejemplos `clasifica_g` y `clasifica_r`): DONE
- Implementar `suma_lista/2` y versión en cola: DONE
- Usar \+ para verificar no pertenencia (`not_member`): DONE
- Mostrar ejemplo donde la negación con variables no instanciadas genera comportamiento inesperado: DONE

Cómo ejecutar en Tau-Prolog (web):
1. Abrir la página de Tau-Prolog.
2. Pegar el contenido de `datos.pl` en el área de programa y presionar "Cargar".
3. Escribir una consulta (por ejemplo `maximo(5,3,M).`) y presionar ejecutar.

Cómo ejecutar con la app Fyne incluida en este repositorio:
1. Desde la carpeta `2S/Clase2` ejecutar `go run main.go`.
2. Cargar el archivo `2S/Clase3/actividad3/datos.pl` usando el botón "Cargar archivo .pl".
3. Escribir las consultas en el campo de texto y presionar "Consultar".

Notas finales:
- Adjunta capturas de pantalla de las consultas exitosas y una breve explicación (2 líneas) en tu PDF de entrega.
