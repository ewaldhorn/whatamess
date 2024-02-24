module LuciansLusciousLasagna
let expectedMinutesInOven = 40
let remainingMinutesInOven spent = 
    expectedMinutesInOven - spent
let preparationTimeInMinutes layers = 
    layers * 2
let elapsedTimeInMinutes layers spent = 
    spent + preparationTimeInMinutes layers
