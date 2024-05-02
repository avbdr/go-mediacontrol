ip="192.168.0.147"
soapaction="urn:schemas-upnp-org:service:ZoneGroupTopology:1"
op="GetZoneGroupAttributes"
url="ZoneGroupTopology/Control"
args=""

while [[ "$#" -gt 0 ]]; do
    case $1 in
        -h|--host) ip="$2"; shift ;;
        -a|--action) soapaction="$2"; shift ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

cat > /tmp/foo <<EOF
<?xml version="1.0" encoding="utf-8"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
<s:Body>
<u:$op xmlns:u="$soapaction">
$args
</u:$op>
</s:Body>
</s:Envelope>
EOF

echo reply from $ip
curl -s -X POST --data-binary @/tmp/foo -H "SoapAction: \"$soapaction#$op\"" $ip:1400/$url > /tmp/out
xmllint --format --nsclean /tmp/out|sed 's/&lt;/</g; s/&gt;/>/g' > /tmp/out2
xmllint --format --nsclean /tmp/out2
echo
