# Bulk Email Generator

Problem:

You're coding on CG, minding your own business, when a notification pops out of nowhere.
Some level 6 PHP coder is following you! It's all fine and dandy. You click on the dismiss button and move on.

Except it doesn't work. The dreaded red popup appears: “Error 492, please send us an email at coders@codingame.com”.
Being the disciplined coder you are, you comply.

But the notification is still there. So you click it again. And read the popup again.
And email coders again. But the notification is still there… You can't really send them the same email all the time, that would be boring!

It's time for some automation!

You are given a bulk email template. You are to expand it to an actual email by selecting a clause randomly when given a choice.

Choices are indicated by (parenthesized text), with clauses separated by pipe symbols |.

Random is defined as using the JBM level-0 twister: “for the ith choice, pick the ith clause (modulo the number of clauses)”.

Input
Line 1: the number N of lines of text to follow
Next N lines: the email template

Important: the email template is the N lines as a whole. This is *not* a line-based problem!
Output
The expanded email.
Constraints
N < 10000
Total template size is less than 10000 characters as well.
Choices don't nest.
