
% hechos: personas con edad y dirección (estructura compuesta)
persona(ana, 28, direccion(av_central, 10)).
persona(carlos, 55, direccion(av_norte, 45)).
persona(maria, 6, direccion(av_central, 10)).
persona(sofia, 34, direccion(av_sur, 7)).
persona(juan, 18, direccion(av_oeste, 2)).

% relaciones simples
padre(carlos, ana).
padre(carlos, juan).
madre(sofia, maria).
amigo(ana, sofia).
amigo(juan, ana).

% regla: abuelo (dos relaciones padre)
% PSEUDOCODIGO: Para `abuelo(A,B)` comprobar:
%  1) existe Z tal que padre(A,Z) es true
%  2) y existe padre(Z,B) es true
% Si ambas son true, la consulta unifica y devuelve A y B (o confirma la
% relación para A y B dados).
abuelo(X, Y) :- padre(X, Z), padre(Z, Y).

% regla: mayor que (comparacion aritmetica)
% PSEUDOCODIGO: Para `edad_mayor(A,B)`:
%  1) buscar hechos persona(A,E1,_) y persona(B,E2,_).
%  2) evaluar la comparación aritmética E1 > E2.
%  3) si verdadero, la consulta succeed; si hay variables, devuelven
%     las combinaciones que cumplan la condición.
edad_mayor(X, Y) :- persona(X, E1, _), persona(Y, E2, _), E1 > E2.

% -------------------------
% Predicados requeridos
% -------------------------

% 1) maximo(X, Y, M) usando corte verde (elige el mayor y evita alternativas)
maximo(X, Y, X) :- X >= Y, !.
maximo(_, Y, Y).

% PSEUDOCODIGO: maximo(A,B,M)
%  - Si A >= B entonces M = A y aplicar '!' para evitar alternativas.
%  - En caso contrario, el segundo caso unifica M = B.
% Ejemplo: maximo(2,4,M) => prueba primera cláusula (2>=4 false), va al
% segundo caso y M=4.

% 2) clasifica(X, Resultado) usando corte (ejemplo didáctico)
%    Mostramos dos versiones: una con corte "verde" (clasifica_g)
%    y otra donde el uso del corte puede alterar el significado (clasifica_r: corte "rojo" en clase)

% versión segura (verde): clasifica según edad (niño/adolescente/adulto)
clasifica_g(X, niño) :- X < 13, !.
clasifica_g(X, adolescente) :- X >= 13, X < 18, !.
clasifica_g(_, adulto).

% PSEUDOCODIGO: clasifica_g(Edad, Categoria)
%  - Si Edad < 13 => Categoria = niño, y '!' evita seguir buscando.
%  - Si 13 <= Edad < 18 => Categoria = adolescente, y '!' evita alternativas.
%  - En cualquier otro caso => Categoria = adulto.
% Nota: Uso del corte aquí es "verde" porque las cláusulas cubren casos
%       mutuamente excluyentes y se evita backtracking innecesario.

% versión con corte que puede comportarse como "rojo" si se usa sin cuidado
% (si la variable no está instanciada o las condiciones no cubren todos los casos,
%  el corte puede impedir alternativas deseadas)
clasifica_r(X, niño) :- X < 13, !.
clasifica_r(X, adolescente) :- X >= 13, X < 18, !.
clasifica_r(X, adulto) :- X >= 18, !.
clasifica_r(_, desconocido).

% PSEUDOCODIGO: clasifica_r(Edad, Resultado)
%  - Similar a clasifica_g pero con un corte final que puede ocultar
%    soluciones si la variable Edad no está instanciada o las cláusulas
%    no son exhaustivas; usar con precaución.

% 3) suma_lista(Lista, Suma) recursiva directa
suma_lista([], 0).
suma_lista([H|T], S) :- suma_lista(T, S1), S is H + S1.

% versión optimizada con recursión de cola
suma_lista_tr(List, Sum) :- suma_lista_tr(List, 0, Sum).
suma_lista_tr([], Acc, Acc).
suma_lista_tr([H|T], Acc, Sum) :- Acc1 is Acc + H, suma_lista_tr(T, Acc1, Sum).

% PSEUDOCODIGO: suma_lista(Lista, Suma)
%  - Caso base: lista vacía => suma 0.
%  - Caso recursivo: sumar cabeza H a la suma de la cola.
% PSEUDOCODIGO (tail-recursive): suma_lista_tr(List, Sum)
%  - Mantener un acumulador Acc que guarda la suma parcial.
%  - Al final, Acc contiene la suma total y se unifica con Sum.

% 4) Uso de \+ para verificar no pertenencia
not_member(X, L) :- \+ member(X, L).

% PSEUDOCODIGO: not_member(Elem, Lista)
%  - Devuelve true si no existe ninguna unificación de member(Elem, Lista).
%  - Importante: \+ (negación como falla) depende de que la subconsulta
%    member/2 falle; con variables libres puede comportarse de forma no
%    intuitiva.

% 5) Ejemplo que muestra cómo la negación como falla puede dar resultados inesperados
%    Si consultamos \+ persona(X, _, _) con X NO instanciada, Prolog intenta satisfacer
%    persona(X,_,_) y si hay al menos un hecho persona/3 la negación falla (es decir devuelve false).
%    En cambio, consultar \+ persona(juan,_,_) devolverá true/false según exista el hecho instanciado.
%    Aquí dejamos un predicado de ejemplo para ejecutar en pruebas.
negacion_ejemplo_uninstanciada(Result) :- ( \+ persona(_,_,_) -> Result = 'negacion_true' ; Result = 'negacion_false' ).

% PSEUDOCODIGO: negacion_ejemplo_uninstanciada(Result)
%  - Evalúa \+ persona(_,_,_): si no existe NINGUN hecho persona/3 entonces
%    la negación succeed y Result = 'negacion_true'. Si hay al menos un
%    hecho persona/3 la negación falla y Result = 'negacion_false'.
%  - Este predicado demuestra por qué no se debe usar \+ sobre metas con
%    variables totalmente no instanciadas si buscamos propiedades universales.


