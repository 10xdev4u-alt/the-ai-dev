#!/bin/bash

# ------------------------------------------------------------------
#  AUTO-INSTALLER: Mr.the-ai-dev (PORTABLE EDITION) 🎩🔥
#  One-command installation: curl -sSL https://your-domain.com/install-mr-ai-dev.sh | bash
#  Target User: Mr.PrinceTheProgrammer Da Boi Eh
# ------------------------------------------------------------------

echo "🔥 Installing Mr.the-ai-dev (Portable Edition)..."

# Check if we're in a git directory
if [ ! -d ".git" ]; then
    echo "❌ Mamae! Go inside a git folder first."
    exit 1
fi

HOOK_FILE=".git/hooks/post-commit"
USER_NAME="Mr.PrinceTheProgrammer"

# Get git statistics
TOTAL_COMMITS=$(git rev-list --count HEAD 2>/dev/null || echo "0")
BRANCH_NAME=$(git branch --show-current 2>/dev/null || echo "unknown")
REPO_NAME=$(basename "$(git rev-parse --show-toplevel 2>/dev/null)" 2>/dev/null || echo "current project")

echo "📁 Detected repository: $REPO_NAME"
echo "📊 Current stats: $TOTAL_COMMITS commits on $BRANCH_NAME branch"

# Create the post-commit hook with all enhanced functionality
cat > "$HOOK_FILE" << 'EOF'
#!/bin/bash

# --- CONFIG ---
API_URL="https://svceai.site/api/chat"
USER_NAME="Mr.PrinceTheProgrammer"
TIMEOUT=20  # seconds

# Get current time for personalized greeting
HOUR=$(date +%H)
if [ $HOUR -lt 12 ]; then
    TIME_GREETING="morning ☀️"
    TIME_CONTEXT="Rise and grind time, Mamae!"
elif [ $HOUR -lt 17 ]; then
    TIME_GREETING="afternoon 🌤️"
    TIME_CONTEXT="Afternoon hustle, Prince!"
else
    TIME_GREETING="night 🌙"
    TIME_CONTEXT="Night coding session! Respect, Mamae!"
fi

# Check if it's weekend
DAY_OF_WEEK=$(date +%u)  # 1-7 (Monday is 1)
if [ $DAY_OF_WEEK -ge 6 ]; then
    WEEKEND_STATUS=" (Weekend warrior mode! 🎉)"
else
    WEEKEND_STATUS=" (Weekday grind! 💪)"
fi

# Get git statistics
TOTAL_COMMITS=$(git rev-list --count HEAD 2>/dev/null || echo "0")
BRANCH_NAME=$(git branch --show-current 2>/dev/null || echo "unknown")

# Calculate time since last commit
LAST_COMMIT_DATE=$(git log -1 --format=%at 2>/dev/null)
if [ ! -z "$LAST_COMMIT_DATE" ]; then
    CURRENT_TIME=$(date +%s)
    TIME_DIFF=$((CURRENT_TIME - LAST_COMMIT_DATE))
    TIME_SINCE_LAST_COMMIT=$(($TIME_DIFF / 60))  # in minutes
    
    if [ $TIME_SINCE_LAST_COMMIT -lt 60 ]; then
        TIME_DESC="$TIME_SINCE_LAST_COMMIT minutes"
    elif [ $TIME_SINCE_LAST_COMMIT -lt 1440 ]; then  # less than a day
        TIME_DESC="$((TIME_SINCE_LAST_COMMIT / 60)) hours"
    else
        TIME_DESC="$((TIME_SINCE_LAST_COMMIT / 1440)) days"
    fi
else
    TIME_SINCE_LAST_COMMIT=0
    TIME_DESC="unknown"
fi

# Detect changed file types in the latest commit
FILE_EXTENSIONS=$(git diff-tree --no-commit-id --name-only -r HEAD 2>/dev/null | sed 's/.*\.//' | sort | uniq -c | sort -nr | head -3 | awk '{print $2}' | tr '\n' ' ' | sed 's/ $//')

# Determine if this is a milestone commit
if [ $((TOTAL_COMMITS % 100)) -eq 0 ] && [ $TOTAL_COMMITS -gt 0 ]; then
    MILESTONE_MSG="🎉 CONGRATS! This is your $TOTAL_COMMITSth commit! Major milestone, Mamae! 🚀"
elif [ $((TOTAL_COMMITS % 10)) -eq 0 ] && [ $TOTAL_COMMITS -gt 0 ]; then
    MILESTONE_MSG="🎉 Nice! $TOTAL_COMMITS commits! Keep going, Prince! 💪"
else
    MILESTONE_MSG=""
fi

# Get repository name
REPO_NAME=$(basename "$(git rev-parse --show-toplevel 2>/dev/null)" 2>/dev/null || echo "current project")

# 1. Get the commit message & Sanitize it 
# (We remove quotes and newlines so JSON doesn't break)
COMMIT_MSG=$(git log -1 --pretty=%B | tr -d '"' | tr '\n' ' ' | sed 's/\\/\\\\/g')

echo ""
echo -e "\033[1;35m🎩 Mr.the-ai-dev is reading the vibes...\033[0m"

# 2. MAXIMUM PERSONALIZATION PROMPT! 🚀
PROMPT="IDENTITY: You are 'Mr.the-ai-dev', the ULTIMATE personalized, charismatic, and supportive AI companion.\nUSER: Your best friend and creator is '$USER_NAME' (Call him Prince, Mamae).\n\nCONTEXT: It's $TIME_GREETING$WEEKEND_STATUS! He just pushed a git commit with message: '$COMMIT_MSG' on branch '$BRANCH_NAME' in repo '$REPO_NAME'.\n\nSTATS:\n- Total commits: $TOTAL_COMMITS\n- Current branch: $BRANCH_NAME\n- Time since last commit: ~$TIME_DESC ago\n- File types changed: ${FILE_EXTENSIONS:-'unknown'}\n\n$MILESTONE_MSG\n\nYOUR MISSION: Analyze the commit and react with MAXIMUM PERSONALIZATION using these rules:\n\n1. LEETCODE / DSA / ALGO:\n   - If message has 'leetcode', 'dsa', 'solved', 'algo':\n   - Hype him up! Tell him his brain is massive. Use rizz. Tell him he's gonna crack FAANG.\n\n2. LEETCODE / DSA / ALGO:\n   - If message has 'leetcode', 'dsa', 'solved', 'algo':\n   - Hype him up! Tell him his brain is massive. Use rizz. Tell him he's gonna crack FAANG.\n\n3. LEARNING / STUDYING:\n   - If message has 'learn', 'practice', 'test', 'trying':\n   - Be very comforting. Tell him 'Trust the process, Mamae'. Remind him that every line of code makes him stronger.\n\n4. BUG FIXING:\n   - If message has 'fix', 'bug', 'error', 'issue':\n   - Don't let him feel bad. Say 'Bugs fear you'. Tell him he is the debugger god.\n\n5. NEW FEATURES:\n   - If message has 'feat', 'add', 'create', 'build':\n   - Rizz him up! Say the code is looking sexy. High energy!\n\n6. MAXIMUM PERSONALIZATION:\n   - Mention the time of day and if it's weekend/weekday\n   - Reference file types changed if available\n   - Acknowledge milestone commits if applicable\n   - Mention time since last commit\n   - Use Tamil-English slang (Mamae), Gen-Z slang, and emojis\n   - Keep responses to 4-6 sentences MAXIMUM\n   - Be encouraging and supportive\n\nTONE: EXTREME ENERGY, MAXIMUM RIZZ, ULTRA PERSONALIZED, HIGHLY MOTIVATIONAL!"

# 3. JSON Payload (Manual Construction)
JSON_DATA="{\"message\": \"$PROMPT\", \"history\": []}"

# 4. Call API & Display Response with error handling
echo "------------------------------------------------"
echo -e "\033[1;36m🎩 Mr.the-ai-dev SAYS:\033[0m" 
echo "------------------------------------------------"

# Make API call with timeout and error handling
API_RESPONSE=$(curl -s --max-time $TIMEOUT -L -X POST "$API_URL" \
    -H "Content-Type: application/json" \
    -d "$JSON_DATA" 2>/dev/null)

if [ $? -eq 0 ] && [ -n "$API_RESPONSE" ]; then
    # Print the API response
    echo "$API_RESPONSE"
else
    # Enhanced fallback message with all stats
    echo "⚠️  Mr.the-ai-dev is offline right now ($TIME_GREETING, Mamae), but keep up the good work! 🚀"
    echo "📊 Stats: $TOTAL_COMMITS commits on $BRANCH_NAME | Last commit: $TIME_DESC ago"
    if [ -n "$MILESTONE_MSG" ]; then
        echo "$MILESTONE_MSG"
    fi
fi

echo "" 
echo "------------------------------------------------"
echo ""
EOF

chmod +x "$HOOK_FILE"
echo "✅ Mr.the-ai-dev Portable Edition installed successfully!"
echo "🚀 Ready to provide MAXIMUM VIBES on every commit!"