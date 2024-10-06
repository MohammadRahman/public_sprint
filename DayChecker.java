package sprint;

import java.time.LocalDate;

public class DayChecker {
    public static String checkDayType(LocalDate date) {
        switch (date.getDayOfWeek()) {
            case SATURDAY, SUNDAY:
                return "Weekend";
            case WEDNESDAY:
                return "Hump Day!";
            default:
                return "Weekday";
        }
    }
}
