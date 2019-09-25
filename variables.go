package main
import "fmt"
func main(){
   //Declaracion de variables.
   Cuo := 1500.0
   //Nom,ApP,ApM,Eda,Cuo,Acr := "Brandon","Cuatecatl","Cuapa",21,1500.00,false
   //Declaracion Matriz
   Meses := [] string{"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre"}
   size := len(Meses)
   
  //For unicamente utlizado para dar formato al mostrar los datos
  i := 0
  for i < size{
    Prc := Cuo * .05
    Cuo =  Cuo + Prc
    fmt.Println(Meses[i:i+1], "Nueva Cuota")
    fmt.Printf("%6.2f\n",Cuo)
    i++
  }
}
