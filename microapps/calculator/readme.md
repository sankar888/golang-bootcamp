# Calculator App
A calculator app takes in an arithmentic expression as input, evaluates it and gives the result back. As of now, this simple calculator app is implemented as a terminal command app, which tkes in inut from console and prints the result back to console.

## Design
### FLow chart
1. App prints messages to console to user to input expression
2. User input an expression, hit enter
3. main fn reads, basicvalidationsuccessful ? invoke calculator : prints usage()
4. calculator validates expression, expressionValid ? evaluate and give result : raise error
5. main fn checks for err, error not nil ? print error : print result
6. wait for next user expression and follow from 2..6
7. Declalre a quit expression and exit the process it input is quit

### ToKnow
1. What is the precision of this calculator ? How big a number can be ? How much precission do existing online calculator supports ?
    1.1 What is the precision of numeric types in go ? What happens if a value exceed a type precision ?
    1.2 Solution in go to handle large numbers ?
2. What is the type of the expression result  ?
3. what operations are supported ?
4. Algorithm to evaluate expressions ?



