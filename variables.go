package main
import "fmt"
func main() {
  //Declaracion de Variables
  //1ra Manera de declarar Variables.
  var Nom  string
  var ApP  string
  var ApM  string
  //2da Forma de Declararlas
  Cuo := 1500.0
  Nom,ApP,ApM,Eda,Cuo,Acr := "Brandon","Cuatecatl","Cuapa",21,1500.00,false
  //Impresion de Variables
  fmt.Println("Resultado de Impresion")
  fmt.Println(Nom,ApP,ApM,Eda,Cuo,Acr)
  if Acr == false {
  fmt.Println("Sin Autorizacion");
}else{
  fmt.Println(Nom,ApP,ApM,Eda,Cuo,Acr)
}
}
