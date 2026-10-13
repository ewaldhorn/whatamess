
void handleObject(Object obj) {
  if (obj instanceof String s && !s.isBlank()) {
    IO.println("Length: " + s.length());
  } else {
    if (obj instanceof String) {
      IO.println("Blank string? No thanks!");
    } else {
      IO.println("Sorry I don't know what to do with " + obj.getClass());
    }
  }
}

void main() {
    Object obj = Integer.valueOf(5);   // should be rejected
    String blank = "";                 // should be rejected
    String fiver = "fiver";            // should have a length of 5
    
    handleObject(obj);
    handleObject(blank);
    handleObject(fiver);  
}
