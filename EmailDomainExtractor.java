package sprint;

import java.util.List;
import java.util.stream.Collectors;

public class EmailDomainExtractor {
    public List<String> extractDomains(List<String> emails) {
        return emails.stream()
                     .filter(email -> isValidEmail(email))
                     .map(email -> extractDomain(email))    
                     .distinct()                          
                     .collect(Collectors.toList()); 
    }
    private boolean isValidEmail(String email){
        String[] parts = email.split("@");
        return parts.length == 2 && !parts[0].isEmpty() && !parts[1].isEmpty();
    }
    private String extractDomain(String email){
        return email.substring(email.indexOf('@') + 1).toLowerCase();
    }
}
