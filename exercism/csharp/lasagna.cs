class Lasagna
{
    public int ExpectedMinutesInOven() {
        return 40;
    }
    public int RemainingMinutesInOven(int been){
        return ExpectedMinutesInOven() - been;
    }
    public int PreparationTimeInMinutes(int layers){
        return layers * 2;
    }
    public int ElapsedTimeInMinutes(int layers, int spent){
        return PreparationTimeInMinutes(layers) + spent;
    }
}
