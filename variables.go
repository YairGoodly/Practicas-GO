package main
import "fmt"
func main(){
   //Declaracion de variables.
   Cuo := 1500.0
   //Nom,ApP,ApM,Eda,Cuo,Acr := "Brandon","Cuatecatl","Cuapa",21,1500.00,false
   //Declaracion Matriz
  // Meses := [] string{"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre"}
   //size := len(Meses)

   //Declaracion de Slice dinamico
    var Mont[]string
  //Declarando los valores que contendra el Slice
   Mont = append(Mont,"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre")

   //Mostrando los valores dentro del slice obteniedo su longitud o capacidad.
   //Mediante la funcion "len"
   i  := 0
   for i < len(Mont){
     Prc := Cuo * .05
     Cuo =  Cuo + Prc
     fmt.Println(Mont[i:i+1], "Nueva Cuota")
     fmt.Printf("%6.2f\n",Cuo)
     i++
   }

}
