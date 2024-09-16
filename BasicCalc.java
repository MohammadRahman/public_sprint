package sprint;

public class BasicCalc {
    public int doOperation(int a, char op, int b) {
      switch (op) {
        case '+':
            return a + b;
        case '-': 
            return a - b;
        case '*':
            if(a == 0 || b ==0){
                return 0;
            }else{
                return a * b;
            }
        case '/':
            if(a == 0 || b ==0){
                return 0;
            }else{
                return a / b;
            }
        case '%':
            if(a == 0 || b ==0){
                return 0;
            }else{
                return a % b;
            }
        default:
           return 0;
      }    
    }

}
