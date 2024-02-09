       IDENTIFICATION DIVISION.
       PROGRAM-ID. YACHT.
       DATA DIVISION.
       WORKING-STORAGE SECTION.
       01 WS-RESULT PIC 99 VALUE 0.
       01 WS-CATEGORY PIC X(15).
       01 WS-DICE.
          05 FILLER OCCURS 5 TIMES.
             10 DICE PIC X.
       01 WS-COUNT PIC 9(2) VALUE 0.
       01 WS-SUM PIC 9(3) VALUE 0.
       01 WS-TMP1 PIC 9 VALUE 0.
       01 WS-TMP2 PIC 9 VALUE 0.
       01 WS-LOOP PIC 9 VALUE 0.
       01 WS-I PIC 9 VALUE 0.
       01 WS-J PIC 9 VALUE 0.      
       01 WS-LITTLE PIC X(5) VALUE '12345'.
       01 WS-BIG PIC X(5) VALUE '23456'.
          PROCEDURE DIVISION.
             SUMM.
               MOVE DICE(WS-LOOP) TO WS-COUNT.
               ADD WS-COUNT TO WS-SUM GIVING WS-SUM.
             END-SUMM.

             BUBBLESORT.
                PERFORM VARYING WS-I FROM 1 BY 1 UNTIL WS-I IS GREATER THAN 6
                   PERFORM VARYING WS-J FROM 1 BY 1 UNTIL WS-J IS GREATER THAN 6
                      IF (DICE(WS-J) < DICE(WS-I))
                         MOVE DICE(WS-J) TO WS-TMP1
                         MOVE DICE(WS-I) TO DICE(WS-J)
                         MOVE WS-TMP1 TO DICE(WS-I)
                      END-IF
                   END-PERFORM
                END-PERFORM.
             END-BUBBLESORT.
      
             YACHT.
                MOVE 0 TO WS-RESULT.
                MOVE 0 TO WS-COUNT.
                EVALUATE WS-CATEGORY
                   WHEN "ones"
                      INSPECT WS-DICE TALLYING WS-RESULT FOR ALL '1'
                   WHEN "twos"
                      INSPECT WS-DICE TALLYING WS-COUNT FOR ALL '2'
                      MULTIPLY WS-COUNT BY 2 GIVING WS-RESULT
                   WHEN "threes"
                      INSPECT WS-DICE TALLYING WS-COUNT FOR ALL '3'
                      MULTIPLY WS-COUNT BY 3 GIVING WS-RESULT
                   WHEN "fours"
                      INSPECT WS-DICE TALLYING WS-COUNT FOR ALL '4'
                      MULTIPLY WS-COUNT BY 4 GIVING WS-RESULT
                   WHEN "fives"
                      INSPECT WS-DICE TALLYING WS-COUNT FOR ALL '5'
                      MULTIPLY WS-COUNT BY 5 GIVING WS-RESULT
                   WHEN "sixes"
                      INSPECT WS-DICE TALLYING WS-COUNT FOR ALL '6'
                      MULTIPLY WS-COUNT BY 6 GIVING WS-RESULT
                   WHEN "full house"
                      PERFORM BUBBLESORT
                      MOVE 0 to WS-TMP1
                      MOVE 0 to WS-TMP2
                      INSPECT WS-DICE TALLYING WS-TMP1 FOR ALL DICE(2)
                      INSPECT WS-DICE TALLYING WS-TMP2 FOR ALL DICE(4)
                      IF WS-TMP1 = 3 AND WS-TMP2 = 2 THEN
                         MOVE 0 TO WS-SUM
                         PERFORM SUMM VARYING WS-LOOP FROM 1 BY 1 UNTIL WS-LOOP=6
                         MOVE WS-SUM TO WS-RESULT
                      END-IF
                      IF WS-TMP1 = 2 AND WS-TMP2 = 3 THEN
                         MOVE 0 TO WS-SUM
                         PERFORM SUMM VARYING WS-LOOP FROM 1 BY 1 UNTIL WS-LOOP=6
                         MOVE WS-SUM TO WS-RESULT
                      END-IF
                   WHEN "four of a kind"
                      MOVE 0 TO WS-COUNT
                      INSPECT WS-DICE TALLYING WS-COUNT FOR ALL DICE(1)                  
                      IF WS-COUNT >= 4 THEN
                         MOVE DICE(1) TO WS-LOOP
                         MULTIPLY WS-LOOP BY 4 GIVING WS-RESULT
                      ELSE
                         MOVE 0 TO WS-COUNT
                         INSPECT WS-DICE TALLYING WS-COUNT FOR ALL DICE(2)
                         IF WS-COUNT >= 4 THEN
                            MOVE DICE(2) TO WS-LOOP
                            MULTIPLY WS-LOOP BY 4 GIVING WS-RESULT
                         END-IF
                      END-IF
                   WHEN "little straight"
                      IF WS-DICE = WS-LITTLE THEN
                         MOVE 30 TO WS-RESULT
                      ELSE
                         INSPECT WS-DICE REPLACING FIRST '1' BY SPACES
                         INSPECT WS-DICE REPLACING FIRST '2' BY SPACES
                         INSPECT WS-DICE REPLACING FIRST '3' BY SPACES
                         INSPECT WS-DICE REPLACING FIRST '4' BY SPACES
                         INSPECT WS-DICE REPLACING FIRST '5' BY SPACES
                         MOVE 0 to WS-RESULT
                         IF WS-DICE = SPACES THEN
                            MOVE 30 TO WS-RESULT
                         END-IF
                      END-IF
                   WHEN "big straight"
                      IF WS-DICE = WS-BIG THEN
                         MOVE 30 TO WS-RESULT
                      ELSE
                         INSPECT WS-DICE REPLACING FIRST '2' BY SPACES
                         INSPECT WS-DICE REPLACING FIRST '3' BY SPACES
                         INSPECT WS-DICE REPLACING FIRST '4' BY SPACES
                         INSPECT WS-DICE REPLACING FIRST '5' BY SPACES
                         INSPECT WS-DICE REPLACING FIRST '6' BY SPACES
                         MOVE 0 to WS-RESULT
                         IF WS-DICE = SPACES THEN
                            MOVE 30 TO WS-RESULT
                         END-IF
                      END-IF
                   WHEN "choice"
                      MOVE 0 TO WS-SUM
                      PERFORM SUMM VARYING WS-LOOP FROM 1 BY 1 UNTIL WS-LOOP=6
                      MOVE WS-SUM TO WS-RESULT
                   WHEN "yacht"
                      INSPECT WS-DICE TALLYING WS-COUNT FOR ALL DICE(1)
                      MOVE 0 TO WS-RESULT
                      IF WS-COUNT = 5 THEN
                         MOVE 50 TO WS-RESULT
                      END-IF
                END-EVALUATE.
             YACHT-EXIT.
          EXIT.
