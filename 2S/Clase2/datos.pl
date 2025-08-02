persona(juan, 25, direccion(calle1, 123)).
persona(maria, 30, direccion(calle2, 456)).
persona(pedro, 18, direccion(calle1, 123)).
persona(luisa, 40, direccion(calle3, 789)).
trabaja_en(maria, universidad).
amigo(juan, maria).
amigo(maria, pedro).
amigo(pedro, juan).
edad_mayor(X, Y) :- persona(X, E1, _), persona(Y, E2, _), E1 > E2.
