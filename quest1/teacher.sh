#!/bin/bash

# Step 1: isolate the interview number from Annabel Church's street line
INTERVIEW=$(grep -h "SEE INTERVIEW"  streets/* | grep -oE '[0-9]+')

# Step 2: print the interview number
echo "$INTERVIEW"

# Step 3: print the interview contents
cat interviews/interview-"$INTERVIEW"

# Step 4: print the main suspect
echo "$MAIN_SUSPECT"
echo ""