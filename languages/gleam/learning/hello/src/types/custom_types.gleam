// ------------------------------------------------------------------------------------------------
type Season {
  Spring
  Summer
  Autumn
  Winter
}

// ------------------------------------------------------------------------------------------------
pub type Option(inner) {
  Some(inner)
  None
}

// An option of string
const name: Option(String) = Some("Annah")

// An option of int
const level: Option(Int) = Some(10)

// ------------------------------------------------------------------------------------------------
pub fn demo() {
  echo weather(Spring)
  echo weather(Autumn)

  echo name
  echo level
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
