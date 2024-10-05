package sprint;

import java.util.List;

public class UsernameProcessor {
    public String findFirstUsername(List<String> listString){
        return listString.stream()
        .findFirst()
        .orElse("Anonymous");
    }
}
