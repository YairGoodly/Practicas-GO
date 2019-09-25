package main
import "fmt"
func main(){
   //Declaracion de variables.
   Cuo := 1500.0
   //Nom,ApP,ApM,Eda,Cuo,Acr := "Brandon","Cuatecatl","Cuapa",21,1500.00,false
   //Declaracion Matriz
   Meses := [] string{"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre"}

   //Ciclo de impresion
   //1ra forma
   i := 0
   for i <= 11 {
     Prc := Cuo * .05
     Cuo =  Cuo + Prc
     fmt.Println(Meses[i] ,",Nueva Cuota =")
     fmt.Printf("%6.2f\n",Cuo)
     i++
   }
   //2da forma
   i := 0
     for{
       if i > 11 {
         break
       }
       Prc := Cuo * .05
       Cuo =  Cuo + Prc
       fmt.Println(Meses[i] ,",Nueva Cuota =")
       fmt.Printf("%6.2f\n",Cuo)
       i++
     }
}
