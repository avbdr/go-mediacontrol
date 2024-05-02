URL=`yt-dlp -f best.1 -g  "$1"`
echo "streaming $URL"
./mediactrl stream "$URL"
