// ------------------------------------------------------------------------------------------------
type Season {
  Spring
  Summer
  Autumn
  Winter
}

// ------------------------------------------------------------------------------------------------
pub fn demo() {
  echo weather(Spring)
  echo weather(Autumn)
}

// ------------------------------------------------------------------------------------------------
fn weather(season: Season) -> String {
  case season {
    Spring -> "Mild"
    Summer -> "Hot"
    Autumn -> "Windy"
    Winter -> "Cold"
  }
}
