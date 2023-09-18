class Enums {
    public static void main(String[] args) {
        var currentJobStatus = JobStatus.Idle;

        switch (currentJobStatus) {
            case None -> System.out.println("Not available");
            case Busy -> System.out.println("Quite busy");
            case Idle -> System.out.println("Available");
        }
    }
}

enum JobStatus {
    None,
    Busy,
    Idle
}