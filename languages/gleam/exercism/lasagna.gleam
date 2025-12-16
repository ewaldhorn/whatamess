// Please define the expected_minutes_in_oven function
pub fn expected_minutes_in_oven() -> Int {
  40
}

// Please define the remaining_minutes_in_oven function
pub fn remaining_minutes_in_oven(spent: Int) -> Int {
  expected_minutes_in_oven() - spent
}

// Please define the preparation_time_in_minutes function
pub fn preparation_time_in_minutes(layers: Int) -> Int {
  layers * 2
}

// Please define the total_time_in_minutes function
pub fn total_time_in_minutes(layers: Int, spent: Int) -> Int {
  preparation_time_in_minutes(layers) + spent
}

// Please define the alarm function
pub fn alarm() {
  "Ding!"
}
