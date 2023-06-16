** Libro Diario **

Es el libro del que dependen los demás libros de contabilidad.

Debe tener fecha de cada transacción, las cuentas involucradas, el importe de la operación y una breve explicación de la transacción. Se le llama apunte o asiento contable.
El conjunto de cuentas involucradas en cada transacción es variable, lo que implíca que a un nivel de base de datos, habrá que definir una unidad más pequeña... ¿subapunte?

Por otra parte, cada apunte tiene que sumar cero, por el principio de la partida doble, el cual se basa en que todo hecho económico tiene origen en otro hecho de igual valor, pero de naturaleza contraria.

Hay un libro diario por cada ejercicio fiscal, luego es un atributo que habrá que poner en cada apunte

Respecto a las cuentas, tienen un número, tienen una cuenta padre, excepto las cabezas de grupo y una descripción. En principio una cuenta no cambia en cada año, aunque si que cambia por algún cambio de ley, así que debería tener un atributo de inicio y otro de fin de vigencia. Sería interesante marcar su naturaleza: Activo, Pasivo, Patrimonio, Gasto, Ingreso, y si crece en el haber o en el debe

Creo que es interesante separar los años fiscales en otra tabla, de tal manera que si el periodo no es el año natural se pueda hacer una comprobación antes de añadirle un apunte, y luego por otra parte se puede cerrar, de tal manera que no se permita meter más apuntes a libro díario ya presentado.