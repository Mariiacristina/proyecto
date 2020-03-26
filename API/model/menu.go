package model

import(
  "log"
  "Binfa/API/schema"
  "Binfa/API/connection"
)

var (
  err error
  errExec error
)

func UsuarioExistente(persona schema.Signin)(usuario schema.Signin,err error){
  var Usuario schema.Signin
  db := connection.Connect()
  err = db.QueryRow("SELECT id_cliente,nombre,password from cliente WHERE cliente.nombre = ?",persona.Nombre).Scan(&Usuario.Id_cliente,&Usuario.Nombre,&Usuario.Password)
  connection.Disconnect(db)
  if(err != nil) {
    log.Println("error en el modelo:", err)
    return Usuario,err
  }else{
    if (persona.Nombre == Usuario.Nombre && persona.Password == Usuario.Password){
      return Usuario,err
    }
    return Usuario,err
  }
}
