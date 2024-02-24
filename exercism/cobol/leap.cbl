       IDENTIFICATION DIVISION.
       PROGRAM-ID. LEAP.
       ENVIRONMENT DIVISION.
       
       DATA DIVISION.
         WORKING-STORAGE SECTION.
           01 WS-YEAR     PIC 9(4).
           01 WS-RESULT   PIC 9(1) VALUE 0.
           01 WS-FOUR     PIC 9(1) VALUE 4.
           01 WS-HUND     PIC 9(3) VALUE 100.
           01 WS-FOURHUND PIC 9(3) VALUE 400.
           01 WS-NUMB     PIC 9(9).
           01 WS-REM      PIC 9(9).
       
       PROCEDURE DIVISION.
          LEAP.
             DIVIDE WS-YEAR BY WS-FOURHUND GIVING WS-NUMB REMAINDER WS-REM.
             IF WS-REM = 0 THEN
               MOVE 1 to WS-RESULT
             ELSE
               DIVIDE WS-YEAR BY WS-FOUR GIVING WS-NUMB REMAINDER WS-REM
               IF WS-REM = 0 THEN
                 MOVE 1 TO WS-RESULT
                 DIVIDE WS-YEAR BY WS-HUND GIVING WS-NUMB REMAINDER WS-REM
                 IF WS-REM = 0 THEN
                   MOVE 0 TO WS-RESULT
                 END-IF
               END-IF
             END-IF.
             CONTINUE.
          LEAP-EXIT.
       EXIT.

