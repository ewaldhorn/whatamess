module CarsAssemble

let successRate (speed: int): float =
    let mutable rate = 0.0

    if (speed >= 0 && speed < 5) then
        rate <- 1.0
    elif (speed > 4 && speed < 9) then
        rate <- 0.9
    elif (speed = 9) then
        rate <- 0.8
    elif (speed > 9) then
        rate <- 0.77
    else
        rate <- 0.0

    rate    

let productionRatePerHour (speed: int): float =
    let rate = successRate speed
    let answer = float(speed) * 221.0 * rate
    answer

let workingItemsPerMinute (speed: int): int =
    let rate = productionRatePerHour speed
    int(rate) / 60

