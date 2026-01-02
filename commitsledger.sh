#!/bin/bash

# ------------------------------------------------------------------
#  THE SMART AUTO-DETECT PUSHER
#  - Detects current branch
#  - Detects all remotes
#  - Calculates missing commits for each remote AUTOMATICALLY
# ------------------------------------------------------------------

# 1. Get Current Branch
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

if [ -z "$CURRENT_BRANCH" ]; then
    echo "❌ Bro, this isn't a git repo."
    exit 1
fi

echo "🔥 ------------------------------------------"
echo "🔥 MODE:      Smart Auto-Detect"
echo "🔥 BRANCH:    $CURRENT_BRANCH"
echo "🔥 ------------------------------------------"

# 2. Loop through every remote you have (Origin, Gitlab, Heroku, etc.)
for REMOTE in $(git remote); do
    echo ""
    echo "📡 CONNECTING TO REMOTE: [$REMOTE]..."
    
    # Fetch latest data so we know the truth
    git fetch $REMOTE > /dev/null 2>&1
    
    # 3. Calculate the missing commits
    # Logic: "Show me commits that are on HEAD but NOT on remote/branch"
    # We use --reverse to get them Base -> Up
    MISSING_COMMITS=$(git log --reverse --pretty=format:"%H" $REMOTE/$CURRENT_BRANCH..HEAD 2>/dev/null)
    
    # Check if the command failed (usually means branch doesn't exist on remote yet)
    if [ $? -ne 0 ]; then
        echo "⚠️  Branch '$CURRENT_BRANCH' not found on '$REMOTE'."
        echo "   Assuming ALL commits need to be pushed..."
        MISSING_COMMITS=$(git log --reverse --pretty=format:"%H")
    fi

    # Count them
    COUNT=$(echo "$MISSING_COMMITS" | sed '/^$/d' | wc -l | xargs)

    if [ "$COUNT" -eq "0" ]; then
        echo "✅ Remote '$REMOTE' is already up to date! Chill."
        continue
    fi

    echo "👨‍🍳 Found $COUNT unpushed commits for '$REMOTE'. Cooking..."
    
    CURRENT=1
    
    # 4. The Push Loop
    for commit_hash in $MISSING_COMMITS; do
        echo "   👉 [$CURRENT/$COUNT] Pushing ${commit_hash:0:7} to $REMOTE..."
        
        git push $REMOTE $commit_hash:refs/heads/$CURRENT_BRANCH
        
        ((CURRENT++))
        # Optional tiny pause to be safe
        sleep 0.5 
    done
    
    echo "🎉 '$REMOTE' is now fully synced!"
done

echo ""
echo "🚀 ALL REMOTES UPDATED SUCCESSFULLY! Bro, you're good."
