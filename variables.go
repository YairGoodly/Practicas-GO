package main
import "fmt"
func main(){
   //Declaracion de variables.
   Cuo := 1500.0
   //Nom,ApP,ApM,Eda,Cuo,Acr := "Brandon","Cuatecatl","Cuapa",21,1500.00,false
   //Declaracion Matriz
   Meses := [] string{"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre"}
   //Ciclo de impresion
   for i :=0; i <= 11; i++{
     Prc := Cuo * .05
     Cuo =  Cuo + Prc
     fmt.Println(Meses[i] ,"Nueva Cuota =")
     //Printf utlizado para dar formato a valores con decimales
     fmt.Printf("%6.2f\n",Cuo)
     }
}
