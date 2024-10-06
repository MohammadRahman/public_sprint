package sprint;

import java.time.DayOfWeek;
import java.time.LocalDate;

public class DayChecker {
    public String checkDayType(LocalDate date) {
        DayOfWeek dayOfWeek = date.getDayOfWeek();

        switch (dayOfWeek) {
            case SATURDAY:
            case SUNDAY:
                return "Weekend";
            case WEDNESDAY:
                return "Hump Day!";          
            default:
                return "Weekday";
        }
    }
}
