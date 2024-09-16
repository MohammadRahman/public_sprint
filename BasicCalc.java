package sprint;

public class BasicCalc {
    public int doOperation(int a, char op, int b) {
       if(a == 0 || b == 0 && op == '%' || op == '*' || op == '/'){
        return 0;
       }else if(op == '+'){
        return a + b;
    }else if(op == '-'){
        return a - b;
    }else if(op == '*'){
        return a * b;
    }else if(op == '/'){
        return a / b;
    }else{
       return a%b;
    }
        
    }

}
