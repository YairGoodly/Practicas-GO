package main
import "fmt"
func main(){
   //Declaracion de variables.
   //Cuo := 1500.0
   //Nom,ApP,ApM,Eda,Cuo,Acr := "Brandon","Cuatecatl","Cuapa",21,1500.00,false
   //Declaracion Matriz
   Meses := [] string{"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre"}

   //Ciclo de impresion
   //3ra forma
   for i, v := range Meses{
     fmt.Println("Mes", i , v)
   }
   //4ta forma
   //variable donde se guarda el tamaño de la Matriz
   size := len(Meses)
   i := 0
   for i < size {
     fmt.Println(Meses[i])
     i++
     }
}
